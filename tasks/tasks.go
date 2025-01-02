package tasks

import "errors"

type Task struct {
	description string
}

func (t *Task) New(task string) error {
	if task == "" {
		return errors.New("Error: Can't add an empty task")
	}

	t.description = task
	return nil
}
