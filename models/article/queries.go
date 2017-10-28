package article

import (
	"time"

	"github.com/olenedr/esamarathon/db"
	"github.com/pkg/errors"
)

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

func (a *Article) Update() error {
	data := map[string]interface{}{
		"title":      a.Title,
		"body":       a.Body,
		"authors":    a.Authors,
		"created_at": time.Now(),
		"updated_at": time.Now(),
	}
	return db.Update(table, data)
}
