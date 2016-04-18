package main

//This file is auto-generated by go-raml
//Do not edit this file by hand since it will be overwritten during the next generation

import (
	"github.com/gorilla/mux"
	"net/http"
)

// ProjectsInterface is interface for /projects root endpoint
type ProjectsInterface interface { // Post is the handler for POST /projects
	// Create a project
	Post(http.ResponseWriter, *http.Request)
}

// ProjectsInterfaceRoutes is routing for /projects root endpoint
func ProjectsInterfaceRoutes(r *mux.Router, i ProjectsInterface) {
	r.HandleFunc("/projects", i.Post).Methods("POST")
}
