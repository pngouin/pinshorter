package main

import (
	"context"
	"flag"
	sqlHelp "github.com/ZooPin/pinshorter/db"
	"github.com/ZooPin/pinshorter/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
	"os/signal"
	"time"
)

var sqlFile string
var dbName string
var secret string

func init() {
	flag.StringVar(&sqlFile, "sql", "create_postgres.sql", "Path to the sql file to create the database.")
	flag.StringVar(&dbName, "database", "db/database.db", "Path to the database.")
	secret = os.Getenv("PINSHORTER_SECRET")
	flag.Parse()
}

func main() {
	if secret == "" {
		log.Fatalln("Environment variable PINSHORTER_SECRET not set")
	}
	e := echo.New()

	//Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db, err := sqlHelp.Open(dbName)
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer db.Close()

	/*	c, err := ioutil.ReadFile(sqlFile)
		if err != nil {
			e.Logger.Fatal(err)
		}
		_, err = db.Exec(string(c))
		if err != nil {
			e.Logger.Fatal(err)
		}*/

	link := handler.NewLink(db)
	user := handler.NewUser(db, secret)

	e.GET("/:api_point", link.Redirect)
	e.POST("login", user.Login)

	p := e.Group("/link")
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
