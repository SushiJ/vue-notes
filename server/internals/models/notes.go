package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Note struct {
	Id        uuid.UUID
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
	Content   string
}

type NoteModel struct {
	DB *sql.DB
}

var ErrNoRecord = errors.New("models: no matching record found")

// TODO: uuid, createdAt and updatedAt should be db thingy
func (m *NoteModel) Insert(id uuid.UUID, title, content string) error {

	stmt := `INSERT INTO notes ( id, title, content ) VALUES ( ?, ?, ? )`

	_, err := m.DB.Exec(stmt, id, title, content)
	if err != nil {
		return err
	}

	return nil
}

func (m *NoteModel) Get(id uuid.UUID) (*Note, error) {
	stmt := `SELECT id,title,content,created_at FROM notes WHERE id = ?`
	row := m.DB.QueryRow(stmt, id)
	n := &Note{}

	err := row.Scan(&n.Id, &n.Title, &n.Content, &n.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}
	return n, nil
}

func (m *NoteModel) GetAll() ([]*Note, error) {
	stmt := `SELECT id,title,content,created_at,updated_at FROM notes`
	rows, err := m.DB.Query(stmt)

	defer rows.Close()
	notes := []*Note{}

	for rows.Next() {
		n := &Note{}

		err := rows.Scan(&n.Id, &n.Title, &n.Content, &n.CreatedAt, &n.UpdatedAt)
		if err != nil {
			return nil, err
		}
		notes = append(notes, n)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return notes, nil
}
