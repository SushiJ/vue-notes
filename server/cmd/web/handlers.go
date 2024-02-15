package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Notes struct {
	Id      uuid.UUID `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
}

func (app *application) getNotes(w http.ResponseWriter, r *http.Request) {

	notes, err := app.notes.GetAll()
	if err != nil {
		app.errorLog.Print(w, err)
		return
	}

	if len(notes) == 0 {
		fmt.Fprint(w, []byte("No notes found"))
		return
	}

	notesJson, err := json.Marshal(notes)
	if err != nil {
		app.errorLog.Print(w, err)
		return
	}
	fmt.Fprint(w, string(notesJson))
}

func (app *application) getNotesById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		app.errorLog.Print(w, err)
		// http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	note, err := app.notes.Get(id)
	if err != nil {
		app.errorLog.Print(w, err)
		return
	}

	noteJson, err := json.Marshal(note)
	fmt.Fprint(w, string(noteJson))
}

func (app *application) createNotes(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		app.errorLog.Print(err)
	}

	var note Notes

	if err := json.Unmarshal(data, &note); err != nil {
		app.errorLog.Print(err)
		return
	}

	note.Id = uuid.New()

	fmt.Println(fmt.Sprintf("%+v", note))

	if err := app.notes.Insert(note.Id, note.Title, note.Content); err != nil {
		fmt.Printf("%+v", err)
		app.errorLog.Print(w, err)
		return
	}

	marshal, err := json.Marshal(note)
	if err != nil {
		app.errorLog.Print(w, err)
		return
	}

	fmt.Fprint(w, string(marshal))
}
