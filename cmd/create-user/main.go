package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/ZooPin/pinshorter/db"
	"github.com/ZooPin/pinshorter/models"
	"log"
	"os"
	"strings"
	"time"
)

var dbName string

func init() {
	flag.StringVar(&dbName, "database", "", "Path to the database.")
	flag.Parse()
}

func main() {
	sql, err := db.Open(dbName)
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
	time.Sleep(10 * time.Second)
	log.Println("Try to connect with the credentials")
	userInfo, ok, err := user.Connection(userConn)
	if err != nil {
		log.Fatalln("Problem with the connection ", err)
	}
	if !ok {
		log.Fatalln("Credentials doesn't match...")
	}
	log.Println("Success !")
}
