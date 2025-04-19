package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/plumpalbert/go-crud/cmd/gocrud/handlers"
	"github.com/plumpalbert/go-crud/internal/core"
)

func main() {
	config := core.NewConfig()
	defer config.Database.Close()

	gRouter := mux.NewRouter()

	gRouter.HandleFunc("/", handlers.HomeHandler)

	// Fetch fragments
	gRouter.HandleFunc("/fragment/{name}", handlers.FragmentHandler).Methods("GET")

	// Add task
	gRouter.HandleFunc("/tasks", handlers.AddTaskHandler).Methods("POST")

	// Update task
	gRouter.HandleFunc("/tasks/{id}", handlers.UpdateTaskHandler).Methods("PUT", "POST")

	// Delete task
	gRouter.HandleFunc("/tasks/{id}", handlers.DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":8080", gRouter)
}
