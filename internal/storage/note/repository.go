package note

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/lilpipidron/sugar-backend/internal/models/notes"
	"github.com/lilpipidron/sugar-backend/internal/models/products"
)

type Repository interface {
	AddNote(note notes.Note, userID int64) error
	GetAllNotes(userID int64) ([]*notes.Note, error)
	GetNotesByDate(userID int64, dateTime time.Time) ([]*notes.Note, error)
	DeleteNote(noteID int64) error
}

type repository struct {
	DB *sql.DB
}

func NewNoteRepository(db *sql.DB) *repository {
	return &repository{DB: db}
}

func (db *repository) AddNote(note notes.Note, userID int64) error {
	const op = "storage.note.AddNote"

	query := "INSERT INTO note_header (note_type, create_date, sugar_level) VALUES (($1), ($2), ($3))"
	result, err := db.DB.Exec(query, note.NoteType, note.DateTime, note.SugarLevel)
	if err != nil {
		return fmt.Errorf("%s: failed add note in note_header: %w", op, err)
	}
	noteID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("%s: failed get noteID: %w", op, err)
	}
	query = "INSERT INTO note_user (note_id, user_id) VALUES ($1, $2)"
	_, err = db.DB.Exec(query, noteID, userID)
	if err != nil {
		return fmt.Errorf("%s: failed add note in note_user: %w", op, err)
	}

	return nil
}

func (db *repository) GetAllNotes(userID int64) ([]*notes.Note, error) {
	const op = "storage.note.GetAllNotes"

	query := "SELECT note_id FROM note_user WHERE user_id = $1"
	rows, err := db.DB.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("%s: failed get all notes: %w", op, err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println(fmt.Errorf("%s: failed close note's rows (note user): %w", op, err))
		}
	}(rows)
	var note []*notes.Note

	for rows.Next() {
		var noteID int64
		err = rows.Scan(&noteID)
		if err != nil {
			return nil, fmt.Errorf("%s: failed scan note's row (note user): %w", op, err)
		}
		query = "SELECT * FROM note_header WHERE note_id = $1"
		row, err := db.DB.Query(query, noteID)
		if err != nil {
			return nil, fmt.Errorf("%s: failed get note header: %w", op, err)
		}

		defer func(rows *sql.Rows) {
			err := rows.Close()
			if err != nil {
				log.Println(fmt.Errorf("%s: failed close note's row (note header): %w", op, err))
			}
		}(row)

		n := &notes.Note{}
		err = row.Scan(&n.NoteID, &n.NoteType, &n.DateTime, &n.SugarLevel)
		if err != nil {
			return nil, fmt.Errorf("%s: failed scan note's row (note header): %w", op, err)
		}

		var np []*notes.NoteProduct

		query = "SELECT * FROM note_detail where note_id = $1"
		pnRows, err := db.DB.Query(query, noteID)
		if err != nil {
			return nil, fmt.Errorf("%s: faileg get note detail: %w", op, err)
		}

		defer func(rows *sql.Rows) {
			err := rows.Close()
			if err != nil {
				log.Println(fmt.Errorf("%s: failed close note's rows (note detail): %w", op, err))
			}
		}(pnRows)

		for pnRows.Next() {
			var productID int64
			var productAmount int
			err = pnRows.Scan(&productID, &productAmount)
			if err != nil {
				return nil, fmt.Errorf("%s: failed scan note's rows (note detail): %w", op, err)
			}

			query = "SELECT * FROM products WHERE product_id = $1"
			pRow, err := db.DB.Query(query, productID)
			if err != nil {
				return nil, fmt.Errorf("%s: failed get product: %w", op, err)
			}

			defer func(rows *sql.Rows) {
				err := rows.Close()
				if err != nil {
					log.Println(fmt.Errorf("%s: failed close product row (note detail): %w", op, err))
				}
			}(pRow)

			p := &products.Product{}
			err = pRow.Scan(&p.ProductID, &p.Name, &p.Carbs)
			if err != nil {
				return nil, fmt.Errorf("%s: failed scan product's row (note detail): %w", op, err)
			}

			noteProduct := &notes.NoteProduct{Product: p, Amount: productAmount}
			np = append(np, noteProduct)
		}
		n.Products = np
		note = append(note, n)
	}

	return note, nil
}

