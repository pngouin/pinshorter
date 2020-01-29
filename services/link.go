package services

import (
	"database/sql"
	"errors"
	"github.com/ZooPin/pinshorter/db"
	"github.com/ZooPin/pinshorter/models"
	"golang.org/x/net/html"
	"math/rand"
	"net/http"
	"time"
)

func NewLink(sql *sql.DB) Link {
	rand.Seed(time.Now().UnixNano())
	return Link{
		database: db.NewLink(sql),
		client:   http.Client{Timeout: 10 * time.Second},
	}
}

type Link struct {
	database db.Link
	client   http.Client
}

const (
	length  = 5
	digits  = "0123456789"
	letters = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	all     = digits + letters
)

func (l Link) Add(link models.Link) (models.Link, error) {
	resp, err := l.client.Get(link.URL)
	if err != nil {
		return link, err
	}

	if resp.StatusCode >= 400 {
		return link, errors.New("cannot GET the URL")
	}

	link.Title, err = l.getTitleFromBody(resp)
	if err != nil {
		return link, err
	}

	link.ApiPoint = l.createRandomApiPoint(length)
	link, err = l.database.Create(link)
	return link, err
}

func (l Link) GetByApiPoint(link models.Link) (models.Link, error) {
	return l.database.GetByAPIPoint(link)
}

func (l Link) Delete(link models.Link) error {
	return l.database.Delete(link)
}

func (l Link) GetById(link models.Link) (models.Link, error) {
	return l.database.GetByID(link)
}

func (l Link) GetAllFromUser(user models.UserInfo) ([]models.Link, error) {
	return l.database.GetAllByUser(user)
}

func (l Link) getTitleFromBody(data *http.Response) (string, error) {
	body, err := html.Parse(data.Body)
	if err != nil {
		return "", err
	}
	var title string
	var crawler func(node *html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.TextNode && node.Parent.Data == "title" && node.Parent.Type == html.ElementNode {
			title = node.Data
			return
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}

	crawler(body)
	if title == "" {
		return "", errors.New("missing <tittle> in the node tree")
	}
	return title, nil
}

func (l Link) createRandomApiPoint(length int) string {
	buf := make([]byte, length)
	buf[0] = digits[rand.Intn(len(digits))]
	buf[1] = letters[rand.Intn(len(letters))]
	for i := 2; i < length; i++ {
		buf[i] = all[rand.Intn(len(all))]
	}
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	if l.database.IsApiPointExist(string(buf)) {
		return l.createRandomApiPoint(length)
	}
	return string(buf)
}
