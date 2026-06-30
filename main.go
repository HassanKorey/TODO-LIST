package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID    int
	Title string
	Done  bool
}

func main() {

	args := os.Args

	if len(args) < 2 {
		fmt.Println("Starting server on :8080...")
		http.HandleFunc("/", listHandler)
		http.ListenAndServe(":8080", nil)
		return
	}
	action := strings.ToLower(args[1])
	switch action {
	case "add":
		if len(args) <= 2 {
			fmt.Println("incorrect argument length")
			return
		}
		title := args[2]
		loadedTask := loadTask()
		res := Task{ID: len(loadedTask) + 1, Title: title}
		loadedTask = append(loadedTask, res)
		saveTask(loadedTask)
		fmt.Println("Done adding task :)")

	case "list":
		loadedTask := loadTask()

		if len(loadedTask) == 0 {
			fmt.Println("no tasks found")
			return
		}

		for _, task := range loadedTask {
			if task.Done == false {
				fmt.Printf("%d. [ ] %s\n", task.ID, task.Title)
			} else {
				fmt.Printf("%d. [x] %s\n", task.ID, task.Title)
			}
		}
	case "done":
		loadedTask := loadTask()
		if len(args) <= 2 {
			fmt.Println("No task targeted")
			return
		}
		taskNum, _ := strconv.Atoi(args[2])
		for i, task := range loadedTask {
			if task.ID == taskNum {
				loadedTask[i].Done = true
			}
		}
		saveTask(loadedTask)
		fmt.Println("marking as done...")
	case "delete":

		if len(args) <= 2 {
			fmt.Println("invalid argument, input the correct one")
			return
		}
		loadedTask := loadTask()
		taskNum, _ := strconv.Atoi(args[2])

		for i, task := range loadedTask {

			if task.ID == taskNum {
				loadedTask = append(loadedTask[:i], loadedTask[i+1:]...)
				break
			}
		}
		saveTask(loadedTask)
		fmt.Println("deleting...")
	default:
		fmt.Println("unknown command")
	}

}
func listHandler(w http.ResponseWriter, r *http.Request) {

	tasks := loadTask()
	tmpl := `
		
			<h1>My Todo List</h1>
	<ul>
	{{range .}}
		<li>{{.Title}} - {{if .Done}}Done{{else}}Pending{{end}}</li>
	{{end}}
	</ul>
		`
	parsed := template.Must(template.New("list").Parse(tmpl))
	parsed.Execute(w, tasks)
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

	savedTask, _ := json.Marshal(tasks)

	os.WriteFile("tasks.json", savedTask, 0644)
}
