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

	r.HandleFunc("/note/{id}", controller.GetNote).Methods(http.MethodGet, http.MethodOptions)      // get note by note Id
	r.HandleFunc("/notes/all", controller.GetNotes).Methods(http.MethodGet, http.MethodOptions)     // get all notes
	r.HandleFunc("/notes", controller.GetNotesByUserId).Methods(http.MethodGet, http.MethodOptions) // get all notes by user Id from JWT
	r.HandleFunc("/note", controller.AddNote).Methods(http.MethodPut, http.MethodOptions)           // add note by user Id from JWT

	r.HandleFunc("/note/{id}", controller.UpdateNote).Methods(http.MethodPatch, http.MethodOptions)  // update note by note Id
	r.HandleFunc("/note/{id}", controller.DeleteNote).Methods(http.MethodDelete, http.MethodOptions) // remove note by note Id

	// run server
	log.Println("started server on :: http://localhost:8080/")
	if oops := http.ListenAndServe(":8080", r); oops != nil {
		log.Fatal(oops)
	}
}
