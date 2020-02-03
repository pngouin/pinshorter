package main

import (
	"context"
	sqlHelp "github.com/ZooPin/pinshorter/db"
	"github.com/ZooPin/pinshorter/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
	"os/signal"
	"time"
)

var queryString string
var secret string

func init() {
	queryString = os.Getenv("DATABASE")
	secret = os.Getenv("PINSHORTER_SECRET")
}

func main() {
	if secret == "" {
		log.Fatalln("Environment variable PINSHORTER_SECRET not set")
	}
	if queryString == "" {
		log.Fatalln("Environment variable DATABASE not set")
	}
	e := echo.New()

	//Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db, err := sqlHelp.Open(queryString)
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer db.Close()

	link := handler.NewLink(db)
	user := handler.NewUser(db, secret)
    g := e.Group("/api")
	g.GET("/:api_point", link.Redirect)
	g.POST("auth", user.Login)

	p := g.Group("/link")
	p.Use(middleware.JWT([]byte(secret)))
	p.PUT("", link.Add)
	p.DELETE("/:api_point", link.Delete)
	p.GET("", link.List)

	// Start server
	go func() {
		if err := e.Start(":1323"); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}
