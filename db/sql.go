package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

const sqlDriver = "sqlite3"

func Open(path string) (*sql.DB, error) {
	database, err := sql.Open(sqlDriver, path)
	if err != nil {
		return nil, err
	}
	return database, nil
}
