package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	username := "root"
	password := "hussain3660"
	host := "127.0.0.1"
	port := "3306"
	dbName := "ResearchPublicationTracker"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&clientFoundRows=true",
		username,
		password,
		host,
		port,
		dbName,
	)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Database ping error:", err)
	}

	fmt.Println("Connected to MySQL successfully")
}
