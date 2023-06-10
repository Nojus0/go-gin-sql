package Mysql

import (
	"database/sql"
	"log"
	"os"
	"time"
)

var dbUrl = os.Getenv("DSN")

var Session *sql.DB

func Connect() error {
	db, err := sql.Open("mysql", dbUrl)

	if err != nil {
		return err
	}

	t := time.Now()
	err = db.Ping()

	if err != nil {
		return err
	}

	log.Println("Connected to MySQL Database with ping time:", time.Since(t))

	return nil
}
