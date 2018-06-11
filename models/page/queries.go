package page

import (
	"math"
	"time"

	"github.com/esamarathon/website/config"
	"github.com/esamarathon/website/db"
	"github.com/pkg/errors"

	r "gopkg.in/gorethink/gorethink.v3"
)

// Get returns an article given an ID
func Get(id string, published *bool) (p Page, err error) {
	filter := map[string]interface{}{
		"id": id,
	}

	// if published is nil, we do not include it in the filter
	if published != nil {
		filter["published"] = *published
	}

	cursor, err := db.GetByFilter(table, filter)
	if err != nil {
		return p, errors.Wrap(err, "page.Get")
	}

	if err = cursor.One(&p); err != nil {
		return p, errors.Wrap(err, "page.Get")
	}

	return p, nil
}

func GetFromName(name string, published *bool) (p Page, err error) {
	filter := map[string]interface{}{
		"friendly_name": name,
	}

	// if published is nil, we do not include it in the filter
	if published != nil {
		filter["published"] = *published
	}

	cursor, err := db.GetByFilter(table, filter)
	if err != nil {
		err = errors.Wrap(err, "page.GetFromName")
		return
	}

	if err = cursor.One(&p); err != nil {
		err = errors.Wrap(err, "page.Get")
		return
	}

	return p, nil
}

// All returns a slice containing all the articles
func All() (ps []Page, err error) {
	rows, err := db.GetAllByOrder(table, "created_at", true)

	if err != nil {
		return ps, errors.Wrap(err, "page.All")
	}

	if err = rows.All(&ps); err != nil {
		return ps, errors.Wrap(err, "page.All")
	}

	return ps, nil
}

// Page returns the pages of a given page
func Pagination(page int, published bool) (ps []Page, err error) {
	var rows *r.Cursor

	if published {
		filter := map[string]interface{}{
			"published": true,
		}

		rows, err = db.GetFilteredPage(table, "created_at", page, config.Config.ArticlesPerPage, filter)
	} else {
		rows, err = db.GetPage(table, "created_at", page, config.Config.ArticlesPerPage)
	}

	if err != nil {
		return ps, errors.Wrap(err, "article.Page")
	}

	if err = rows.All(&ps); err != nil {
		return ps, errors.Wrap(err, "article.Page")
	}

	return
}

// PageCount Calculates the number of pages based on the number of articles
// and articles per page (minus 1 because we start at 0)
func PageCount(published bool) (pageCount int, err error) {
	// Get the article count
	if published {
		filter := map[string]interface{}{
			"published": true,
		}

		pageCount, err = db.GetFilteredCount(table, filter)
	} else {
		pageCount, err = db.GetCount(table)
	}

	if err != nil {
		return 0, errors.Wrap(err, "article.PageCount")
	}

	// Convert values to float64 in order to do ceil on the result
	count := float64(pageCount)
	perPage := float64(config.Config.ArticlesPerPage)

	// Divide the count to the article per page and convert to int
	return int(math.Ceil(count/perPage)) - 1, nil
}

func (p *Page) Create() error {
	data := map[string]interface{}{
		"title":         p.Title,
		"body":          p.Body,
		"published":     p.Published,
		"created_at":    time.Now(),
		"updated_at":    time.Now(),
		"friendly_name": p.FriendlyName,
	}

	_, err := db.Insert(table, data)
	return err
}

// Update updates an article entry in the database
func (a *Page) Update() error {
	data := map[string]interface{}{
		"title":         a.Title,
		"body":          a.Body,
		"published":     a.Published,
		"created_at":    a.CreatedAt,
		"updated_at":    time.Now(),
		"friendly_name": a.FriendlyName,
	}
	return db.Update(table, a.ID, data)
}

// Delete attempts to delete an article given an ID
func Delete(id string) error {
	return db.Delete(table, id)
}
