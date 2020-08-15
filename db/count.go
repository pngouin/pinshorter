package db

import (
	"database/sql"
	"github.com/pngouin/pinshorter/db/query"
	"github.com/pngouin/pinshorter/models"
)

func NewCount(sql *sql.DB) Count {
	return Count{db: sql}
}

type Count struct {
	db *sql.DB
}

func (c Count) Add(link models.Link) error {
	id := createUUID()
	_, err := c.db.Exec(query.CountCreate, id, link.Id)
	return err
}

func (c Count) Count(link models.Link) (int, error) {
	var count int
	row := c.db.QueryRow(query.CountGet, link.Id)
	err := row.Scan(&count)
	return count, err
}
