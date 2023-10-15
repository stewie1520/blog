package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func NewPostgresDBX(connectionURL string, options ...Option) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(connectionURL)
	if err != nil {
		zap.S().Errorf("Unable to parse connection pool config: %v\n", err)
		return nil, err
	}

	for _, option := range options {
		option(config)
	}

	conn, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		zap.S().Errorf("Unable to create connection pool: %v\n", err)
		return nil, err
	}

	err = conn.Ping(context.Background())
	if err != nil {
		zap.S().Errorf("Unable to connect to database: %v\n", err)
		return nil, err
	}

	zap.S().Info("Connected to database 🎉")

	return conn, nil
}
