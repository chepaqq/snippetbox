package models

import (
	"database/sql"
	"time"
)

// A Snippet represents data for an individual snippet
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// SnippetModel wraps a DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

// Insert new snippet into the database
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
    VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

// Get a specific snippet based on its id
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

// Latest 10 snippets
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
