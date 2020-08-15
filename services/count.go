package services

import (
	"database/sql"
	"github.com/pngouin/pinshorter/db"
	"github.com/pngouin/pinshorter/models"
)

func NewCount(sql *sql.DB) Count {
	return Count{database: db.NewCount(sql)}
}

type Count struct {
	database db.Count
}

func (c Count) GetCount(link models.Link) (int, error) {
	return c.database.Count(link)
}

func (c Count) Add(link models.Link) error {
	return c.database.Add(link)
}
