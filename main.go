package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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

// save tasks to JSON file
func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(taskFile, data, 0644)
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

func addTask(args []string) {
	if len(args) == 0 {
		fmt.Println("Please provide a task description")
		return
	}
	text := strings.Join(args, " ")

	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	tasks = append(tasks, Task{Text: text, Done: false})

	err = saveTasks(tasks)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	fmt.Println("Added tasks:", text)

}

func markDone(args []string) {

	if len(args) != 1 {
		fmt.Println("Usage: todo done <task number>")
		return
	}

	// convert argument string to int
	index, err := strconv.Atoi(args[0])
	if err != nil || index < 1 {
		fmt.Println("Please provide a valid task number")
		return
	}

	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	if index > len(tasks) {
		fmt.Println("Task number out of range")
		return
	}

	// mark as done
	tasks[index-1].Done = true
	err = saveTasks(tasks)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	fmt.Printf("Task %d marked as done!\n", index)
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

	case "done":
		markDone((os.Args[2:]))

	case "help", "-h", "--help":
		usage()

	default:
		fmt.Println("Unknown command:", command)
		usage()
	}

}
