package main

import "time"

type Event struct {
	ID          int
	Name        string
	Description string
	Location    string
	Date        time.Time
}
var events []Event
func (e Event) save(){
 events = append(events, e)
}
