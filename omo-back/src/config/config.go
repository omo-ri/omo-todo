package config

import (
	"os"
	"time"
)

var PostgresHost = Get("POSTGRES_HOST", "localhost")
var PostgresPort = Get("POSTGRES_PORT", "5432")
var PostgresUser = Get("POSTGRES_USER", "postgres")
var PostgresPassword = Get("POSTGRES_PASSWORD", "postgres")
var PostgresDatabase = Get("POSTGRES_DATABASE", "postgres")

var JwtSecret = Get("JWT_SECRET", "secret")
var JwtExpireTime = 24 * time.Hour

var Port = Get("PORT", "8080")

func Get(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
