package handler

import (
	"database/sql"
	"encoding/json"
	"github.com/ZooPin/pinshorter/models"
	"github.com/ZooPin/pinshorter/services"
	"github.com/labstack/echo/v4"
	"log"
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
	if link.ApiPoint == "" || len(link.ApiPoint) > services.ApiLength {
		return echo.ErrNotFound
	}
	link, err := l.link.GetByApiPoint(link)
	if err != nil {
		log.Println("Error: GET /", link.ApiPoint, "err:", err)
		return echo.ErrBadRequest
	}
	defer l.link.AddCountToLink(link)
	return c.Redirect(http.StatusPermanentRedirect, link.URL)
}

// PUT /add to create a short link !secure
func (l Link) Add(c echo.Context) error {
	var link models.Link

	user := getUserJWT(c)
	if !l.user.IsUserExist(user) {
		log.Println("Error: PUT /add err: user don't exist id:", user.Id, "name:", user.Name)
		return echo.ErrUnauthorized
	}
	link.User = user

	err := json.NewDecoder(c.Request().Body).Decode(&link)
	if err != nil {
		log.Println("Error: PUT /add User:", user.Id, "err:", err)
		return echo.ErrBadRequest
	}

	link, err = l.link.Add(link)
	if err != nil {
		log.Println("Error: PUT /add User:", user.Id, "err:", err)
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, link)
}

// GET /list list all link by an user !secure
func (l Link) List(c echo.Context) error {
	user := getUserJWT(c)
	if !l.user.IsUserExist(user) {
		log.Println("Error: GET /list", c.Param("id"), " err: user don't exist id:", user.Id, "name:", user.Name)
		return echo.ErrUnauthorized
	}
	links, err := l.link.GetAllFromUser(user)
	if err != nil {
		log.Println("Error: PUT /add User:", user.Id, "err:", err)
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, links)
}

// DELETE /:id a link !secure
func (l Link) Delete(c echo.Context) error {
	user := getUserJWT(c)
	if !l.user.IsUserExist(user) {
		log.Println("Error: DELETE /", c.Param("id"), " err: user don't exist id:", user.Id, "name:", user.Name)
		return echo.ErrUnauthorized
	}

	var link models.Link
	link.Id = c.Param("id")
	link.User = user

	err := l.link.Delete(link)
	if err != nil {
		log.Println("Error: DELETE /", link.Id, "user:", user.Id, "err:", err)
		return echo.ErrInternalServerError
	}
	log.Println("Info: DELETE /", link.Id, "user:", user.Id, "deleted")

	return c.NoContent(http.StatusOK)
}
