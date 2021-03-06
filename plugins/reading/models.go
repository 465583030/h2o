package reading

import (
	"time"

	"github.com/kapmahc/h2o/plugins/auth"
	"github.com/kapmahc/h2o/web"
)

// Book book
type Book struct {
	web.Model

	Author      string    `json:"author"`
	Publisher   string    `json:"publisher"`
	Title       string    `json:"title"`
	Type        string    `json:"type"`
	Lang        string    `json:"lang"`
	File        string    `json:"-"`
	Subject     string    `json:"subject"`
	Description string    `json:"description"`
	PublishedAt time.Time `json:"publishedAt"`
	Cover       string    `json:"cover"`
}

// TableName table name
func (Book) TableName() string {
	return "reading_books"
}

// Note note
type Note struct {
	web.Model
	Type   string    `json:"type"`
	Body   string    `json:"body"`
	UserID uint      `json:"userId"`
	User   auth.User `json:"-"`
	BookID uint      `json:"bookId"`
	Book   Book      `json:"book"`
}

// TableName table name
func (Note) TableName() string {
	return "reading_notes"
}
