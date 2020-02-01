package db

import (
	"database/sql"
	"github.com/ZooPin/pinshorter/models"
)

func NewCount(sql *sql.DB) Count {
	return Count{db: sql}
}

type Count struct {
	db *sql.DB
}

func (c Count) Add(link models.Link) error {
	id := createUUID()
	_, err := c.db.Exec("INSERT INTO count (count_id, date, link_id) VALUES ($1, now(), $2);", id, link.Id)
	return err
}

func (c Count) Count(link models.Link) (int, error) {
	var count int
	row := c.db.QueryRow("SELECT count(*) FROM count where link_id=$1", link.Id)
	err := row.Scan(&count)
	return count, err
}
