package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP port")
	flag.Parse()

	r := chi.NewRouter()
	r.Get("/check", checkHealth)

	r.Route("/notes", func(r chi.Router) {
		r.Get("/", getNotes)
		r.Get("/{id}", getNotesById)
		r.Post("/create", createNotes)
	})

	log.Printf("Live at http://localhost%s", *addr)
	log.Fatal((http.ListenAndServe(*addr, r)))
}

type Notes struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Content   string    `json:"content"`
}

func checkHealth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("200"))
}

var notes = []Notes{
	{Id: "1", Title: "First note", Content: "This is the first Note", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{Id: "2", Title: "Second note", Content: "This is the second Note", CreatedAt: time.Now(), UpdatedAt: time.Now()},
}

func getNotes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(notes)
}

func getNotesById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	fmt.Fprintf(w, "Id: %s", idParam)
}

func createNotes(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allow", http.StatusMethodNotAllowed)
	}
}
