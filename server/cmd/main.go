package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP port")
	flag.Parse()

	r := chi.NewRouter()
	r.Get("/check", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("200"))
	})

	r.Route("/notes", notesRoutes)

	log.Printf("Live at http://localhost%s", *addr)
	log.Fatal((http.ListenAndServe(*addr, r)))
}
