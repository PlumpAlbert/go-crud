package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var tmpl *template.Template // For later use when introducing HTML rendering

type Task struct {
	Id   int
	Task string
	Done bool
}

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func initDB() {
	var err error

	db, err = sql.Open("mysql", "root:root@(127.0.0.1)/testdb?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}

	// Check DB connection
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	initDB()
	defer db.Close()

	gRouter := mux.NewRouter()

	gRouter.HandleFunc("/", HomeHandler)

	// Get tasks:
	gRouter.HandleFunc("/tasks", fetchTasks).Methods("GET")

	http.ListenAndServe(":8080", gRouter)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.ExecuteTemplate(w, "base", nil); err != nil {
		http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError) // use http.StatusInternalServerError to signal an error
	}
}

func fetchTasks(w http.ResponseWriter, r *http.Request) {
	// Fetch all the tasks from our database
	tasks, err := GetTasks(db)
	if err != nil {
		log.Fatal("Error while fetching tasks: ", err)
		http.Error(w, "Error while fetching tasks: "+err.Error(), http.StatusInternalServerError)
	}

	if err := tmpl.ExecuteTemplate(w, "todoList", tasks); err != nil {
		http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
	}
}

func GetTasks(dbPointer *sql.DB) ([]Task, error) {
	query := "SELECT id, task, done FROM tasks;"

	rows, err := dbPointer.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []Task

	for rows.Next() { // Loop over rows in result set
		var todo Task

		err = rows.Scan(&todo.Id, &todo.Task, &todo.Done)
		if err != nil {
			return nil, err
		}

		// Build the array of Tasks from returned rows
		tasks = append(tasks, todo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
