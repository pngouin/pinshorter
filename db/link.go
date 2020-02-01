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
	_, err := l.db.Exec("INSERT INTO links (link_id, title, url, api_point, created_at, user_id) VALUES ($1, $2, $3, $4, now(), $5)",
		link.Id, link.Title, link.URL, link.ApiPoint, link.User.Id)
	return link, err
}

func (l Link) Delete(link models.Link) error {
	_, err := l.db.Exec("UPDATE links SET deleted_at=now() WHERE link_id=$1 AND user_id=$2 AND deleted_at is NULL", link.Id, link.User.Id)
	return err
}

func (l Link) GetByID(link models.Link) (models.Link, error) {
	row := l.db.QueryRow("SELECT link_id, title, url, api_point, created_at from links where link_id=$1 and deleted_at is null", link.Id)
	result, err := l.Scan(row)
	return result, err
}

func (l Link) GetAllByUser(user models.UserInfo) ([]models.Link, error) {
	rows, err := l.db.Query("select link_id, title, url, api_point, link.created_at from links join user u on link.user_id = u.user_id where u.user_id=$1 and u.deleted_at is null", user.Id)
	if err != nil {
		return nil, err
	}
	return l.ScanRows(rows)
}

func (l Link) IsApiPointExist(str string) bool {
	var result int
	row := l.db.QueryRow("select count(*) from links where api_point=? ", str)
	err := row.Scan(&result)
	if err != nil {
		return false
	}
	return result != 0
}

func (l Link) GetByAPIPoint(link models.Link) (models.Link, error) {
	row := l.db.QueryRow("SELECT link_id, title, url, api_point, links.created_at from links where api_point=$1 and deleted_at is null ", link.ApiPoint)
	return l.Scan(row)
}
