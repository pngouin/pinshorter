package handler

import (
	"database/sql"
	"encoding/json"
	"github.com/pngouin/pinshorter/models"
	"github.com/pngouin/pinshorter/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"log"
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

	if user.Name == "" || user.Password == "" {
		jsonErr := models.Error{Error: "Input can't be empty."}
		return c.JSON(http.StatusBadRequest, jsonErr)
	}

	uInfo, ok, err := u.user.Connection(user)
	if err != nil {
		log.Println("Error: POST /auth user:", user.Name, "err:", err)
		return echo.ErrBadRequest
	}
	if !ok {
		return echo.ErrUnauthorized
	}

	claims := models.Token{
		Name: uInfo.Name,
		Id:   uInfo.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(48 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(u.secret))
	if err != nil {
		log.Println("Error signing token: ", err)
		return echo.ErrInternalServerError
	}
	log.Println("Info: POST /auth user:", uInfo.Id, "connected")

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
