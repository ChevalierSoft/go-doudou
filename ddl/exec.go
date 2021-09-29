package ddl

import (
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"github.com/unionj-cloud/go-doudou/ddl/columnenum"
	"github.com/unionj-cloud/go-doudou/ddl/ddlast"
	"github.com/unionj-cloud/go-doudou/ddl/extraenum"
	"github.com/unionj-cloud/go-doudou/ddl/sortenum"
	"github.com/unionj-cloud/go-doudou/sliceutils"
	"github.com/unionj-cloud/go-doudou/stringutils"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"

	// here must import mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/iancoleman/strcase"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/unionj-cloud/go-doudou/astutils"
	"github.com/unionj-cloud/go-doudou/ddl/codegen"
	"github.com/unionj-cloud/go-doudou/ddl/config"
	"github.com/unionj-cloud/go-doudou/ddl/table"
	"path/filepath"
)

// Ddl is for ddl command
type Ddl struct {
	Dir     string
	Reverse bool
	Dao     bool
	Pre     string
	Df      string
	Conf    config.DbConfig
}

// Exec executes the logic for ddl command
// if Reverse is true, it will generate code from database tables,
// otherwise it will update database tables from structs defined in domain pkg
func (d Ddl) Exec() {
	var db *sqlx.DB
	var err error
	conf := d.Conf
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		conf.User,
		conf.Passwd,
		conf.Host,
		conf.Port,
		conf.Schema,
		conf.Charset)
	conn += `&loc=Asia%2FShanghai&parseTime=True`
	db, err = sqlx.Connect("mysql", conn)
	if err != nil {
		logrus.Panicln(err)
	}
	defer db.Close()
	db.MapperFunc(strcase.ToSnake)
	db = db.Unsafe()

	var existTables []string
	if err = db.Select(&existTables, "show tables"); err != nil {
		logrus.Panicln(err)
	}

	var tables []table.Table
	if !d.Reverse {
		struct2Table(d, tables, existTables, db)
	} else {
		table2struct(d, tables, existTables, db)
	}

	if d.Dao {
		genDao(d, tables)
	}

}

func genDao(d Ddl, tables []table.Table) {
	var err error
	if err = codegen.GenBaseGo(d.Dir, d.Df); err != nil {
		logrus.Errorf("FATAL: %+v\n", err)
	}
	for _, t := range tables {
		if err = codegen.GenDaoGo(d.Dir, t, d.Df); err != nil {
			logrus.Errorf("FATAL: %+v\n", err)
			break
		}
		if err = codegen.GenDaoImplGo(d.Dir, t, d.Df); err != nil {
			logrus.Errorf("FATAL: %+v\n", err)
			break
		}
		if err = codegen.GenDaoSQL(d.Dir, t, d.Df); err != nil {
			logrus.Errorf("FATAL: %+v\n", err)
			break
		}
	}
}

func table2struct(d Ddl, tables []table.Table, existTables []string, db *sqlx.DB) {
	var err error
	if err = os.MkdirAll(d.Dir, os.ModePerm); err != nil {
		logrus.Panicln(err)
	}
	for _, t := range existTables {
		if stringutils.IsNotEmpty(d.Pre) && !strings.HasPrefix(t, d.Pre) {
			continue
		}
		var dbIndice []table.DbIndex
		if err = db.Select(&dbIndice, fmt.Sprintf("SHOW INDEXES FROM %s", t)); err != nil {
			logrus.Panicln(err)
		}

		idxMap := make(map[string][]table.DbIndex)

		for _, idx := range dbIndice {
			if val, exists := idxMap[idx.KeyName]; exists {
				val = append(val, idx)
				idxMap[idx.KeyName] = val
			} else {
				idxMap[idx.KeyName] = []table.DbIndex{
					idx,
				}
			}
		}

		indexes, colIdxMap := idxListAndMap(idxMap)

		var columns []table.DbColumn
		if err = db.Select(&columns, fmt.Sprintf("SHOW FULL COLUMNS FROM %s", t)); err != nil {
			logrus.Panicln(err)
		}

		var cols []table.Column
		var fields []astutils.FieldMeta
		for _, item := range columns {
			col := dbColumn2Column(item, colIdxMap, t)
			fields = append(fields, col.Meta)
			cols = append(cols, col)
		}

		domain := astutils.StructMeta{
			Name:   strcase.ToCamel(strings.TrimPrefix(t, d.Pre)),
			Fields: fields,
		}

		var pkColumn table.Column
		for _, column := range cols {
			if column.Pk {
				pkColumn = column
				break
			}
		}

		tables = append(tables, table.Table{
			Name:    t,
			Columns: cols,
			Pk:      pkColumn.Name,
			Indexes: indexes,
			Meta:    domain,
		})

		dfile := filepath.Join(d.Dir, strings.ToLower(domain.Name)+".go")
		if _, err = os.Stat(dfile); os.IsNotExist(err) {
			if err = codegen.GenDomainGo(d.Dir, domain); err != nil {
				logrus.Errorf("FATAL: %+v\n", err)
			}
		} else {
			logrus.Warnf("file %s already exists", dfile)
		}
	}
}

