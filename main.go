package main

import (
	"flag"
	"fmt"
)

func main() {
	taskPtr := flag.String("add", "", "A new task to add (use quotes for multiple words, e.g., -add \"Buy groceries\")")
	flag.Parse()

	fmt.Println(*taskPtr)
}
