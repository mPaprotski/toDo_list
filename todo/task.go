package todo

import (
	"fmt"
	"time"
)

type Task struct {
	Title  string
	Text   string
	IsDone bool

	CreatedAt time.Time
	DoneAt    *time.Time
}

func NewTask(title, text string) Task {
	return Task{
		Title: title,
		Text:  text,

		IsDone:    false,
		CreatedAt: time.Now(),
		DoneAt:    nil,
	}
}

func (t *Task) Done() {
	if t.IsDone == true {
		fmt.Println(taskAlreadyDone)
		return
	}
	t.IsDone = true

	now := time.Now()
	t.DoneAt = &now
}
