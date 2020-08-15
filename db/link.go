package db

import (
	"database/sql"
	"github.com/pngouin/pinshorter/db/query"
	"github.com/pngouin/pinshorter/models"
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
		&result.User.Id,
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
			&result.User.Id,
			&result.User.Name,
			&result.Count,
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
	_, err := l.db.Exec(query.LinkCreate,
		link.Id, link.Title, link.URL, link.ApiPoint, link.User.Id)
	return link, err
}

func (l Link) Delete(link models.Link) error {
	_, err := l.db.Exec(query.LinkDelete, link.Id, link.User.Id)
	return err
}

func (l Link) GetByID(link models.Link) (models.Link, error) {
	row := l.db.QueryRow(query.LinkById, link.Id)
	result, err := l.Scan(row)
	return result, err
}

func (l Link) GetAllByUser(user models.UserInfo) ([]models.Link, error) {
	rows, err := l.db.Query(query.LinkAllByUser, user.Id)
	if err != nil {
		return nil, err
	}
	return l.ScanRows(rows)
}

func (l Link) IsApiPointExist(str string) bool {
	var result int
	row := l.db.QueryRow(query.LinkIsApiPointExist, str)
	err := row.Scan(&result)
	if err != nil {
		return false
	}
	return result != 0
}

func (l Link) GetByAPIPoint(link models.Link) (models.Link, error) {
	row := l.db.QueryRow(query.LinkByAPIPoint, link.ApiPoint)
	return l.Scan(row)
}
