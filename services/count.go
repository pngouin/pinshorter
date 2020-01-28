package services

import (
	"database/sql"
	"github.com/ZooPin/pinshorter/db"
	"github.com/ZooPin/pinshorter/models"
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
