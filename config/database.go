// config/database.go

package config

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	db "njajal-go/db/sqlc"
	"os"
)

func DbInit() *db.Queries {
	DatabaseURL := os.Getenv("DB_URL")
	if DatabaseURL == "" {
		panic("DB_URL environment variable is not set")
	}

	ctx := context.Background()
	connPool, err := pgxpool.New(ctx, DatabaseURL)
	if err != nil {
		panic(fmt.Sprintf("Error connecting to the database: %v", err))
	}

	return db.New(connPool)
}
