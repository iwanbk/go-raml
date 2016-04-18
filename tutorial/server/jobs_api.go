package main

import (
	"net/http"
)

// JobsAPI is API implementation of /jobs root endpoint
type JobsAPI struct {
}

// Post is the handler for POST /jobs
// Create a job
func (api JobsAPI) Post(w http.ResponseWriter, r *http.Request) {
	// uncomment below line to add header
	// w.Header().Set("key","value")
}
