package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.ExecuteTemplate(w, "base", nil); err != nil {
		http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
	}
}

func FragmentHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	switch params["name"] {
	case "todoList":
		// Fetch all the tasks from our database
		tasks, err := GetTasks(db)
		if err != nil {
			log.Fatal("Error while fetching tasks: ", err)
			http.Error(w, "Error while fetching tasks: "+err.Error(), http.StatusInternalServerError)
		}

		if err := tmpl.ExecuteTemplate(w, "todoList", tasks); err != nil {
			http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
		}

	case "updateTaskForm":
		taskId, _ := strconv.Atoi(r.URL.Query().Get("taskId"))
		if taskId > 0 {
			task, err := GetTaskById(db, taskId)
			if err != nil {
				log.Fatal("Error while fetching task by id: ", err)
				http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
			}

			if err = tmpl.ExecuteTemplate(w, params["name"], task); err != nil {
				http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
			}
		}

	default:
		// load fragment without data
		if err := tmpl.ExecuteTemplate(w, params["name"], nil); err != nil {
			http.Error(w, "Fragment not found: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

func AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	// get value associated with "task" form field
	task := r.FormValue("task")

	query := "INSERT INTO tasks (task) VALUES (?)"

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(task)
	if err != nil {
		log.Fatal(err)
	}

	// return a fresh list
	todos, _ := GetTasks(db)
	tmpl.ExecuteTemplate(w, "todoList", todos)
}
