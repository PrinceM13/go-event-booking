package models

import "time"

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events []Event = []Event{}

func (e Event) Save() {
	// later: save to a database
	events = append(events, e)
}

func GetAllEvents() []Event {
	// later: fetch from a database
	return events
}
