package postgresql

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog"
)

func ConnectDB(log zerolog.Logger, ctx context.Context) *pgx.Conn {
	conn, err := pgx.Connect(ctx, "postgres://postgres:6355@localhost:5432/crawler_db")
	if err != nil {
		log.Error().Err(err).Msg("Failed to connec to postgres!")
		os.Exit(1)
	}

	return conn
}

func InsertUrl(ctx context.Context, db *pgx.Conn, url string) error {
	query := `
		INSERT INTO urls (url)
		VALUES ($1)
	`
	_, err := db.Exec(ctx, query, url)
	return err
}
