package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

const sqlDriver = "postgres"

func Open(path string) (*sql.DB, error) {
	database, err := sql.Open(sqlDriver, path)
	if err != nil {
		return nil, err
	}
	return database, nil
}
