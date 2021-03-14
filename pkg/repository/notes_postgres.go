package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/msvetkov/notes-app/pkg/domain"
	"strings"
)

type NotesPostgres struct {
	db *sqlx.DB
}

func NewNotesPostgres(db *sqlx.DB) *NotesPostgres {
	return &NotesPostgres{db: db}
}

func (r *NotesPostgres) Create(note domain.Note) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createNoteQuery := fmt.Sprintf(`INSERT INTO %s (title, body, date_created, user_id) values ($1, $2, $3, $4) RETURNING id`, notesTable)
	row := tx.QueryRow(createNoteQuery, note.Title, note.Body, note.DateCreated, note.UserId)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *NotesPostgres) GetAll(userId int) ([]domain.Note, error) {
	var list []domain.Note

	getNotesQuery := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", notesTable)
	err := r.db.Select(&list, getNotesQuery, userId)

	return list, err
}

func (r *NotesPostgres) GetById(userId, noteId int) (domain.Note, error) {
	var note domain.Note

	getNoteQuery := fmt.Sprintf("SELECT * FROM %s WHERE id = $1 AND user_id = $2", notesTable)
	err := r.db.Get(&note, getNoteQuery, noteId, userId)

	return note, err
}

func (r *NotesPostgres) Delete(userId, noteId int) error {
	deleteQuery := fmt.Sprintf("DELETE FROM %s WHERE id = $1 AND user_id = $2", notesTable)
	_, err := r.db.Exec(deleteQuery, noteId, userId)

	return err
}

func (r *NotesPostgres) Update(userId, noteId int, input domain.UpdateNoteInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}
	if input.Body != nil {
		setValues = append(setValues, fmt.Sprintf("body=$%d", argId))
		args = append(args, *input.Body)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d AND user_id = $%d", notesTable, setQuery, argId, argId+1)
	args = append(args, noteId, userId)

	_, err := r.db.Exec(query, args...)
	return err
}
