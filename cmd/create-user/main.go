package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/pngouin/pinshorter/db"
	"github.com/pngouin/pinshorter/models"
	"log"
	"os"
	"strings"
)

var queryString string

func init() {
	queryString = os.Getenv("DATABASE_URL")
	flag.Parse()
}

func main() {
	if queryString == "" {
		log.Fatalln("Environment variable DATABASE_URL not set")
	}
	sql, err := db.Open(queryString)
	if err != nil {
		log.Fatalln("Cannot open the database: ", err)
	}

	user := db.NewUser(sql)
	var userConn models.UserConn
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Username: ")
	in, _ := reader.ReadString('\n')
	userConn.Name = strings.TrimSuffix(in, "\n")
	fmt.Print("Password: ")
	in, _ = reader.ReadString('\n')
	userConn.Password = strings.TrimSuffix(in, "\n")
	userInfo, err := user.Add(userConn)
	if err != nil {
		log.Fatalln("Error creating user: ", err)
	}
	log.Println("User created with name ", userInfo.Name, " and id ", userInfo.Id)
}
