package todo

import (
	"fmt"
	"time"
)

type Task struct {
	Title     string
	Text      string
	Completed bool

	CreatedAt time.Time
	// TODO: Поле можно переименовать в CompletedAt для более привычного
	// английского нейминга и единообразия с Completed.
	CompleteAt *time.Time
}

func NewTask(title, text string) Task {
	return Task{
		Title: title,
		Text:  text,

		Completed:  false,
		CreatedAt:  time.Now(),
		CompleteAt: nil,
	}
}

func (t *Task) Complete() {
	// TODO: Доменный слой лучше модифицировать так, чтобы он возвращал error,
	// а не печатал в stdout: сейчас HTTP-слою трудно различать причины отказа.
	if t.Completed == true {
		fmt.Println(ErrTaskAlreadyExists)
		return
	}
	t.Completed = true

	now := time.Now()
	t.CompleteAt = &now
}

func (t *Task) Uncomplete() {
	if t.Completed == false {
		fmt.Println(ErrTaskAlreadyExists)
		return
	}
	t.Completed = false

	t.CompleteAt = nil
}
