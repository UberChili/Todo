package tasks

import (
	"reflect"
	"testing"
)

const (
	filePath     = "test.md"
	one_task     = "Buy eggs"
	another_task = "Buy milk"
)

func TestNewTask(t *testing.T) {
	tasks := Tasks{}

	t.Run("Creating and a new first task", func(t *testing.T) {

		got := tasks.New(one_task)
		if got != nil {
			t.Errorf("%q", got)
		}
	})

}

func TestReadTasks(t *testing.T) {

	t.Run("Reading from simple strings", func(t *testing.T) {
		tasks := Tasks{}

		err := tasks.New(one_task)
		if err != nil {
			t.Errorf("%q", err)
		}
		err = tasks.New(another_task)
		if err != nil {
			t.Errorf("%q", err)
		}

		got, err := tasks.ListAsSlice()
		if err != nil {
			t.Errorf("%q", err)
		}

		want := []string{one_task, another_task}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %q, want %q", got, want)
		}

	})

	// Using file test.md
	t.Run("Reading from a file", func(t *testing.T) {
		tasks := Tasks{}

		want := []string{one_task, another_task}

		err := tasks.ReadFromFile(filePath)
		if err != nil {
			t.Errorf("%q", err)
		}

		got := tasks.tasks
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %q, got %q", got, want)
		}
	})
}

func TestDetectOpenTasks(t *testing.T) {
	t.Run("Checking correct task lines", func(t *testing.T) {
		example := "- [ ] Buy eggs"
		another_example := "- [ ] buy milk"

		lines := []string{example, another_example}

		for _, i := range lines {
			if !CheckForTask(i) {
				t.Errorf("Expected all to be valid task lines, given %q", lines)
			}
		}

	})

	t.Run("Checking incorrect task lines", func(t *testing.T) {
		feeling := "- Today I felt kind of depressed"
		just_a_thought := "There was a lot of traffic in the street today"

		lines := []string{feeling, just_a_thought}

		for _, i := range lines {
			if CheckForTask(i) {
				t.Errorf("Expected none to be valid task lines, given %q", lines)
			}
		}
	})
}
