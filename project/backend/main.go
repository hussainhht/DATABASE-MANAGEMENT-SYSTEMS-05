package main

import (
	"log"
	"net/http"
	"os"
	_ "github.com/go-sql-driver/mysql"

	"RPT/handlers"
)

func main() {
	ConnectDB()
	defer DB.Close()

	fs := http.FileServer(http.Dir("../frontend"))
	http.Handle("/", fs)

	handlers.RegisterRoutes(DB)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
