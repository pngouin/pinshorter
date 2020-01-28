package handler

import (
	"database/sql"
	"encoding/json"
	"github.com/ZooPin/pinshorter/models"
	"github.com/ZooPin/pinshorter/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

func NewLink(sql *sql.DB) Link {
	return Link{
		link:  services.NewLink(sql),
		count: services.NewCount(sql),
		user:  services.NewUser(sql),
	}
}

type Link struct {
	link  services.Link
	count services.Count
	user  services.User
}

// GET /:api_point redirect to the URL
func (l Link) Redirect(c echo.Context) error {
	link := models.Link{
		ApiPoint: c.Param("api_point"),
	}
	link, err := l.link.GetByApiPoint(link)
	if err != nil {
		jsonErr := models.Error{Error: err.Error()}
		return c.JSON(http.StatusBadRequest, jsonErr)
	}
	return c.Redirect(http.StatusPermanentRedirect, link.URL)
}

// POST /add to create a short link !secure
func (l Link) Add(c echo.Context) error {
	var link models.Link

	err := json.NewDecoder(c.Request().Body).Decode(&link)
	if err != nil {
		jsonErr := models.Error{Error: err.Error()}
		return c.JSON(http.StatusBadRequest, jsonErr)
	}

	user := getUserJWT(c)
	if !l.user.IsUserExist(user) {
		jsonErr := models.Error{Error: errUserDontExist}
		return c.JSON(http.StatusBadRequest, jsonErr)
	}
	link.User = user

	link, err = l.link.Add(link)
	if err != nil {
		jsonErr := models.Error{Error: err.Error()}
		return c.JSON(http.StatusBadRequest, jsonErr)
	}

	return c.JSON(http.StatusCreated, link)
}

// GET /list list all link by an user !secure
func (l Link) List(c echo.Context) error {
	user := getUserJWT(c)
	if !l.user.IsUserExist(user) {
		jsonErr := models.Error{Error: errUserDontExist}
		return c.JSON(http.StatusBadRequest, jsonErr)
	}
	links, err := l.link.GetAllFromUser(user)
	if err != nil {
		jsonErr := models.Error{Error: err.Error()}
		return c.JSON(http.StatusInternalServerError, jsonErr)
	}
	return c.JSON(http.StatusOK, links)
}

// DELETE /:id a link !secure
func (l Link) Delete(c echo.Context) error {
	user := getUserJWT(c)
	if !l.user.IsUserExist(user) {
		jsonErr := models.Error{Error: errUserDontExist}
		return c.JSON(http.StatusBadRequest, jsonErr)
	}

	var link models.Link
	link.Id = c.Param("id")
	link.User = user

	err := l.link.Delete(link)
	if err != nil {
		jsonErr := models.Error{Error: err.Error()}
		return c.JSON(http.StatusInternalServerError, jsonErr)
	}

	return c.NoContent(http.StatusOK)
}
