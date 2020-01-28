package handler

import (
	"database/sql"
	"encoding/json"
	"github.com/ZooPin/pinshorter/models"
	"github.com/ZooPin/pinshorter/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func NewUser(sql *sql.DB, str string) User {
	return User{
		user:   services.NewUser(sql),
		secret: str,
	}
}

type User struct {
	user   services.User
	secret string
}

func (u User) Login(c echo.Context) error {
	var user models.UserConn

	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		jsonErr := models.Error{Error: err.Error()}
		return c.JSON(http.StatusBadRequest, jsonErr)
	}

	ok, err := u.user.Connection(user)
	if err != nil {
		jsonErr := models.Error{Error: err.Error()}
		return c.JSON(http.StatusBadRequest, jsonErr)
	}
	if !ok {
		return echo.ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS512)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.Id
	claims["name"] = user.Name
	claims["exp"] = time.Now().Add(72 * time.Hour)

	t, err := token.SignedString([]byte(u.secret))
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