func (db *repository) GetNotesByDate(userID int64, dateTime time.Time) ([]*notes.Note, error) {
	const op = "storage.note.GetAllNotes"

	query := "SELECT note_id FROM note_user WHERE user_id = $1"
	rows, err := db.DB.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("%s: failed get all notes: %w", op, err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println(fmt.Errorf("%s: failed close note's rows (note user): %w", op, err))
		}
	}(rows)
	var note []*notes.Note

	for rows.Next() {
		var noteID int64
		err = rows.Scan(&noteID)
		if err != nil {
			return nil, fmt.Errorf("%s: failed scan note's row (note user): %w", op, err)
		}
		query = "SELECT * FROM note_header WHERE note_id = $1 AND create_date = $2"
		row, err := db.DB.Query(query, noteID, dateTime)
		if err != nil {
			return nil, fmt.Errorf("%s: failed get note header: %w", op, err)
		}

		defer func(rows *sql.Rows) {
			err := rows.Close()
			if err != nil {
				log.Println(fmt.Errorf("%s: failed close note's row (note header): %w", op, err))
			}
		}(row)

		n := &notes.Note{}
		err = row.Scan(&n.NoteID, &n.NoteType, &n.DateTime, &n.SugarLevel)
		if err != nil {
			return nil, fmt.Errorf("%s: failed scan note's row (note header): %w", op, err)
		}

		var np []*notes.NoteProduct

		query = "SELECT * FROM note_detail where note_id = $1"
		pnRows, err := db.DB.Query(query, noteID)
		if err != nil {
			return nil, fmt.Errorf("%s: faileg get note detail: %w", op, err)
		}

		defer func(rows *sql.Rows) {
			err := rows.Close()
			if err != nil {
				log.Println(fmt.Errorf("%s: failed close note's rows (note detail): %w", op, err))
			}
		}(pnRows)

		for pnRows.Next() {
			var productID int64
			var productAmount int
			err = pnRows.Scan(&productID, &productAmount)
			if err != nil {
				return nil, fmt.Errorf("%s: failed scan note's rows (note detail): %w", op, err)
			}

			query = "SELECT * FROM products WHERE product_id = $1"
			pRow, err := db.DB.Query(query, productID)
			if err != nil {
				return nil, fmt.Errorf("%s: failed get product: %w", op, err)
			}

			defer func(rows *sql.Rows) {
				err := rows.Close()
				if err != nil {
					log.Println(fmt.Errorf("%s: failed close product row (note detail): %w", op, err))
				}
			}(pRow)

			p := &products.Product{}
			err = pRow.Scan(&p.ProductID, &p.Name, &p.Carbs)
			if err != nil {
				return nil, fmt.Errorf("%s: failed scan product's row (note detail): %w", op, err)
			}

			noteProduct := &notes.NoteProduct{Product: p, Amount: productAmount}
			np = append(np, noteProduct)
		}
		n.Products = np
		note = append(note, n)
	}

	return note, nil
}

func (db *repository) DeleteNote(noteID int64) error {
	const op = "storage.note.DeleteNote"

	query := "DELETE * FROM  note_user WHERE note_id = $1"
	_, err := db.DB.Exec(query, noteID)
	if err != nil {
		return fmt.Errorf("%s: failed delete note from note_user: %w", op, err)
	}

	query = "DELETE * FROM note_detail WHERE note_id = $1"
	_, err = db.DB.Exec(query, noteID)
	if err != nil {
		return fmt.Errorf("%s: failed delete note from note_detail: %w", op, err)
	}

	query = "DELETE * FROM note_header WHERE note_id = $1"
	_, err = db.DB.Exec(query, noteID)
	if err != nil {
		return fmt.Errorf("%s: failed delete note from note_header: %w", op, err)
	}

	return nil
}
