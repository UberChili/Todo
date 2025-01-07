package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/UberChili/todoapp/config"
	"github.com/UberChili/todoapp/tasks"
)

func main() {
	taskPtr := flag.String("add", "", "A new task to add (use quotes for multiple words, e.g., -add \"Buy groceries\").")
	taskList := flag.Bool("l", true, "Lists open tasks.")
	delTask := flag.Int("r", 0, "Remove a task by specifing id.")
	flag.Parse()

	filename := time.Now().Format("2006-01-02")

	directory, exists := config.ConfigExists()
	if !exists {
		fmt.Println("You have no tasks directory configured. Refer to documentation.")
		return
	}

	filename = config.FormatFilename(directory, filename)

	tasklist, err := tasks.ReadOpenTasks(filename)
	if err != nil {
		fmt.Println("You have no open tasks! Free day!")
		// return
	}

	// Add a new task
	if *taskPtr != "" {
		tasks.AddNewTask(filename, *taskPtr)
		return
	}

	// Deleting a task, specified by id
	if *delTask > 0 {
		// Checking if valid task
		if *delTask > len(tasklist) {
			fmt.Printf("Task %d does not exist.", *delTask)
			return
		}
		// Remove the task
		tasks.CompleteTask(filename, *delTask)
		return
	}

	// List open tasks, this is the default behavior when no argument or flag is given
	if *taskList {
		config.ListOpenTasks(tasklist)
		return
	}
}
