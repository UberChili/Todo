package main

import (
	"flag"
	"fmt"
	"github.com/UberChili/todo/config"
	"github.com/UberChili/todo/tasks"
	"time"
)

func main() {
	taskPtr := flag.String("add", "", "A new task to add (use quotes for multiple words, e.g., -add \"Buy groceries\").")
	taskList := flag.Bool("l", false, "Lists open tasks.") // Changed default to false
	delTask := flag.Int("r", 0, "Remove a task by specifing id.")
	flag.Parse()

	filename := time.Now().Format("2006-01-02")
	directory, exists := config.ConfigExists()
	if !exists {
		fmt.Println("You have no tasks directory configured. Refer to documentation.")
		return
	}

	filename = config.FormatFilename(directory, filename)

	// Handle the -add flag
	if *taskPtr != "" {
		tasks.AddNewTask(filename, *taskPtr)
		return
	}

	// Handle the -r flag
	if *delTask > 0 {
		tasklist, err := tasks.ReadOpenTasks(filename)
		if err != nil {
			fmt.Println("No tasks to remove!")
			return
		}

		if *delTask > len(tasklist) {
			fmt.Printf("Task %d does not exist.\n", *delTask)
			return
		}

		tasks.CompleteTask(filename, *delTask)
		return
	}

	// List tasks if either no flags are provided or -l flag is explicitly set
	if flag.NFlag() == 0 || *taskList {
		tasklist, err := tasks.ReadOpenTasks(filename)
		if err != nil {
			fmt.Println("You have no open tasks! Free day!")
			return
		}

		// fmt.Printf("id\tTask\n")
		fmt.Printf("\033[1mid\tTask\033[0m\n")
		config.ListOpenTasks(tasklist)
		return
	}
}
