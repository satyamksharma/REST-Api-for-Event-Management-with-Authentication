package models

import "time"

type Event struct {
	ID int 
	Name string `binding:"required"`
	Description string `binding:"required"`
	Location string `binding:"required"`
	DateTIme time.Time `binding:"required"`
	UserID int 
}

var events []Event = []Event{}


func (e Event) Save() {
	//add it to database
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}