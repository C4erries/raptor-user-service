package postgres

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/c4erries/raptor-user-service/internal/config"
)

type Postgres struct {
	DB *sql.DB
}

func New(config config.StorageConfig) (*Postgres, error) {
	const op = "postgres.Connect"

	db, err := connect()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Postgres{
		db,
	}, nil
}

func connect() (*sql.DB, error) {
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDB := os.Getenv("POSTGRES_DB")
	postgresHost := os.Getenv("POSTGRES_HOST")

	connectStr := "host=" + postgresHost + " user=" + postgresUser +
		" password=" + postgresPassword +
		" dbname=" + postgresDB + " sslmode=disable"

	db, err := sql.Open("postgres", connectStr)

	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
