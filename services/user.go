package services

import (
	"database/sql"
	"github.com/ZooPin/pinshorter/db"
	"github.com/ZooPin/pinshorter/models"
)

func NewUser(sql *sql.DB) User {
	return User{db.NewUser(sql)}
}

type User struct {
	database db.User
}

func (u User) Connection(user models.UserConn) (bool, error) {
	return u.database.Connection(user)
}

func (u User) Add(conn models.UserConn) (models.UserInfo, error) {
	return u.database.Add(conn)
}

func (u User) GetById(user models.UserInfo) (models.UserInfo, error) {
	return u.database.GetById(user)
}

func (u User) IsUserExist(user models.UserInfo) bool {
	return u.database.IsUserExist(user)
}
