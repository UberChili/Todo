package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/UberChili/todoapp/config"
	"github.com/UberChili/todoapp/tasks"
)

func main() {
	taskPtr := flag.String("add", "", "A new task to add (use quotes for multiple words, e.g., -add \"Buy groceries\")")
	taskList := flag.Bool("l", true, "Lists open tasks")
	flag.Parse()

	filename := time.Now().Format("2006-01-02")
	fmt.Println(filename)

	directory, exists := config.ConfigExists()
	if !exists {
		fmt.Println("You have no tasks directory configured. Refer to documentation.")
		return
	}

	fmt.Println(directory)

	filename = directory + "/" + filename + ".md"

	// TODO
	// List all open tasks in the dailies file.
	// Now that I think about it... This is probably the easiest thing and
	// should be implemented first
	tasks, err := tasks.ReadOpenTasks(filename)
	if err != nil {
		fmt.Println("You have no open tasks! Free day!" + err.Error()) // Remember to delete the error, just testing
		return
	}

	for _, task := range tasks {
		fmt.Println(task)
	}

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

	// fmt.Println(*taskPtr)
	// fmt.Println(*taskList)
}
