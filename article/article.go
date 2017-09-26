package article

import (
	"time"

	"labix.org/v2/mgo/bson"
)

// Article describes the format of an article
type Article struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
