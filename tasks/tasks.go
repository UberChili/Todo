package tasks

import "errors"

type Tasks struct {
	tasks []Task
}

type Task struct {
	id          int
	description string
}

func (t *Tasks) New(description string) error {
	if description == "" {
		return errors.New("Error: Can't add an empy task.")
	}

	task := Task{id: len(t.tasks) + 1, description: description}
	t.tasks = append(t.tasks, task)

	return nil
}

func (t Tasks) ListAsSlice() ([]string, error) {
	var tasks []string
	for _, task := range t.tasks {
		tasks = append(tasks, task.description)
	}
	return tasks, nil
}

func (t *Tasks) GetTasksFromFile(file string) error {
	return nil
}
