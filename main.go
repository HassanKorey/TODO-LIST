package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID    int
	Title string
	Done  bool
}

func main() {

	args := os.Args

	if len(args) < 2 {
		fmt.Println("please provide a command")
		return
	}

	switch args[1] {
	case "add":
		if len(args) <= 2 {
			fmt.Println("please provide a task title")
			return
		}
		title := args[2]
		fmt.Println("Adding...", title)
	case "list":
		fmt.Println("listing...")
	case "done":
		fmt.Println("marking as done...")
	case "delete":
		fmt.Println("deleting...")
	default:
		fmt.Println("unknown command")
	}

}

func loadTask() []Task {
	var tasks []Task
	taskFile, err := os.ReadFile("tasks.json")

	if err != nil {
		return tasks
	}
	json.Unmarshal(taskFile, &tasks)

	return tasks
}

func saveTask(tasks []Task) {

}
