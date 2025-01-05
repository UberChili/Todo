package tasks

import (
	"fmt"
	"reflect"
	"testing"
)

const (
	inputFile        = "test.md"
	inputFileNoTasks = "testempty.md"
)

func TestAdd(t *testing.T) {
	t.Run("Giving only task description as input", func(t *testing.T) {
		task := Task{}
		input := "Buy Coca-Cola"

		err := task.New(input)
		assertError(t, err, nil)

		got := task.Description

		if got != input {
			t.Errorf("got %q, want %q", got, input)
		}
	})

	t.Run("Giving an empty input", func(t *testing.T) {
		task := Task{}
		input := ""

		err := task.New(input)
		assertError(t, err, ErrCantAddEmptyTask)

		got := task.Description

		if got != input {
			t.Errorf("got %q, want %q", got, input)
		}
	})
}

func TestReadLines(t *testing.T) {
	t.Run("Reading a file with correct tasks", func(t *testing.T) {
		want := []string{"- [ ] Buy eggs", "- [ ] Buy milk", "- [ ] Start working on todo CLI app"}

		got, err := ReadOpenTasks(inputFile)
		assertError(t, err, nil)

		for _, line := range got {
			fmt.Println(line)
		}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("Got %q, want %q", got, want)
		}
	})

	t.Run("Reading completed tasks", func(t *testing.T) {
		want := []string{"- [x] This is a completed task", "- [x] Made some progress with my project today"}

		got, err := ReadClosedTasks(inputFile)
		assertError(t, err, nil)

		for _, line := range got {
			fmt.Println(line)
		}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("Got %q, want %q", got, want)
		}
	})

	t.Run("Reading a file with no correct tasks", func(t *testing.T) {
		want := []string{}

		got, err := ReadTasks(inputFileNoTasks)
		assertError(t, err, nil)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("Got %q, want %q", got, want)
		}
	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q, want %q", got, want)
	}
}
