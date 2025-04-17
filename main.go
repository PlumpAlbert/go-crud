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

	http.ListenAndServe(":8080", gRouter)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.ExecuteTemplate(w, "base", nil); err != nil {
		http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError) // use http.StatusInternalServerError to signal an error
	}
}
