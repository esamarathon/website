package article

import (
	"github.com/olenedr/esamarathon/models/user"

	"time"
)

const table = "articles"

// Article describes the format of an article
type Article struct {
	ID        string      `json:"_id,omitempty" gorethink:"id,omitempty"`
	Title     string      `json:"title,omitempty" gorethink:"title,omitempty"`
	Body      string      `json:"body,omitempty" gorethink:"body,omitempty"`
	Authors   []user.User `json:"authors,omitempty" gorethink:"authors,omitempty"`
	CreatedAt time.Time   `json:"created_at,omitempty" gorethink:"created_at,omitempty"`
	UpdatedAt time.Time   `json:"updated_at,omitempty" gorethink:"updated_at,omitempty"`
}

func (a *Article) AuthorExists(user user.User) bool {
	for _, u := range a.Authors {
		if u.ID == user.ID {
			return true
		}
	}

	return false
}
