package tasks

import (
	"reflect"
	"testing"
)

const filePath = "test.md"

func TestNewTask(t *testing.T) {
	tasks := Tasks{}

	t.Run("Creating and a new first task", func(t *testing.T) {
		thingToDo := "Buy milk and eggs"

		got := tasks.New(thingToDo)
		if got != nil {
			t.Errorf("%q", got)
		}
	})

}

func TestReadTasks(t *testing.T) {
	tasks := Tasks{}

	t.Run("Reading from simple strings", func(t *testing.T) {
		one_task := "Buy eggs"
		another_task := "Buy milk"

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
}
