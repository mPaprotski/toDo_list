package scanner

import (
	"time"
)

type Event struct {
	description string
	userInput   string
	dateAt      time.Time
}

func NewEvent(description, userInput string) Event {
	return Event{
		description: description,
		userInput:   userInput,
		dateAt:      time.Now(),
	}
}
