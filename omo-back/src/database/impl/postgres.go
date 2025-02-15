package impl

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"net/url"
	omoconfig "omo-back/src/config"
	"time"
)

type PostgresServiceImpl struct {
	db *pgxpool.Pool
}

type UUID = uuid.UUID

func NewPostgresServiceImpl() *PostgresServiceImpl {
	config, err := GetConfig()

	if err != nil {
		log.Fatalf("Failed to get config: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	db, err := pgxpool.NewWithConfig(ctx, config)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 测试连接
	if err := db.Ping(ctx); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	return &PostgresServiceImpl{db: db}
}

func (p *PostgresServiceImpl) Close() {
	if p.db != nil {
		p.db.Close()
		p.db = nil
	}
}

func (p *PostgresServiceImpl) Exec(ctx context.Context, query string, args ...any) (pgconn.CommandTag, error) {
	resp, err := p.db.Exec(ctx, query, args...)
	if err != nil {
		return resp, fmt.Errorf("failed to exec query: %v", err)

	}
	return resp, nil
}

func (p *PostgresServiceImpl) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return p.db.QueryRow(ctx, sql, args...)
}

func (p *PostgresServiceImpl) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return p.db.Query(ctx, sql, args...)
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
