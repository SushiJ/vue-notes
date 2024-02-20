package main

import "github.com/go-chi/chi/v5"

func (app *application) routes(r chi.Router) {
	r.Get("/", app.getNotes)
	r.Post("/", app.createNotes)
	r.Get("/{id}", app.getNotesById)
	r.Delete("/{id}", app.getNotesById)
	r.Put("/{id}", app.updateNote)
}
