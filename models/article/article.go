package article

import (
	"html/template"

	"github.com/olenedr/esamarathon/config"
	"github.com/olenedr/esamarathon/models/user"
	blackfriday "gopkg.in/russross/blackfriday.v2"

	"time"
)

const table = "articles"

// Article describes the format of an article
type Article struct {
	ID        string        `json:"_id,omitempty" gorethink:"id,omitempty"`
	Title     string        `json:"title,omitempty" gorethink:"title,omitempty"`
	Body      string        `json:"body,omitempty" gorethink:"body,omitempty"`
	HTML      template.HTML `json:"html,omitempty" gorethink:"html,omitempty"`
	Published bool          `json:"published,omitempty" gorethink:"published,omitempty"`
	Timestamp string        `json:"timestamp,omitempty" gorethink:"timestamp,omitempty"`
	Authors   []user.User   `json:"authors,omitempty" gorethink:"authors,omitempty"`
	CreatedAt time.Time     `json:"created_at,omitempty" gorethink:"created_at,omitempty"`
	UpdatedAt time.Time     `json:"updated_at,omitempty" gorethink:"updated_at,omitempty"`
}

// AuthorExists checks if a user is in the author-slice
func (a *Article) AuthorExists(user user.User) bool {
	for _, u := range a.Authors {
		if u.ID == user.ID {
			return true
		}
	}

	return false
}

// AddAuthor appends an author to the article
func (a *Article) AddAuthor(u user.User) {
	if a.Authors == nil {
		a.Authors = []user.User{}
	}

	a.Authors = append(a.Authors, u)
}

// ShortenBody shaves off some of the body to make an appropriate ingress
func (a *Article) ShortenBody() {
	if len(a.Body) >= 340 {
		a.Body = a.Body[0:340] + "..."
	}
}

// ParseHTML parses the markdown to HTML
func (a *Article) ParseHTML() {
	markdown := []byte(a.Body)
	html := blackfriday.Run(markdown, blackfriday.WithExtensions(config.Config.MarkdownExtensions))
	a.HTML = template.HTML(string(html))
}

// FormatTimestamp adds a UTC timestamp to the article
func (a *Article) FormatTimestamp() {
	a.Timestamp = a.CreatedAt.UTC().String()
}
