package models

import (
	"database/sql"
	"errors"
	"time"
)

type SnippetModelInterface interface {
	Insert(userID int, title string, content string, expires int) (int, error)
	Update(id int, userID int, title string, content string, expires int) error
	Delete(id int, userID int) error
	Get(id int) (Snippet, error)
	List() ([]Snippet, error)
	ListByUserID(userID int) ([]Snippet, error)
}

// Define a Snippet type to hold the data for an individual snippet. Notice how
// the fields of the struct correspond to the fields in our MySQL snippets
// table?
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
	UserId  int
}

// Define a SnippetModel type which wraps a sql.DB connection pool.
type SnippetModel struct {
	DB *sql.DB
}

// This will insert a new snippet into the database.
func (m *SnippetModel) Insert(userID int, title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets (user_id, title, content, created, expires)
    VALUES(?, ?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, userID, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *SnippetModel) Update(id int, userID int, title string, content string, expires int) error {
	stmt := `UPDATE snippets SET title = ?, content = ?, expires = DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY)
	WHERE id = ? AND user_id = ?`

	_, err := m.DB.Exec(stmt, title, content, expires, id, userID)
	return err
}

func (m *SnippetModel) Delete(id int, userID int) error {
	stmt := `DELETE FROM snippets WHERE id = ? AND user_id = ?`

	_, err := m.DB.Exec(stmt, id, userID)
	return err
}

// This will return a specific snippet based on its id.
func (m *SnippetModel) Get(id int) (Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
    WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := m.DB.QueryRow(stmt, id)

	var s Snippet

	if err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Snippet{}, ErrNoRecord
		} else {
			return Snippet{}, err
		}
	}

	return s, nil
}

// This will return the 10 most recently created snippets.
func (m *SnippetModel) List() ([]Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
    WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var snippets []Snippet

	for rows.Next() {
		var s Snippet

		if err := rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires); err != nil {
			return nil, err
		}

		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}

// This will return the 10 most recently created snippets.
func (m *SnippetModel) ListByUserID(userID int) ([]Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
    WHERE user_id = ? AND expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(stmt, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var snippets []Snippet

	for rows.Next() {
		var s Snippet

		if err := rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires); err != nil {
			return nil, err
		}

		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
