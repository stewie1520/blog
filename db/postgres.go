package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/stewie1520/blog/log"
)

func NewPostgresDBX(connectionURL string, options ...Option) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(connectionURL)
	if err != nil {
		log.S().Errorf("Unable to parse connection pool config: %v\n", err)
		return nil, err
	}

	for _, option := range options {
		option(config)
	}

	conn, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.S().Errorf("Unable to create connection pool: %v\n", err)
		return nil, err
	}

	err = conn.Ping(context.Background())
	if err != nil {
		log.S().Errorf("Unable to connect to database: %v\n", err)
		return nil, err
	}

	log.S().Info("Connected to database 🎉")

	return conn, nil
}
