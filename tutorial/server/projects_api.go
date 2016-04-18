package main

import (
	"net/http"
)

// ProjectsAPI is API implementation of /projects root endpoint
type ProjectsAPI struct {
}

// Post is the handler for POST /projects
// Create a project
func (api ProjectsAPI) Post(w http.ResponseWriter, r *http.Request) {
	// uncomment below line to add header
	// w.Header().Set("key","value")
}
