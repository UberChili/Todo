package tasks

import "testing"

func TestNewTask(t *testing.T) {
	tasks := Tasks{}

	t.Run("Creating and a new first task", func(t *testing.T) {
		thingToDo := "Buy milk and eggs"

		tasks.New(thingToDo)
	})
}

// func TestManagingTasks(t *testing.T) {
// 	tasks := Tasks{}
//
// 	t.Run("Creating a fresh task list", func(t *testing.T) {
// 	})
// }
