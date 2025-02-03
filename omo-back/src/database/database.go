package database

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/url"
	omoconfig "omo-back/src/config"
	"time"
)

var Pool *pgxpool.Pool

type UUID = uuid.UUID

func Init() error {
	password := omoconfig.PostgresPassword
	escapedPassword := url.QueryEscape(password)
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		omoconfig.PostgresUser,
		escapedPassword,
		omoconfig.PostgresHost,
		omoconfig.PostgresPort,
		omoconfig.PostgresDatabase,
	)
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return fmt.Errorf("failed to parse config: %v", err)
	}
	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnLifetime = 30 * time.Minute

	Pool, err = pgxpool.NewWithConfig(context.Background(), config)

	if err != nil {
		return fmt.Errorf("failed to create connection pool: %v", err)
	}

	// 测试连接
	if err := Pool.Ping(context.Background()); err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}
	return nil
}
