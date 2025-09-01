package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Println(`Usage:
	todo <command> [args]

	Commands:
	list:	List tasks
	add <text...> Add a new task

	Examples:
	todo list
	todo add "empty the dishwasher"

	`)
}

func main() {
	// os.Args holds CLI arguments. os.Args[0] is the program name.
	if len(os.Args) < 2 {
		usage()
		return
	}

	command := os.Args[1]

	switch command {
	case "list":
		// implement this in the next step
		fmt.Println("No tasks yet! (feature coming next)")

	case "add":
		// implement this in the next step
		fmt.Println("Add feature coming in next patch!")

	case "help", "-h", "--help":
		usage()

	default:
		fmt.Println("Unknown command:", command)
		usage()
	}

}
