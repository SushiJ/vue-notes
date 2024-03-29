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
		fmt.Fprint(w, string("No notes found"))
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
		http.Error(w, "Invalid ID", http.StatusBadRequest)
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

	if len(note.Title) == 0 || len(note.Content) == 0 {
		fmt.Fprint(w, http.StatusBadRequest)
		return
	}

	note.Id = uuid.New()

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

	fmt.Println(fmt.Sprintf("%+v", string(marshal)))

	fmt.Fprint(w, string(marshal))
}

func (app *application) updateNote(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		app.errorLog.Print(w, err)
		return
	}

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

	if len(note.Title) == 0 || len(note.Content) == 0 {
		fmt.Fprint(w, http.StatusBadRequest)
		return
	}

	res, err := app.notes.Update(id, note.Title, note.Content)
	if err != nil || res == 0 {
		app.errorLog.Print(w, err)
		return
	}

	fmt.Fprint(w, http.StatusOK)
}

func (app *application) deleteNoteById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		app.errorLog.Print(w, err)
		return
	}

	err = app.notes.Delete(id)
	if err != nil {
		app.errorLog.Print(w, err)
		return
	}

	fmt.Fprint(w, http.StatusOK)
}
