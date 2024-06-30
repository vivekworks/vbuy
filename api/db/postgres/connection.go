package postgres

import (
    "context"
    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/vivekworks/vbuy"
    "go.uber.org/zap"
)

const UrlFormat = "postgres://%s:%s@%s:%d/%s"

type DB struct {
    pool *pgxpool.Pool
}

func NewDB(ctx context.Context, dbURL string) (*DB, error) {
    logger := vbuy.LoggerFromContext(ctx)
    dbPool, err := pgxpool.New(ctx, dbURL)
    if err != nil {
        logger.Error("unable to create connection pool", zap.Error(err))
        return nil, err
    }
    return &DB{pool: dbPool}, nil
}

func (d *DB) Close() {
    d.pool.Close()
}
