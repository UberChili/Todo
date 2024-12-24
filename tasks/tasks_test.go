package tasks

import (
	"reflect"
	"testing"
)

func TestNewTask(t *testing.T) {
	tasks := Tasks{}

	t.Run("Creating and a new first task", func(t *testing.T) {
		thingToDo := "Buy milk and eggs"

		tasks.New(thingToDo)
	})

}

// func TestReadTasksFromFile(t *testing.T) {
// 	task := Tasks{}
//
// 	t.Run("Reading tasks from an existing file", func(t *testing.T) {
// 	})
// }

func TestReadTasks(t *testing.T) {
	tasks := Tasks{}

	t.Run("Reading from simple strings", func(t *testing.T) {
		one_task := "Buy eggs"
		another_task := "Buy milk"

		tasks.New(one_task)
		tasks.New(another_task)

		got := tasks.ListAsSlice()
		want := []string{one_task, another_task}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %q, want %q", got, want)
		}

	})
}
