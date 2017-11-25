package article

import (
	"math"
	"time"

	"github.com/olenedr/esamarathon/config"
	"github.com/olenedr/esamarathon/db"
	"github.com/pkg/errors"
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

	return db.Insert(table, data)
}

// All returns a slice containing all the articles
func All() ([]Article, error) {
	rows, err := db.GetAllOrderBy(table, "created_at")
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
func Page(page int) ([]Article, error) {
	rows, err := db.GetPage(table, page, config.Config.ArticlesPerPage)
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
func PageCount() (int, error) {
	// get the article count
	articleCount, err := db.GetCount(table)
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
func Get(id string) (Article, error) {
	var a Article
	cursor, err := db.GetOneById(table, id)
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
