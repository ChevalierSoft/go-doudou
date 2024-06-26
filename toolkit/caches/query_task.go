package caches

import "github.com/wubin1989/gorm"

type queryTask struct {
	id      string
	db      *gorm.DB
	queryCb func(db *gorm.DB)
}

func (q *queryTask) GetId() string {
	return q.id
}

func (q *queryTask) Run() {
	q.queryCb(q.db)
}