func idxListAndMap(idxMap map[string][]table.DbIndex) ([]table.Index, map[string][]table.IndexItem) {
	var indexes []table.Index
	colIdxMap := make(map[string][]table.IndexItem)
	for k, v := range idxMap {
		if len(v) == 0 {
			continue
		}
		items := make([]table.IndexItem, len(v))
		for i, idx := range v {
			var sort sortenum.Sort
			if idx.Collation == "B" {
				sort = sortenum.Desc
			} else {
				sort = sortenum.Asc
			}
			items[i] = table.IndexItem{
				Unique: !v[0].NonUnique,
				Name:   k,
				Column: idx.ColumnName,
				Order:  idx.SeqInIndex,
				Sort:   sort,
			}
			if val, exists := colIdxMap[idx.ColumnName]; exists {
				val = append(val, items[i])
				colIdxMap[idx.ColumnName] = val
			} else {
				colIdxMap[idx.ColumnName] = []table.IndexItem{
					items[i],
				}
			}
		}
		indexes = append(indexes, table.Index{
			Unique: !v[0].NonUnique,
			Name:   k,
			Items:  items,
		})
	}
	return indexes, colIdxMap
}

func dbColumn2Column(item table.DbColumn, colIdxMap map[string][]table.IndexItem, t string) table.Column {
	extra := item.Extra
	if strings.Contains(extra, "auto_increment") {
		extra = ""
	}
	extra = strings.TrimSpace(strings.TrimPrefix(extra, "DEFAULT_GENERATED"))
	if stringutils.IsNotEmpty(item.Comment) {
		extra += fmt.Sprintf(" comment '%s'", item.Comment)
	}
	extra = strings.TrimSpace(extra)
	var defaultVal string
	if item.Default != nil {
		defaultVal = *item.Default
	}
	col := table.Column{
		Table:         t,
		Name:          item.Field,
		Type:          columnenum.ColumnType(item.Type),
		Default:       defaultVal,
		Pk:            table.CheckPk(item.Key),
		Nullable:      table.CheckNull(item.Null),
		Unsigned:      table.CheckUnsigned(item.Type),
		Autoincrement: table.CheckAutoincrement(item.Extra),
		Extra:         extraenum.Extra(extra),
		AutoSet:       table.CheckAutoSet(defaultVal),
		Indexes:       colIdxMap[item.Field],
	}
	col.Meta = table.NewFieldFromColumn(col)
	return col
}

func struct2Table(d Ddl, tables []table.Table, existTables []string, db *sqlx.DB) {
	var (
		files []string
		err   error
	)
	err = filepath.Walk(d.Dir, astutils.Visit(&files))
	if err != nil {
		logrus.Panicln(err)
	}
	sc := astutils.NewStructCollector(astutils.ExprString)
	for _, file := range files {
		fset := token.NewFileSet()
		root, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
		if err != nil {
			logrus.Panicln(err)
		}
		ast.Walk(sc, root)
	}

	flattened := ddlast.FlatEmbed(sc.Structs)
	for _, sm := range flattened {
		tables = append(tables, table.NewTableFromStruct(sm, d.Pre))
	}
	for _, t := range tables {
		if sliceutils.StringContains(existTables, t.Name) {
			var columns []table.DbColumn
			if err = db.Select(&columns, fmt.Sprintf("desc %s", t.Name)); err != nil {
				logrus.Panicln(err)
			}
			var existColumnNames []interface{}
			for _, dbCol := range columns {
				existColumnNames = append(existColumnNames, dbCol.Field)
			}
			existColSet := mapset.NewSetFromSlice(existColumnNames)

			for _, col := range t.Columns {
				if existColSet.Contains(col.Name) {
					if err = table.ChangeColumn(db, col); err != nil {
						logrus.Infof("FATAL: %+v\n", err)
					}
				} else {
					if err = table.AddColumn(db, col); err != nil {
						logrus.Infof("FATAL: %+v\n", err)
					}
				}
			}
		} else {
			if err = table.CreateTable(db, t); err != nil {
				logrus.Errorf("FATAL: %+v\n", err)
			}
		}
	}
}
