package tasks

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

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

func ReadTasks(filename string) ([]string, error) {
	result := []string{}

	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("Error could not open file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "- [ ] ") {
			result = append(result, line)
		}
	}

	return result, nil
}
