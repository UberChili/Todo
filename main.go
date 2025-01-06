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
		fmt.Println("You have no open tasks! Free day!" + err.Error()) // Remember to delete the error, just testing
		return
	}

	if *delTask > 0 {
		if *delTask < 1 || *delTask > len(tasklist) {
			fmt.Printf("Task %d does not exist.", *delTask)
			return
		}
		// Remove the task
		tasks.CompleteTask(filename, *delTask)
		return
	}

	if *taskList {
		config.ListOpenTasks(tasklist)
		return
	}

	// TODO
	// Mark a task as completed.
	// 1. Remove from tasks slice
	// 2. Mark completed (by editing the .md file)

	if *taskPtr != "" {
		// We do not necessarily list open tasks
		*taskList = false

		// TODO
		// 1. Get today's date
		// 2. Check if daily file with today's date as title already exists
		// 3. If it doesn't exist, create file in correct directory
		// 4. Open file
		// 5. Get to end of File and write new task
		//
		// We should just use a function. Something like:
		// AddNewTask(*taskPtr, date)

	}

	if *taskList {
		config.ListOpenTasks(tasklist)
		return
	}

	// fmt.Println(*taskPtr)
	// fmt.Println(*taskList)
}
