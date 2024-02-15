package main

import "github.com/go-chi/chi/v5"

func (app *application) routes(r chi.Router) {
	r.Get("/", app.getNotes)
	r.Get("/{id}", app.getNotesById)
	r.Post("/", app.createNotes)
}
