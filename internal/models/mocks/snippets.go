package mocks

import (
	"time"

	"github.com/johanpham2711/snippet-box/internal/models"
)

var mockSnippet = models.Snippet{
	ID:      1,
	Title:   "An old silent pond",
	Content: "An old silent pond...",
	Created: time.Now(),
	Expires: time.Now(),
}

type SnippetModel struct{}

func (m *SnippetModel) Insert(userID int, title string, content string, expires int) (int, error) {
	return 2, nil
}

func (m *SnippetModel) Update(id int, userID int, title string, content string, expires int) error {
	return nil
}

func (m *SnippetModel) Delete(id int, userID int) error {
	return nil
}

func (m *SnippetModel) Get(id int) (models.Snippet, error) {
	switch id {
	case 1:
		return mockSnippet, nil
	default:
		return models.Snippet{}, models.ErrNoRecord
	}
}

func (m *SnippetModel) List() ([]models.Snippet, error) {
	return []models.Snippet{mockSnippet}, nil
}

func (m *SnippetModel) ListByUserID(userID int) ([]models.Snippet, error) {
	return []models.Snippet{mockSnippet}, nil
}