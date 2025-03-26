package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Database interface {
	Close(ctx context.Context) error
	Ping() error
	ExecContext(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
	QueryContext(ctx context.Context, query string, args ...any) (pgx.Rows, error)
	Scheme() string
	QueryRowContext(ctx context.Context, query string, args ...interface{}) pgx.Row
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
}

type Rows interface {
	Scan(dest ...any) error
	Next() bool
	Close() error
}
