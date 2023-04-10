package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Notes struct {
	Id        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Content   string    `json:"content"`
}

func notesRoutes(r chi.Router) {
	r.Get("/", getNotes)
	r.Get("/{id}", getNotesById)
	r.Post("/create", createNotes)
}

var notes = []Notes{
	{Id: uuid.New(), Title: "First note", Content: "This is the first Note", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{Id: uuid.New(), Title: "Second note", Content: "This is the second Note", CreatedAt: time.Now(), UpdatedAt: time.Now()},
}

func getNotes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(notes)
}

func getNotesById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for _, n := range notes {
		if n.Id == id {
			json.NewEncoder(w).Encode(n)
			return
		}
	}
}

func createNotes(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allow", http.StatusMethodNotAllowed)
	}
}
