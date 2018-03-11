package menu

import (
	"fmt"

	"github.com/esamarathon/website/db"
	"github.com/pkg/errors"
)

const table = "menu_items"

type Menu struct {
	Items []MenuItem
}

type MenuItem struct {
	ID     string `json:"_id,omitempty" gorethink:"id,omitempty"`
	Title  string `json:"title,omitempty" gorethink:"title,omitempty"`
	Link   string `json:"link,omitempty" gorethink:"link,omitempty"`
	NewTab bool   `json:"new_tab,omitempty" gorethink:"new_tab,omitempty"`
}

func Default() Menu {
	return Menu{
		[]MenuItem{
			MenuItem{
				Title:  "Home",
				Link:   "/",
				NewTab: false,
			},
			MenuItem{
				Title:  "News",
				Link:   "/news",
				NewTab: false,
			},
			MenuItem{
				Title:  "Schedule",
				Link:   "/schedule",
				NewTab: false,
			},
			MenuItem{
				Title:  "Donate",
				Link:   "https://www.speedrun.com/esa2017/donate",
				NewTab: true,
			},
			MenuItem{
				Title:  "Tickets",
				Link:   "https://esamarathon.eventbrite.com/",
				NewTab: true,
			},
			MenuItem{
				Title:  "Forum",
				Link:   "https://www.speedrun.com/ESA_Winter_2018/forum",
				NewTab: true,
			},
		},
	}
}

// Get returns an instance of the Menu struct
func Get() Menu {
	m, err := All()
	if err != nil {
		fmt.Printf("%v", err)
		return Default()
	}
	fmt.Printf("%v", m)
	return Default()
}

// All returns a slice containing all the menu items
func All() ([]MenuItem, error) {
	rows, err := db.GetAll(table)
	var m []MenuItem
	if err != nil {
		return m, errors.Wrap(err, "menu.All")
	}

	if err = rows.All(&m); err != nil {
		return m, errors.Wrap(err, "menu.All")
	}

	return m, nil
}

// Create inserts a new menu item into the database
func (m *MenuItem) Create() error {
	data := map[string]interface{}{
		"title":   m.Title,
		"link":    m.Link,
		"new_tab": m.NewTab,
	}

	_, err := db.Insert(table, data)
	return err
}
