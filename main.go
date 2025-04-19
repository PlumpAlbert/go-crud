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

	// Fetch fragments
	gRouter.HandleFunc("/fragment/{name}", FragmentHandler).Methods("GET")

	// Add task
	gRouter.HandleFunc("/tasks", AddTaskHandler).Methods("POST")

	// Update task
	gRouter.HandleFunc("/tasks/{id}", UpdateTaskHandler).Methods("PUT", "POST")

	// Delete task
	gRouter.HandleFunc("/tasks/{id}", DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":8080", gRouter)
}
