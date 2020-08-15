package main

import (
	"context"
	"flag"
	sqlHelp "github.com/pngouin/pinshorter/db"
	"github.com/pngouin/pinshorter/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
	"os/signal"
	"time"
)

var queryString string
var secret string
var port string
var dev bool

func init() {
	queryString = os.Getenv("DATABASE_URL")
	secret = os.Getenv("PINSHORTER_SECRET")
	port = os.Getenv("PORT")
	flag.BoolVar(&dev, "dev", false, "Dev configuration server.")
	flag.Parse()
}

func main() {
	if secret == "" {
		log.Fatalln("Environment variable PINSHORTER_SECRET not set")
	}
	if queryString == "" {
		log.Fatalln("Environment variable DATABASE_URL not set")
	}
	e := echo.New()

	//Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	if dev {
		e.Use(middleware.CORS())
	}

	db, err := sqlHelp.Open(queryString)
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer db.Close()

	e.Static("/user", "www/dist")
	link := handler.NewLink(db)
	user := handler.NewUser(db, secret)
	e.GET("/:api_point", link.Redirect)
	g := e.Group("/api")
	g.POST("/auth", user.Login)

	p := g.Group("/link")
	p.Use(middleware.JWT([]byte(secret)))
	p.PUT("", link.Add)
	p.DELETE("/:id", link.Delete)
	p.GET("", link.List)

	// Start server
	go func() {
		if err := e.Start(":" + port); err != nil {
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
