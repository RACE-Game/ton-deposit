package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type DB struct {
	scheme  string
	pgxConn *pgxpool.Pool
}

func New(ctx context.Context, databaseURL string, scheme string, maxConn int32) (*DB, error) {

	// urlExample := "postgres://username:password@localhost:5432/database_name"
	//conn, err := pgx.Connect(ctx, databaseURL)
	conn, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		return nil, errors.Wrap(err, "unable to connect to database")
	}

	conn.Config().MaxConns = int32(maxConn)

	return &DB{
		scheme:  scheme,
		pgxConn: conn,
	}, nil
}

func (db *DB) Close(ctx context.Context) error {
	db.pgxConn.Close()
	return nil
}

func (db *DB) Scheme() string {
	return db.scheme
}
func (db *DB) Ping() error {
	return db.pgxConn.Ping(context.Background())
}

func (db *DB) ExecContext(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return db.pgxConn.Exec(ctx, query, args...)
}

func (db *DB) QueryContext(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	return db.pgxConn.Query(ctx, query, args...)
}

func (db *DB) QueryRowContext(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return db.pgxConn.QueryRow(ctx, query, args...)
}

func (db *DB) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	return db.pgxConn.SendBatch(ctx, b)
}
