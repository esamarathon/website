package esamarathon

import (
	"time"

	"github.com/olenedr/esamarathon/db"
	"labix.org/v2/mgo"

	"labix.org/v2/mgo/bson"
)

type Article struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	title      string
	body       string
	created_at time.Time
	updated_at time.Time
}

var collection *mgo.Collection
var collectionName = "articles"

func init() {
	collection = db.Connection.C(collectionName)
}

func (a *Article) All() ([]Article, error) {
	var result []Article
	if err := collection.Find(nil).All(&result); err != nil {
		return nil, err
	}

	return result, nil
}
