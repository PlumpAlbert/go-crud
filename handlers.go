package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.ExecuteTemplate(w, "base", nil); err != nil {
		http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
	}
}

func FragmentHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if err := tmpl.ExecuteTemplate(w, params["name"], nil); err != nil {
		http.Error(w, "Fragment not found: "+err.Error(), http.StatusInternalServerError)
	}
}
