package config

import (
	"fmt"
	"os"
	"strings"
)

func ConfigExists() (string, bool) {
	directory := os.Getenv("DAILIES_DIRECTORY")

	if directory == "" {
		return "", false
	}

	if _, err := os.Stat(directory); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Directory does not exist")
			return "", false
		}
	}

	return directory, true
}

func ListOpenTasks(tasks []string) {
	for i, task := range tasks {
		task = strings.TrimPrefix(task, "- [ ] ")
		fmt.Printf("%d\t%s\n", i+1, task)
	}
}

func FormatFilename(directory, text string) string {
	return directory + "/" + text + ".md"
}
