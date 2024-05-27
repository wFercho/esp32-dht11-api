package postgres

import (
	"database/sql"
	"fmt"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	envFile, err := godotenv.Read(".env.local")

	if err != nil {
		return nil, err
	}

	POSTGRES_DATABASE := envFile["POSTGRES_DATABASE"]
	POSTGRES_USERNAME := envFile["POSTGRES_USERNAME"]
	POSTGRES_PASSWORD := envFile["POSTGRES_PASSWORD"]

	if len(POSTGRES_DATABASE) == 0 || len(POSTGRES_USERNAME) == 0 || len(POSTGRES_PASSWORD) == 0 {
		return nil, &argError{3, "Provide all the postgres db credentials"}
	}

	//"user= dbname= password= sslmode=disable"
	connStr := fmt.Sprintf(`user=%s dbname=%s password=%s sslmode=disable`, POSTGRES_USERNAME, POSTGRES_DATABASE, POSTGRES_PASSWORD)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}
