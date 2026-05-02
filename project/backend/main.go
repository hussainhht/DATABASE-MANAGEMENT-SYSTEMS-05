package main

import (
	"log"
	"net/http"

	"RPT/handlers"
)

func main() {
	ConnectDB()
	defer DB.Close()

	fs := http.FileServer(http.Dir("../frontend"))
	http.Handle("/", fs)

	handlers.RegisterRoutes(DB)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
