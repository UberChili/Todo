package tasks

import (
	"errors"
	"fmt"
	"time"
)

type Tasks struct {
	tasks []Task
}

type Task struct {
	id          int
	description string
	time        string
}

func (t *Tasks) New(description string) error {
	if description == "" {
		return errors.New("Error: Can't add an empy task.")
	}

	currentTime := time.Now().Format("01-02-2006 15:04")
	task := Task{id: len(t.tasks) + 1, description: description, time: currentTime}
	t.tasks = append(t.tasks, task)

	fmt.Println(task)

	return nil
}
