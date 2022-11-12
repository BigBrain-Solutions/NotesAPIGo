package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Note struct {
	Id      int       `json:"Id"`
	Title   string    `json:"Title"`
	Content string    `json:"Content"`
	Created time.Time `json:"CreatedAt"`

	UserId string `json:"UserId"`
}

var notes []Note

func AddNote(w http.ResponseWriter, r *http.Request) {
	var note Note
	note.Created = time.Now()

	params := mux.Vars(r)
	note.UserId = params["id"]

	_ = json.NewDecoder(r.Body).Decode(&note)
	notes = append(notes, note)
}

func GetNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

func GetNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for i := 0; i < len(notes); i++ {
		if notes[i].UserId == params["id"] {
			json.NewEncoder(w).Encode(notes[i])
		}
	}

}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		return
	}

	for index, note := range notes {
		if note.Id == id {
			notes = append(notes[:index], notes[index+1:]...)
			break
		}
	}

}

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		return
	}

	for index, note := range notes {
		if note.Id == id {

			notes = append(notes[:index], notes[index+1:]...)
			var note Note
			_ = json.NewDecoder(r.Body).Decode(&note)

			note.Id = id
			notes = append(notes, note)

			json.NewEncoder(w).Encode(note)
		}
	}
}
