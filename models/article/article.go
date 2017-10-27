package article

import (
	"time"
)

// Article describes the format of an article
type Article struct {
	ID        string    `gorethink:"id" json:"id,omitempty"`
	Title     string    `gorethink:"title" json:"title,omitempty"`
	Body      string    `gorethink:"body" json:"body,omitempty"`
	CreatedAt time.Time `gorethink:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorethink:"updated_at" json:"updated_at,omitempty"`
}
