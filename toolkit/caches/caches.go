package caches

import (
	"gorm.io/gorm/callbacks"
	"gorm.io/gorm/clause"
	"sync"

	"gorm.io/gorm"
)

type Caches struct {
	Conf *Config

	queue   *sync.Map
	queryCb func(*gorm.DB)
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

	c.queryCb = db.Callback().Query().Get("gorm:query")

	err := db.Callback().Query().Replace("gorm:query", c.Query)
	if err != nil {
		return err
	}

	err = db.Callback().Create().After("gorm:after_create").Register("cache:after_create", c.AfterWrite)
	if err != nil {
		return err
	}

	err = db.Callback().Create().After("gorm:after_delete").Register("cache:after_delete", c.AfterWrite)
	if err != nil {
		return err
	}

	err = db.Callback().Create().After("gorm:after_update").Register("cache:after_update", c.AfterWrite)
	if err != nil {
		return err
	}

	return nil
}

func (c *Caches) Query(db *gorm.DB) {
	if c.Conf.Easer == false && c.Conf.Cacher == nil {
		c.queryCb(db)
		return
	}

	identifier := buildIdentifier(db)

	if c.checkCache(db, identifier) {
		return
	}

	c.ease(db, identifier)
	if db.Error != nil {
		return
	}

	c.storeInCache(db, identifier)
	if db.Error != nil {
		return
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

func (c *Caches) ease(db *gorm.DB, identifier string) {
	if c.Conf.Easer == false {
		c.queryCb(db)
		return
	}

	res := ease(&queryTask{
		id:      identifier,
		db:      db,
		queryCb: c.queryCb,
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

func (c *Caches) checkCache(db *gorm.DB, identifier string) bool {
	if c.Conf.Cacher != nil {
		if res := c.Conf.Cacher.Get(identifier); res != nil {
			res.replaceOn(db)
			return true
		}
	}
	return false
}

func getFromClause(db *gorm.DB) *clause.From {
	if db == nil || db.Statement == nil {
		return &clause.From{}
	}
	c, ok := db.Statement.Clauses[clause.From{}.Name()]
	if !ok || c.Expression == nil {
		return &clause.From{}
	}
	from, ok := c.Expression.(clause.From)
	if !ok {
		return &clause.From{}
	}
	return &from
}

func getTables(db *gorm.DB) []string {
	// Find all table names within the sql statement as cache tags
	from := getFromClause(db)
	var tables []string
	for _, item := range from.Tables {
		tables = append(tables, item.Name)
	}
	for _, item := range from.Joins {
		tables = append(tables, item.Table.Name)
	}
	return tables
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
