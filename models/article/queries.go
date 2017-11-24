package article

import (
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
		"created_at": time.Now(),
		"updated_at": time.Now(),
	}

	return db.Insert(table, data)
}

// All returns a slice containing all the articles
func All() ([]Article, error) {
	rows, err := db.GetAll(table)
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
	count, err := db.GetCount(table)
	if err != nil {
		return 0, errors.Wrap(err, "article.PageCount")
	}
	return (count / config.Config.ArticlesPerPage) - 1, nil
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
		"created_at": a.CreatedAt,
		"updated_at": time.Now(),
	}
	return db.Update(table, a.ID, data)
}
