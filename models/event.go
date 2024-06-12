package models

import "time"

type Event struct {
	ID int 
	Name string
	Description string
	Location string
	DateTIme time.Time
	userID int
}

var events []Event = []Event{}


func (e Event) Save() {
	//add it to database
	events = append(events, e)
}