package handler

import (
	"github.com/ZooPin/pinshorter/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

const (
	errUserDontExist = "User don't exist."
)

func getUserJWT(c echo.Context) models.UserInfo {
	var user models.UserInfo

	t := c.Get("id").(*jwt.Token)
	claims := t.Claims.(jwt.MapClaims)
	user.Id = claims["id"].(string)
	user.Name = claims["name"].(string)

	return user
}
