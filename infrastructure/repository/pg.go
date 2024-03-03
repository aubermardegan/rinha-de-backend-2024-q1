package repository

import (
	"context"
	"fmt"

	"github.com/amardegan/rinha-de-backend-2024-q1/config"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

type DBConn struct {
	Pool *pgxpool.Pool
}

func InitPostgreSQL(ctx context.Context) (*DBConn, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable&pool_max_conns=100",
		config.DATABASE_USER, config.DATABASE_PASSWORD, config.DATABASE_HOST, config.DATABASE_PORT, config.DATABASE_NAME)
	cfg, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}
	cfg.MaxConns = 10

	db, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}

	err = db.Ping(ctx)

	return &DBConn{Pool: db}, err
}
