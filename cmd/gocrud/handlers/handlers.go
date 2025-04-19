package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/plumpalbert/go-crud/cmd/gocrud/utils"
	"github.com/plumpalbert/go-crud/internal/core"
	"github.com/plumpalbert/go-crud/pkg/gocrud"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if err := core.Config.Templates.ExecuteTemplate(w, "base", nil); err != nil {
		http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
	}
}

func FragmentHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	switch params["name"] {
	case "todoList":
		// Fetch all the tasks from our database
		tasks, err := utils.GetTasks(core.Config.Database)
		if err != nil {
			log.Fatal("Error while fetching tasks: ", err)
			http.Error(w, "Error while fetching tasks: "+err.Error(), http.StatusInternalServerError)
		}

		if err := core.Config.Templates.ExecuteTemplate(w, "todoList", tasks); err != nil {
			http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
		}

	case "updateTaskForm":
		taskId, _ := strconv.Atoi(r.URL.Query().Get("taskId"))
		if taskId > 0 {
			task, err := utils.GetTaskById(core.Config.Database, taskId)
			if err != nil {
				log.Fatal("Error while fetching task by id: ", err)
				http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
			}

			if err = core.Config.Templates.ExecuteTemplate(w, params["name"], task); err != nil {
				http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
			}
		}

	default:
		// load fragment without data
		if err := core.Config.Templates.ExecuteTemplate(w, params["name"], nil); err != nil {
			http.Error(w, "Fragment not found: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

func AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	// get value associated with "task" form field
	task := r.FormValue("task")

	query := "INSERT INTO tasks (task) VALUES (?)"

	stmt, err := core.Config.Database.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(task)
	if err != nil {
		log.Fatal(err)
	}

	// return a fresh list
	todos, _ := utils.GetTasks(core.Config.Database)
	core.Config.Templates.ExecuteTemplate(w, "todoList", todos)
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskId, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	title := r.FormValue("task")
	var done bool
	switch strings.ToLower(r.FormValue("done")) {
	case "yes", "on":
		done = true
	default:
		done = false
	}

	task := gocrud.Task{
		Id: taskId, Task: title, Done: done,
	}

	query := "UPDATE tasks SET task = ?, done = ? WHERE id = ?"

	result, err := core.Config.Database.Exec(query, task.Task, task.Done, task.Id)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		fmt.Println("No rows updated")
	}

	tasks, _ := utils.GetTasks(core.Config.Database)

	core.Config.Templates.ExecuteTemplate(w, "todoList", tasks)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskId, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	query := "DELETE FROM tasks WHERE id = ?"

	result, err := core.Config.Database.Exec(query, taskId)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		fmt.Println("No rows deleted")
	}

	tasks, _ := utils.GetTasks(core.Config.Database)

	core.Config.Templates.ExecuteTemplate(w, "todoList", tasks)
}
