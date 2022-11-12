package controller

import (
	"NotesyAPI/services"
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
	UserId  string    `json:"UserId"`
}

func AddNote(w http.ResponseWriter, r *http.Request) {
	var notes []Note

	var note Note
	note.Created = time.Now()

	params := mux.Vars(r)
	note.UserId = params["id"]

	_ = json.NewDecoder(r.Body).Decode(&note)

	notesArrJson := services.GetString("notes")

	json.Unmarshal([]byte(notesArrJson), &notes)

	notes = append(notes, note)

	noteJsoned, err := json.Marshal(notes)

	if err != nil {
		println("cannot add note")
		return
	}

	services.SetString("notes", string(noteJsoned))
}

func GetNotes(w http.ResponseWriter, r *http.Request) {
	var notes []Note

	w.Header().Set("Content-Type", "application/json")

	notesArrJson := services.GetString("notes")

	json.Unmarshal([]byte(notesArrJson), &notes)

	json.NewEncoder(w).Encode(notes)
}

func GetNotesByUserId(w http.ResponseWriter, r *http.Request) {

	var notes []Note
	var userNotes []Note

	w.Header().Set("Content-Type", "application/json")

	notesArrJson := services.GetString("notes")

	params := mux.Vars(r)

	json.Unmarshal([]byte(notesArrJson), &notes)

	for index, note := range notes {
		if note.UserId == params["id"] {
			userNotes = append(userNotes, notes[index])
		}
	}

	json.NewEncoder(w).Encode(userNotes)
}

func GetNote(w http.ResponseWriter, r *http.Request) {
	var notes []Note

	w.Header().Set("Content-Type", "application/json")

	notesArrJson := services.GetString("notes")

	params := mux.Vars(r)

	json.Unmarshal([]byte(notesArrJson), &notes)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		return
	}

	for index, note := range notes {
		if note.Id == id {
			json.NewEncoder(w).Encode(notes[index])
		}
	}

}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	var notes []Note

	w.Header().Set("Content-Type", "application/json")

	notesArrJson := services.GetString("notes")

	params := mux.Vars(r)

	json.Unmarshal([]byte(notesArrJson), &notes)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		return
	}

	for index, note := range notes {
		if note.Id == id {
			notes = append(notes[:index], notes[index+1:]...)

			notesJsoned, err := json.Marshal(notes)

			if err != nil {
				break
			}

			services.SetString("notes", string(notesJsoned))
			break
		}
	}

}

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	var notes []Note

	w.Header().Set("Content-Type", "application/json")

	notesArrJson := services.GetString("notes")

	params := mux.Vars(r)

	json.Unmarshal([]byte(notesArrJson), &notes)

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

			notesJsoned, err := json.Marshal(notes)

			if err != nil {
				break
			}

			services.SetString("notes", string(notesJsoned))
			break
		}
	}
}
