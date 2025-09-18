package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// old // var tasks []string // in-memory, temporary storage

type Task struct {
	Text string `json: "text"`
	Done bool   `json: "done"`
}

const taskFile = "tasks.json"

// load tasks from JSON file
func loadTasks() ([]Task, error) {
	data, err := os.ReadFile(taskFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil //no file, return empty list
		}
		return nil, err
	}
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

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

// load tasks from JSON file
func loadTasks() ([]Task, error) {
	data, err := os.ReadFile(taskFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil // no file yet, return empty list
		}
		return nil, err
	}
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// save tasks to JSON file
func saveTasks(tasks []Task) err {
	data, err := json.MarshalIndent(tasks, "", " ") {
		if err != nil {
			return err
		}
		return os.WriteFile(taskFile, data, 0644)
	}
}

func listTasks() {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}
	if len(tasks) == 0 {
		fmt.Println("Your todo list is empty!")
		return
	}
	fmt.Println("Your tasks:")
	for i, t := range tasks {
		status := "[ ]"
		if t.Done {
			status = "[x]"
		}
		fmt.Printf("%d. %s %s\n", i+1, status, t.Text)
	}
}

//implement actual add task
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
