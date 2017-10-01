package article

import (
	"github.com/olenedr/esamarathon/db"
	"labix.org/v2/mgo"
)

var collection *mgo.Collection
var collectionName = "articles"

func query() *mgo.Collection {
	return db.Connection.C(collectionName)
}

// All returns a collection of all articles
func (a *Article) All() ([]Article, error) {
	var result []Article
	if err := query().Find(nil).All(&result); err != nil {
		return nil, err
	}

	return result, nil
}
