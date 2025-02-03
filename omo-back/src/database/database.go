package database

import (
	"context"
	"errors"
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

	config, err := GetConfig()

	if err != nil {
		return fmt.Errorf("failed to get config: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	Pool, err = pgxpool.NewWithConfig(ctx, config)

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return fmt.Errorf("failed to create database connection pool: timeout")
		}
		return fmt.Errorf("failed to create database connection pool: %v", err)
	}

	// 测试连接
	if err := Pool.Ping(ctx); err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}
	return nil
}

func GetConfig() (*pgxpool.Config, error) {
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
		return nil, fmt.Errorf("failed to parse config: %v", err)
	}
	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnLifetime = 30 * time.Minute
	return config, nil
}
