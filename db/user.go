package db

import (
	"database/sql"
	"github.com/ZooPin/pinshorter/models"
)

func NewUser(db *sql.DB) User {
	return User{
		db:    db,
		crypt: NewCrypt(db),
	}
}

type User struct {
	db    *sql.DB
	crypt Crypt
}

func (u User) ScanConnection(row *sql.Row) (models.UserConn, error) {
	var result models.UserConn
	err := row.Scan(&result.Id, &result.Name, &result.Password)
	return result, err
}

func (u User) ScanInfo(row *sql.Row) (models.UserInfo, error) {
	var result models.UserInfo
	err := row.Scan(&result.Id, &result.Name)
	return result, err
}

func (u User) Add(user models.UserConn) (models.UserInfo, error) {
	user.Id = createUUID()
	_, err := u.db.Exec("INSERT INTO user (user_id, name, password, created_at) VALUES (?, ?, ?, cast(datetime('now') as text))",
		user.Id, user.Name, u.crypt.Hash(user.Password))
	return user.ToUserInfo(), err
}

func (u User) Connection(conn models.UserConn) (models.UserInfo, bool, error) {
	var pass int
	row := u.db.QueryRow("SELECT user_id, count(*) from user where name=? and password=? and deleted_at is null", conn.Name, u.crypt.Hash(conn.Password))
	err := row.Scan(&conn.Id, &pass)
	return conn.ToUserInfo(), pass != 0, err
}

func (u User) GetById(user models.UserInfo) (models.UserInfo, error) {
	row := u.db.QueryRow("SELECT user_id, name FROM user WHERE user_id=? AND deleted_at IS NULL", user.Id)
	return u.ScanInfo(row)
}

func (u User) IsUserExist(user models.UserInfo) bool {
	var result int
	row := u.db.QueryRow("select count(*) from user where user_id=? and name=? and deleted_at is null", user.Id, user.Name)
	err := row.Scan(&result)
	if err != nil {
		return false
	}
	return result != 0
}
