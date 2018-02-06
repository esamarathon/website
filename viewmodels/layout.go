package viewmodels

type meta struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Image       string `json:"image,omitempty"`
}

// DefaultMata is a set of default metadata values
var DefaultMeta = meta{
	"ESA Marathon",
	"Welcome to European Speedrunner Assembly!",
	"http://www.esamarathon.com/static/img/og-image.png",
}
