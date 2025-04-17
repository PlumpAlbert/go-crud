package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	gRouter := mux.NewRouter()

	gRouter.HandleFunc("/", HomeHandler)

	http.ListenAndServe(":8080", gRouter)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to home page\n")
}
