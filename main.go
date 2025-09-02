package main

import (
	"fmt"
	"os"
	"strings"
)

var tasks []string // in-memory, temporary storage

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

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("Your todo list is empty!")
		return
	}
	fmt.Println("Your tasks:")
	for i, t := range tasks {
		fmt.Printf("%d. %s\n", i+1, t)
	}
}

func addTask(args []string) {
	if len(args) == 0 {
		fmt.Println("Please provide a task description")
		return
	}
	text := strings.Join(args, " ")
	tasks = append(tasks, text)
	fmt.Println("Added task:", text)
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
		listTasks()

	case "add":
		addTask(os.Args[2:])

	case "help", "-h", "--help":
		usage()

	default:
		fmt.Println("Unknown command:", command)
		usage()
	}

}
