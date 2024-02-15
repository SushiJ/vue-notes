package main

import (
	"encoding/json"
	"fmt"
	"io"
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
	r.Post("/", createNotes)
}

// TODO: uuid, createdAt and updatedAt should be db thingy
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
	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var note Notes
	if err := json.Unmarshal(data, &note); err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}
	note.Id = uuid.New()
	note.CreatedAt = time.Now()
	note.UpdatedAt = time.Now()

	marshal, err := json.Marshal(note)

	notes = append(notes, note)
	fmt.Fprint(w, string(marshal))
}
