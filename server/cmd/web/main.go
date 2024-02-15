package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sushij/vue-notes/server/internals/models"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	notes    *models.NoteModel
	router   *chi.Mux
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP port")
	dsn := flag.String("dsn", "./note.db", "sqlite source name")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	stmt := `CREATE TABLE IF NOT EXISTS notes (
    id TEXT NOT NULL PRIMARY KEY,
    title TEXT,
    content TEXT,
    created_at DATETIME default CURRENT_TIMESTAMP,
    updated_at DATETIME default CURRENT_TIMESTAMP
    );`

	if _, err := db.Exec(stmt); err != nil {
		errorLog.Fatal(err)
		return
	}

	r := chi.NewRouter()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		router:   r,
		notes:    &models.NoteModel{DB: db},
	}

	r.Route("/notes", app.routes)

	infoLog.Printf("Live at http://localhost%s", *addr)
	errorLog.Fatal((http.ListenAndServe(*addr, r)))
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
