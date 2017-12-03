package article

import (
	"math"
	"time"

	"github.com/esamarathon/website/config"
	"github.com/esamarathon/website/db"
	"github.com/pkg/errors"
	r "gopkg.in/gorethink/gorethink.v3"
)

// Create inserts a new article entry into the database
func (a *Article) Create() error {
	data := map[string]interface{}{
		"title":      a.Title,
		"body":       a.Body,
		"authors":    a.Authors,
		"published":  a.Published,
		"created_at": time.Now(),
		"updated_at": time.Now(),
	}

	_, err := db.Insert(table, data)
	return err
}

// All returns a slice containing all the articles
func All() ([]Article, error) {
	rows, err := db.GetAllByOrder(table, "created_at")
	var a []Article
	if err != nil {
		return a, errors.Wrap(err, "article.All")
	}

	if err = rows.All(&a); err != nil {
		return a, errors.Wrap(err, "article.All")
	}

	return a, nil
}

// Page returns the articles of a given page
func Page(page int, published bool) ([]Article, error) {
	var rows *r.Cursor
	var err error
	if published {
		filter := map[string]interface{}{
			"published": true,
		}

		rows, err = db.GetFilteredPage(table, "created_at", page, config.Config.ArticlesPerPage, filter)
	} else {
		rows, err = db.GetPage(table, "created_at", page, config.Config.ArticlesPerPage)
	}

	var a []Article
	if err != nil {
		return a, errors.Wrap(err, "article.Page")
	}

	if err = rows.All(&a); err != nil {
		return a, errors.Wrap(err, "article.Page")
	}

	return a, nil
}

// PageCount Calculates the number of pages based on the number of articles
// and articles per page (minus 1 because we start at 0)
func PageCount(published bool) (int, error) {
	var articleCount int
	var err error

	// Get the article count
	if published {
		filter := map[string]interface{}{
			"published": true,
		}

		articleCount, err = db.GetFilteredCount(table, filter)
	} else {
		articleCount, err = db.GetCount(table)
	}

	if err != nil {
		return 0, errors.Wrap(err, "article.PageCount")
	}

	// Convert values to float64 in order to do ceil on the result
	count := float64(articleCount)
	perPage := float64(config.Config.ArticlesPerPage)

	// Divide the count to the article per page and convert to int
	return int(math.Ceil(count/perPage)) - 1, nil
}

// Get returns an article given an ID
func Get(id string, published *bool) (Article, error) {
	var a Article

	filter := map[string]interface{}{
		"id": id,
	}

	// if published is nil, we do not include it in the filter
	if published != nil {
		filter["published"] = *published
	}

	cursor, err := db.GetByFilter(table, filter)
	if err != nil {
		return a, errors.Wrap(err, "article.Get")
	}

	if err = cursor.One(&a); err != nil {
		return a, errors.Wrap(err, "article.Get")
	}

	return a, nil
}

// Update updates an article entry in the database
func (a *Article) Update() error {
	data := map[string]interface{}{
		"title":      a.Title,
		"body":       a.Body,
		"authors":    a.Authors,
		"published":  a.Published,
		"created_at": a.CreatedAt,
		"updated_at": time.Now(),
	}
	return db.Update(table, a.ID, data)
}

// Delete attempts to delete an article given an ID
func Delete(id string) error {
	return db.Delete(table, id)
}
