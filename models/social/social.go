package social

import (
	"github.com/esamarathon/website/db"
	"github.com/pkg/errors"
)

const table = "social_items"

const (
	TwitterOrder = iota
	DiscordOrder
	YoutubeOrder
	TwitchOrder
	ForumOrder
	FacebookOrder
)

type SocialLinks struct {
	Items []SocialItem
}

type SocialItem struct {
	ID       string `json:"_id,omitempty" gorethink:"id,omitempty"`
	Title    string `json:"title,omitempty" gorethink:"title,omitempty"`
	Link     string `json:"link,omitempty" gorethink:"link,omitempty"`
	Image    string `json:"image,omitempty" gorethink:"image,omitempty"`
	ImageAlt string `json:"imagealt,omitempty" gorethink:"imagealt,omitempty"`
	NewTab   bool   `json:"new_tab,omitempty" gorethink:"new_tab,omitempty"`
	Order    int    `json:"order,omitempty" gorethink:"order,omitempty"`
}

func Default() SocialLinks {
	return SocialLinks{
		[]SocialItem{
			SocialItem{
				Title:    "Twitter",
				Link:     "https://www.twitter.com/esamarathon",
				Image:    "/static/img/ico-twitter.svg",
				ImageAlt: "Twitter icon",
				NewTab:   true,
				Order:    TwitterOrder,
			},
			SocialItem{
				Title:    "Discord",
				Link:     "https://www.discord.gg/0TZ2NlveujtasAqb",
				Image:    "/static/img/ico-discord.svg",
				ImageAlt: "Discord icon",
				NewTab:   true,
				Order:    DiscordOrder,
			},
			SocialItem{
				Title:    "Youtube",
				Link:     "https://www.youtube.com/user/EuroSpeedAssembly",
				Image:    "/static/img/ico-youtube.svg",
				ImageAlt: "Youtube icon",
				NewTab:   true,
				Order:    YoutubeOrder,
			},
			SocialItem{
				Title:    "Twitch",
				Link:     "https://www.twitch.tv/esamarathon",
				Image:    "/static/img/ico-twitch.svg",
				ImageAlt: "Twitch icon",
				NewTab:   true,
				Order:    TwitchOrder,
			},
			SocialItem{
				Title:    "Forum",
				Link:     "https://www.speedrun.com/esa2018/forum",
				Image:    "/static/img/ico-forum.svg",
				ImageAlt: "Forum icon",
				NewTab:   true,
				Order:    ForumOrder,
			},
			SocialItem{
				Title:    "Facebook",
				Link:     "https://www.facebook.com/europeanspeedrunnerassembly",
				Image:    "/static/img/ico-facebook.svg",
				ImageAlt: "Facebook icon",
				NewTab:   true,
				Order:    FacebookOrder,
			},
		},
	}
}

// Find a menu item given an id
func Find(id string) (SocialItem, error) {
	var si SocialItem

	filter := map[string]interface{}{
		"id": id,
	}

	cursor, err := db.GetByFilter(table, filter)
	if err != nil {
		return si, errors.Wrap(err, "social.Find")
	}

	if err = cursor.One(&si); err != nil {
		return si, errors.Wrap(err, "social.Find")
	}

	return si, nil
}

// IsStored checks whether we're using the social links from the DB or Default
func IsStored() bool {
	m, err := All()
	if err != nil || len(m) == 0 {
		return false
	}
	return true
}

func Get() SocialLinks {
	sls, err := All()
	if err != nil || len(sls) == 0 {
		return Default()
	}
	return SocialLinks{sls}
}

func All() ([]SocialItem, error) {
	rows, err := db.GetAllByOrder(table, "order", false)
	var sis []SocialItem

	if err != nil {
		return sis, errors.Wrap(err, "social.All")
	}

	if err = rows.All(&sis); err != nil {
		return sis, errors.Wrap(err, "social.All")
	}

	return sis, nil
}

func (s *SocialLinks) Insert() []error {
	err := make([]error, 0)
	for _, si := range s.Items {
		e := si.Create()
		if e != nil {
			err = append(err, e)
		}
	}
	return err
}

// Create inserts a new social item into the database
func (si *SocialItem) Create() error {
	data := si.ToMap()
	_, err := db.Insert(table, data)
	return err
}

// Update updates a social item entry
func (si *SocialItem) Update() error {
	data := si.ToMap()
	return db.Update(table, si.ID, data)
}

func (si *SocialItem) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"title":    si.Title,
		"link":     si.Link,
		"image":    si.Image,
		"imageAlt": si.ImageAlt,
		"new_tab":  si.NewTab,
		"order":    si.Order,
	}
}
