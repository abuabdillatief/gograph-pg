package database

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v9"
)

//I'd like to see my query in the future, thus
//I need to implement an interface of the logger

//DBLogger ...
type DBLogger struct{}

//BeforeQuery ...
func (d DBLogger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

//AfterQuery ...
func (d DBLogger) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	fmt.Println(q.FormattedQuery())
	return nil
}

//New ...
func New(opts *pg.Options) *pg.DB {
	return pg.Connect(opts)
}
