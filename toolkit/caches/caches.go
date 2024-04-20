package caches

import (
	"fmt"
	"github.com/auxten/postgresql-parser/pkg/sql/parser"
	"github.com/auxten/postgresql-parser/pkg/sql/sem/tree"
	"github.com/auxten/postgresql-parser/pkg/walk"
	"github.com/samber/lo"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/stringutils"
	"github.com/xwb1989/sqlparser"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/callbacks"
	"strings"
	"sync"
)

type Caches struct {
	Conf  *Config
	queue *sync.Map
}

type Config struct {
	Easer  bool
	Cacher Cacher
}

func (c *Caches) Name() string {
	return "gorm:caches"
}

func (c *Caches) Initialize(db *gorm.DB) error {
	if c.Conf == nil {
		c.Conf = &Config{
			Easer:  false,
			Cacher: nil,
		}
	}

	if c.Conf.Easer {
		c.queue = &sync.Map{}
	}

	callback := db.Callback().Query().Get("gorm:query")

	err := db.Callback().Query().Replace("gorm:query", c.Query(callback))
	if err != nil {
		return err
	}

	err = db.Callback().Create().After("gorm:after_create").Register("cache:after_create", c.AfterWrite)
	if err != nil {
		return err
	}

	err = db.Callback().Delete().After("gorm:after_delete").Register("cache:after_delete", c.AfterWrite)
	if err != nil {
		return err
	}

	err = db.Callback().Update().After("gorm:after_update").Register("cache:after_update", c.AfterWrite)
	if err != nil {
		return err
	}

	err = db.Callback().Raw().After("gorm:raw").Register("cache:after_raw", c.AfterWrite)
	if err != nil {
		return err
	}

	return nil
}

func (c *Caches) Query(callback func(*gorm.DB)) func(*gorm.DB) {
	return func(db *gorm.DB) {
		if c.Conf.Easer == false && c.Conf.Cacher == nil {
			callback(db)
			return
		}

		identifier := buildIdentifier(db)

		if db.DryRun {
			return
		}

		if res, ok := c.checkCache(identifier); ok {
			res.replaceOn(db)
			return
		}

		c.ease(db, identifier, callback)
		if db.Error != nil {
			return
		}

		c.storeInCache(db, identifier)
		if db.Error != nil {
			return
		}
	}
}

func (c *Caches) AfterWrite(db *gorm.DB) {
	if c.Conf.Easer == false && c.Conf.Cacher == nil {
		return
	}

	callbacks.BuildQuerySQL(db)

	tables := getTables(db)
	if len(tables) == 1 {
		c.deleteCache(db, tables[0])
	} else {
		c.deleteCache(db, tables[0], tables[1:]...)
	}

	if db.Error != nil {
		return
	}
}

func (c *Caches) ease(db *gorm.DB, identifier string, callback func(*gorm.DB)) {
	if c.Conf.Easer == false {
		callback(db)
		return
	}

	res := ease(&queryTask{
		id:      identifier,
		db:      db,
		queryCb: callback,
	}, c.queue).(*queryTask)

	if db.Error != nil {
		return
	}

	if res.db.Statement.Dest == db.Statement.Dest {
		return
	}

	q := Query{
		Dest:         db.Statement.Dest,
		RowsAffected: db.Statement.RowsAffected,
	}
	q.replaceOn(res.db)
}

func (c *Caches) checkCache(identifier string) (res *Query, ok bool) {
	if c.Conf.Cacher != nil {
		if res = c.Conf.Cacher.Get(identifier); res != nil {
			return res, true
		}
	}
	return nil, false
}

func getTables(db *gorm.DB) []string {
	callbacks.BuildQuerySQL(db)
	switch db.Dialector.(type) {
	case *mysql.Dialector:
		return getTablesMysql(db)
	case *postgres.Dialector:
		return getTablesPostgres(db)
	}
	return nil
}

func getTablesMysql(db *gorm.DB) []string {
	stmt, err := sqlparser.Parse(db.Statement.SQL.String())
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
	tableNames := make([]string, 0)
	_ = sqlparser.Walk(func(node sqlparser.SQLNode) (kontinue bool, err error) {
		switch node := node.(type) {
		case sqlparser.TableName:
			tableNames = append(tableNames, node.Name.CompliantName())
		}
		return true, nil
	}, stmt)
	tableNames = lo.Filter(tableNames, func(x string, index int) bool {
		return stringutils.IsNotEmpty(x)
	})
	tableNames = lo.Uniq(tableNames)
	return tableNames
}

func getTablesPostgres(db *gorm.DB) []string {
	tableNames := make([]string, 0)
	sql := db.Statement.SQL.String()
	w := &walk.AstWalker{
		Fn: func(ctx interface{}, node interface{}) (stop bool) {
			//log.Printf("%T", node)
			switch expr := node.(type) {
			case *tree.TableName:
				var sb strings.Builder
				fmtCtx := tree.NewFmtCtx(tree.FmtSimple)
				expr.TableNamePrefix.Format(fmtCtx)
				sb.WriteString(fmtCtx.String())

				if sb.Len() > 0 {
					sb.WriteString(".")
				}

				fmtCtx = tree.NewFmtCtx(tree.FmtSimple)
				expr.TableName.Format(fmtCtx)
				sb.WriteString(fmtCtx.String())

				tableNames = append(tableNames, sb.String())
			case *tree.Insert:
				fmtCtx := tree.NewFmtCtx(tree.FmtSimple)
				expr.Table.Format(fmtCtx)
				tableName := fmtCtx.String()
				tableNames = append(tableNames, tableName)
			case *tree.Update:
				fmtCtx := tree.NewFmtCtx(tree.FmtSimple)
				expr.Table.Format(fmtCtx)
				tableName := fmtCtx.String()
				tableNames = append(tableNames, tableName)
			case *tree.Delete:
				fmtCtx := tree.NewFmtCtx(tree.FmtSimple)
				expr.Table.Format(fmtCtx)
				tableName := fmtCtx.String()
				tableNames = append(tableNames, tableName)
			}
			return false
		},
	}
	stmts, err := parser.Parse(sql)
	if err != nil {
		return nil
	}
	_, err = w.Walk(stmts, nil)
	if err != nil {
		return nil
	}
	return tableNames
}

func (c *Caches) storeInCache(db *gorm.DB, identifier string) {
	if c.Conf.Cacher != nil {
		err := c.Conf.Cacher.Store(identifier, &Query{
			Tags:         getTables(db),
			Dest:         db.Statement.Dest,
			RowsAffected: db.Statement.RowsAffected,
		})
		if err != nil {
			_ = db.AddError(err)
		}
	}
}

func (c *Caches) deleteCache(db *gorm.DB, tag string, tags ...string) {
	if c.Conf.Cacher != nil {
		err := c.Conf.Cacher.Delete(tag, tags...)
		if err != nil {
			_ = db.AddError(err)
		}
	}
}
