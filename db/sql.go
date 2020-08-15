package db

import (
	"database/sql"
	"github.com/pngouin/pinshorter/db/query"
	_ "github.com/lib/pq"
)

const sqlDriver = "postgres"

func Open(path string) (*sql.DB, error) {
	database, err := sql.Open(sqlDriver, path)
	if err != nil {
		return nil, err
	}

	_, err = database.Exec(query.CreatePostgresql)

	return database, err
}
