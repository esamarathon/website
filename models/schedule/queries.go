package schedule

import (
	"github.com/esamarathon/website/db"
	"github.com/pkg/errors"
)

func All() (scheds []ScheduleRef, err error) {
	rows, err := db.GetAllByOrder(table, "order", false)
	if err != nil {
		return scheds, errors.Wrap(err, "schedule.All")
	}

	err = rows.All(&scheds)
	if err != nil {
		err = errors.Wrap(err, "schedule.All")
	}

	return
}

// Get returns an schedule given an ID
func Get(id string) (sched ScheduleRef, err error) {
	filter := map[string]interface{}{
		"id": id,
	}

	cursor, err := db.GetByFilter(table, filter)
	if err != nil {
		return sched, errors.Wrap(err, "schedule.Get")
	}

	if err = cursor.One(&sched); err != nil {
		return sched, errors.Wrap(err, "schedule.Get")
	}

	return
}

func (s *ScheduleRef) Create() error {
	data := s.ToMap()
	_, err := db.Insert(table, data)
	return err
}

func (s *ScheduleRef) Update() error {
	return db.Update(table, s.ID, s.ToMap())
}

func Delete(id string) error {
	return db.Delete(table, id)
}

// IsStored checks whether we're using the menu from the DB or Default
func IsStored() bool {
	m, err := All()
	if err != nil || len(m) == 0 {
		return false
	}
	return true
}

func (s *ScheduleRef) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"url":   s.Url,
		"title": s.Title,
		"order": s.Order,
	}
}
