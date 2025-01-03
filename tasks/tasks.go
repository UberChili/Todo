package tasks

import "errors"

var ErrCantAddEmptyTask = errors.New("Can't add an empty task")

type Task struct {
	Description string
}

func (t *Task) New(text string) error {
	if text == "" {
		return ErrCantAddEmptyTask
	}
	t.Description = text

	return nil
}
