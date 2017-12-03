package viewmodels

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/esamarathon/website/cache"
	"github.com/esamarathon/website/config"
	"github.com/pkg/errors"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

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

type indexView struct {
	Meta          meta
	Frontpage     interface{}
	Livemode      bool
	CopyrightYear int
}

type frontPage struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
}

// Index returns the viewmodel for the indexview
func Index() indexView {
	// Convert from markdown to html
	view := indexView{
		Meta:          DefaultMeta,
		Livemode:      config.Config.LiveMode,
		CopyrightYear: time.Now().Year(),
	}

	// Attempt to withdraw cached frontpage
	frontpage, found := cache.Get("frontpage")
	if found {
		// Use found cache
		view.Frontpage = frontpage
	} else {
		// Withdraw frontpage data from file
		fPage := getFrontpage()
		// Parse body to markdown
		body := blackfriday.Run([]byte(fPage.Body), blackfriday.WithExtensions(config.Config.MarkdownExtensions))
		// Create map with interfaces due to caching
		view.Frontpage = map[string]interface{}{
			"Title": fPage.Title,
			"Body":  template.HTML(body),
		}
		// Cache result
		cache.Set("frontpage", view.Frontpage, cache.Duration())
	}
	return view
}

func getFrontpage() frontPage {

	fPage := frontPage{
		Title: DefaultMeta.Description,
		Body:  "Change this default text in the admin panel",
	}

	file, err := ioutil.ReadFile(config.Config.FrontpageDataPath)
	if err != nil {
		// Failed to read file, we should create the file and return the default
		log.Println(errors.Wrap(err, "Error reading frontpage data"))
		createFrontpage(fPage.Title, fPage.Body)
		return fPage
	}

	// Unmarshal to frontpage object
	err = json.Unmarshal(file, &fPage)
	if err != nil {
		log.Println(errors.Wrap(err, "Unmarshal frontpage"))
		// Something went wrong in unmarshaling
		// Suspecting messed up syntax or something similar
		// Try deleting the file and recreating it
		deleteFrontpage()
		createFrontpage(fPage.Title, fPage.Body)
		return fPage
	}
	return fPage

}

// deleteFrontpage deletes the data page for the frontpage
func deleteFrontpage() {
	err := os.Remove(config.Config.FrontpageDataPath)
	if err != nil {
		log.Println(errors.Wrap(err, "Error deleting frontpage data file"))
	}
}

// createFrontpage creates the data file if it doesn't exist
func createFrontpage(title, body string) {
	log.Println("Attempting to create frontpage data file")
	file, err := os.Create(config.Config.FrontpageDataPath)
	if err != nil {
		log.Println(errors.Wrap(err, "Error creating frontpage data file"))
	}
	defer file.Close()

	UpdateFrontpage(title, body)
}

// UpdateFrontpage takes a new title and body and marshals
// it to JSON before writing it to file
func UpdateFrontpage(title, body string) error {
	fPage := frontPage{title, body}

	jsonByte, _ := json.Marshal(fPage)
	err := ioutil.WriteFile(config.Config.FrontpageDataPath, jsonByte, 0644)
	if err != nil {
		log.Println(errors.Wrap(err, "Error writing to frontpage data file"))
	} else {
		// If we have cache
		_, found := cache.Get("frontpage")
		if found {
			// We delete it
			cache.Cache.Delete("frontpage")
		}
	}
	return err
}
