package main

import (
	"NotesyAPI/config"
	"NotesyAPI/controller"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	// load configs
	godotenv.Load(".env")
	config.LoadConfig()

	// create a router
	mux := http.NewServeMux()

	// define routes
	mux.HandleFunc("/google_login", controller.GoogleLogin)
	mux.HandleFunc("/google_callback", controller.GoogleCallback)

	// run server
	log.Println("started server on :: http://localhost:8080/")
	if oops := http.ListenAndServe(":8080", mux); oops != nil {
		log.Fatal(oops)
	}
}
