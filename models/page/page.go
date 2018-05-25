package page

import (
	"github.com/esamarathon/website/models/article"
)

const table = "pages"

type Page struct {
	article.Article
	FriendlyName string `json:"friendly_name,omitempty" gorethink:"friendly_name,omitempty"`
}