package menu

import (
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
	Order  int    `json:"order,omitempty" gorethink:"order,omitempty"`
}

func Default() Menu {
	return Menu{
		[]MenuItem{
			MenuItem{
				Title:  "Home",
				Link:   "/",
				NewTab: false,
				Order:  0,
			},
			MenuItem{
				Title:  "News",
				Link:   "/news",
				NewTab: false,
				Order:  1,
			},
			MenuItem{
				Title:  "Schedule",
				Link:   "/schedule",
				NewTab: false,
				Order:  2,
			},
			MenuItem{
				Title:  "Donate",
				Link:   "https://www.speedrun.com/esa2017/donate",
				NewTab: true,
				Order:  3,
			},
			MenuItem{
				Title:  "Tickets",
				Link:   "https://esamarathon.eventbrite.com/",
				NewTab: true,
				Order:  4,
			},
			MenuItem{
				Title:  "Forum",
				Link:   "https://www.speedrun.com/ESA_Winter_2018/forum",
				NewTab: true,
				Order:  5,
			},
		},
	}
}

// Get returns an instance of the Menu struct
func Get() Menu {
	m, err := All()
	if err != nil || len(m) == 0 {
		return Default()
	}
	return Menu{m}
}

// IsStored checks whether we're using the menu from the DB or Default
func IsStored() bool {
	m, err := All()
	if err != nil || len(m) == 0 {
		return false
	}
	return true
}

// All returns a slice containing all the menu items
func All() ([]MenuItem, error) {
	rows, err := db.GetAllByOrder(table, "order", false)
	var m []MenuItem
	if err != nil {
		return m, errors.Wrap(err, "menu.All")
	}

	if err = rows.All(&m); err != nil {
		return m, errors.Wrap(err, "menu.All")
	}

	return m, nil
}

func (m *Menu) Insert() []error {
	err := make([]error, 0)
	for _, m := range m.Items {
		e := m.Create()
		if e != nil {
			err = append(err, e)
		}
	}
	return err
}

// Create inserts a new menu item into the database
func (m *MenuItem) Create() error {
	data := map[string]interface{}{
		"title":   m.Title,
		"link":    m.Link,
		"new_tab": m.NewTab,
		"order":   m.Order,
	}

	_, err := db.Insert(table, data)
	return err
}
