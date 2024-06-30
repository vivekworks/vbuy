package postgres

import (
    "context"
    "fmt"
    "github.com/jackc/pgx/v5/pgxpool"
)

const UrlFormat = "postgres://%s:%s@%s:%d/%s"

type DB struct {
    pool *pgxpool.Pool
}

func NewDB(ctx context.Context, dbURL string) (*DB, error) {
    dbPool, err := pgxpool.New(ctx, dbURL)
    if err != nil {
        return nil, fmt.Errorf("unable to create connection pool: %v", err)
    }
    return &DB{pool: dbPool}, nil
}

func (d *DB) Close() {
    d.pool.Close()
}
