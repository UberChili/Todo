package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	taskPtr := flag.String("add", "", "A new task to add (use quotes for multiple words, e.g., -add \"Buy groceries\")")
	taskList := flag.Bool("l", true, "Lists open tasks")
	flag.Parse()

	// TODO Check if "configuration" has been set
	if !configExists() {
		fmt.Println("You have no tasks directory configured. Refer to documentation.")
		return
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
	}

	// fmt.Println(*taskPtr)
	// fmt.Println(*taskList)
}

func configExists() bool {
	directory := os.Getenv("DAILIES_DIRECTORY")

	if directory == "" {
		return false
	}

	fmt.Println(directory)

	if _, err := os.Stat(directory); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Directory does not exist")
			return false
		}
		return true
	}

	return true
}
