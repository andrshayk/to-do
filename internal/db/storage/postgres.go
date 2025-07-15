package storage

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitPostgres() error {
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		getEnv("POSTGRES_USER", "user"),
		getEnv("POSTGRES_PASSWORD", "password"),
		getEnv("POSTGRES_HOST", "localhost"),
		getEnv("POSTGRES_PORT", "5432"),
		getEnv("POSTGRES_DB", "checklist"),
	)
	var err error
	DB, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		return err
	}
	return DB.Ping(context.Background())
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
