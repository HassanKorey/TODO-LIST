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
// Here's what the add case needs to do:

// Load existing tasks from file
// Create a new Task with the title from args[2]
// Give it an ID (just len(tasks) + 1 is fine)
// Append it to the slice
// Save the updated slice back to file
// Print confirmation
// Your task — replace the add case with this logic. Try writing it yourself first without me 
// giving you the code.You have all the tools: loadTask(), saveTask(), append(), and the Task struct.
	switch args[1] {
	case "add":
		if len(args) <= 2 {
			fmt.Println("incorrect argument length")
			return
		}
		title := args[2]
		loadedTask := loadTask()
			res := Task{ID: len(loadedTask)+1,Title: title, Done: true}
			loadedTask =  append(loadedTask, res)
			saveTask(loadedTask)
			fmt.Println(loadedTask)
		
		fmt.Println("Done adding task :)")
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

	savedTask,_ := json.Marshal(tasks)
	
	os.WriteFile("tasks.json", savedTask, 0644)
}
