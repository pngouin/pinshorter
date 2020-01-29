package db

import (
	"crypto/sha512"
	"database/sql"
	"encoding/base64"
	"golang.org/x/crypto/pbkdf2"
	"log"
)

func NewCrypt(sql *sql.DB) Crypt {
	c := Crypt{
		db: sql,
	}
	err := c.db.QueryRow("SELECT salt FROM params").Scan(&c.salt)
	if err != nil {
		return c
	}

	if c.salt == "" {
		log.Fatalln("Cannot retrieve salt for password: ", err)
	}

	return c
}

type Crypt struct {
	salt string
	db   *sql.DB
}

func (c Crypt) Hash(str string) string {
	return base64.StdEncoding.EncodeToString(pbkdf2.Key([]byte(str), []byte(c.salt), 100000, 64, sha512.New))
}
