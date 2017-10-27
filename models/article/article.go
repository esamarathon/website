package article

import (
	"time"
)

const table = "articles"

// Article describes the format of an article
type Article struct {
	ID        string `json:"_id,omitempty" gorethink:"id,omitempty"`
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
