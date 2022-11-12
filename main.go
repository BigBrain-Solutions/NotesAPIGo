package main

import (
	"NotesyAPI/config"
	"NotesyAPI/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// load configs
	godotenv.Load(".env")
	config.LoadConfig()

	// create a router
	r := mux.NewRouter()

	// define routes
	r.HandleFunc("/google_login", controller.GoogleLogin).Methods("GET")
	r.HandleFunc("/google_callback", controller.GoogleCallback).Methods("GET")

	r.HandleFunc("/note/{id}", controller.GetNote).Methods("GET")
	r.HandleFunc("/notes", controller.GetNotes).Methods("GET")
	r.HandleFunc("/note/{id}", controller.AddNote).Methods("PUT")

	r.HandleFunc("/note/{id}", controller.UpdateNote).Methods("PATCH")
	r.HandleFunc("/note/{id}", controller.DeleteNote).Methods("DELETE")

	// run server
	log.Println("started server on :: http://localhost:8080/")
	if oops := http.ListenAndServe(":8080", r); oops != nil {
		log.Fatal(oops)
	}
}
