package db

import (
	"context"
	"database/sql"
	"errors"
)

type DB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row

	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

func NewPostgres(dsn string) (DB, error) {
	return nil, errors.New("fuck")
}

func InTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return fn(ctx)
}

func Migrate(dsn string) error {
	return errors.New("fuck")
}
