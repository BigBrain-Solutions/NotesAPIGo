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

	headerValue := r.Header.Get("Authorization")

	if services.IsAuthorized(headerValue) == false {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	claims := services.JwtParse(headerValue) // parsing JWT

	var note Note
	note.Created = time.Now()

	note.UserId = claims.Id // signing note.UserId from JWT payload Id

	// Redis
	_ = json.NewDecoder(r.Body).Decode(&note)

	notesArrJson := services.GetString("notes")

	json.Unmarshal([]byte(notesArrJson), &notes)

	if len(notes) == 0 {
		note.Id = 0
	} else {
		note.Id = notes[len(notes)-1].Id + 1
	}

	notes = append(notes, note)

	noteJsoned, err := json.Marshal(notes)

	if err != nil {
		println("cannot add note")
		return
	}

	services.SetString("notes", string(noteJsoned))
	w.WriteHeader(http.StatusCreated)
}

func GetNotes(w http.ResponseWriter, r *http.Request) {
	var notes []Note

	headerValue := r.Header.Get("Authorization")

	if services.IsAuthorized(headerValue) == false {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	notesArrJson := services.GetString("notes")

	json.Unmarshal([]byte(notesArrJson), &notes)

	json.NewEncoder(w).Encode(notes)
}

func GetNotesByUserId(w http.ResponseWriter, r *http.Request) {

	var notes []Note
	var userNotes []Note

	headerValue := r.Header.Get("Authorization")

	if services.IsAuthorized(headerValue) == false {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	notesArrJson := services.GetString("notes")

	json.Unmarshal([]byte(notesArrJson), &notes)

	for index, note := range notes {
		if note.UserId == services.JwtParse(headerValue).Id {
			userNotes = append(userNotes, notes[index])
		}
	}

	if userNotes == nil {
		w.WriteHeader(http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(userNotes)
}

func GetNote(w http.ResponseWriter, r *http.Request) {
	var notes []Note

	w.Header().Set("Content-Type", "application/json")

	headerValue := r.Header.Get("Authorization")
	if services.IsAuthorized(headerValue) == false {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	claims := services.JwtParse(headerValue)

	notesArrJson := services.GetString("notes")

	params := mux.Vars(r)

	json.Unmarshal([]byte(notesArrJson), &notes)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		return
	}

	for index, note := range notes {
		if note.Id == id && note.UserId == claims.Id {
			json.NewEncoder(w).Encode(notes[index])
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	w.WriteHeader(http.StatusBadRequest)
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	var notes []Note

	w.Header().Set("Content-Type", "application/json")

	headerValue := r.Header.Get("Authorization")
	if services.IsAuthorized(headerValue) == false {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	notesArrJson := services.GetString("notes")

	params := mux.Vars(r)

	json.Unmarshal([]byte(notesArrJson), &notes)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		return
	}

	claims := services.JwtParse(headerValue)

	for index, note := range notes {
		if note.Id == id && note.UserId == claims.Id {
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

	headerValue := r.Header.Get("Authorization")
	if services.IsAuthorized(headerValue) == false {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	notesArrJson := services.GetString("notes")

	params := mux.Vars(r)

	json.Unmarshal([]byte(notesArrJson), &notes)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		return
	}

	claims := services.JwtParse(headerValue)

	for index, note := range notes {
		if note.Id == id && note.UserId == claims.Id {

			notes = append(notes[:index], notes[index+1:]...)
			var note Note
			_ = json.NewDecoder(r.Body).Decode(&note)

			note.Id = id
			note.UserId = claims.Id
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
