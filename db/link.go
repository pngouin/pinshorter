package db

import (
	"database/sql"
	"github.com/ZooPin/pinshorter/models"
)

func NewLink(db *sql.DB) Link {
	return Link{
		db: db,
	}
}

type Link struct {
	db    *sql.DB
	count Count
}

func (l Link) Scan(row *sql.Row) (models.Link, error) {
	var result models.Link
	var created string

	err := row.Scan(
		&result.Id,
		&result.Title,
		&result.URL,
		&result.ApiPoint,
		&result.Count,
		&created,
	)

	if err != nil {
		return result, err
	}

	result.CreatedAt, err = scanDate(created)
	if err != nil {
		return result, err
	}

	return result, err
}

func (l Link) ScanRows(rows *sql.Rows) ([]models.Link, error) {
	defer rows.Close()
	var links []models.Link

	for rows.Next() {
		var result models.Link
		var created string

		err := rows.Scan(
			&result.Id,
			&result.Title,
			&result.URL,
			&result.ApiPoint,
			&created,
		)

		if err != nil {
			return links, err
		}

		result.CreatedAt, err = scanDate(created)
		if err != nil {
			return links, err
		}
		links = append(links, result)
	}
	return links, nil
}

func (l Link) Create(link models.Link) (models.Link, error) {
	link.Id = createUUID()
	_, err := l.db.Exec("INSERT INTO link (link_id, title, url, api_point, created_at, user_id) VALUES (?, ?, ?, ?, CAST(datetime('now') as TEXT), ?)",
		link.Id, link.Title, link.URL, link.Count, link.ApiPoint, link.User.Id)
	return link, err
}

func (l Link) Delete(link models.Link) error {
	_, err := l.db.Exec("UPDATE link SET deleted_at=CAST(datetime('now') as TEXT) WHERE link_id=? AND user_id=? AND deleted_at is NULL", link.Id, link.User.Id)
	return err
}

func (l Link) GetByID(link models.Link) (models.Link, error) {
	row := l.db.QueryRow("SELECT link_id, title, url, api_point, created_at from link where link_id=? and deleted_at is null", link.Id)
	result, err := l.Scan(row)
	return result, err
}

func (l Link) GetAllByUser(user models.UserInfo) ([]models.Link, error) {
	rows, err := l.db.Query("select link_id, title, url, api_point, link.created_at from link join user u on link.user_id = u.user_id where u.user_id=? and u.deleted_at is null", user.Id)
	if err != nil {
		return nil, err
	}
	return l.ScanRows(rows)
}

func (l Link) IsApiPointExist(str string) bool {
	var result int
	row := l.db.QueryRow("select count(*) from link where api_point=? ", str)
	err := row.Scan(&result)
	if err != nil {
		return false
	}
	return result != 0
}

func (l Link) GetByAPIPoint(link models.Link) (models.Link, error) {
	row := l.db.QueryRow("SELECT link_id, title, url, api_point, link.created_at from link where api_point=? and deleted_at is null ", link.ApiPoint)
	return l.Scan(row)
}
