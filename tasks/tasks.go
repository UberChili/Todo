package tasks

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

var ErrCouldNotReadFile = errors.New("Error: Could not read tasks from file")

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

func (t *Tasks) ReadFromFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return ErrCouldNotReadFile
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	// Read line by line
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if CheckForTask(line) {
			formatted := strings.TrimSuffix(strings.TrimPrefix(line, "- [ ]"), "\n")
			task := Task{len(t.tasks) + 1, formatted}
			t.tasks = append(t.tasks, task)
		}
	}

	return nil
}

func CheckForTask(line string) bool {
	prefix := "- [ ]"
	if !strings.HasPrefix(line, prefix) {
		return false
	}
	return true
}
