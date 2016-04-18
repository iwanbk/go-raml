package main

import (
	"net/http"
	"fmt"
)

// ResourcesAPI is API implementation of /resources root endpoint
type ResourcesAPI struct {
}

// Get is the handler for GET /resources
// Get a resource
func (api ResourcesAPI) Get(w http.ResponseWriter, r *http.Request) {
	// uncomment below line to add header
	// w.Header().Set("key","value")
	fmt.Fprintf(w, "response from get call")
}

// Post is the handler for POST /resources
// Post a resource
func (api ResourcesAPI) Post(w http.ResponseWriter, r *http.Request) {
	// uncomment below line to add header
	// w.Header().Set("key","value")
	fmt.Fprintf(w, "response from post call ")
}

// resourceIdGet is the handler for GET /resources/{resourceId}
// Get a resource ID
func (api ResourcesAPI) resourceIdGet(w http.ResponseWriter, r *http.Request) {
	// uncomment below line to add header
	// w.Header().Set("key","value")
	fmt.Fprintf(w, "this should be the resource with the requested id")
}

// resourceIdPut is the handler for PUT /resources/{resourceId}
// Put resource ID
func (api ResourcesAPI) resourceIdPut(w http.ResponseWriter, r *http.Request) {
	// uncomment below line to add header
	// w.Header().Set("key","value")
	fmt.Fprintf(w, "response from resourceIdput")
}

// resourceIdDelete is the handler for DELETE /resources/{resourceId}
// Delete a resource ID
func (api ResourcesAPI) resourceIdDelete(w http.ResponseWriter, r *http.Request) {
	// uncomment below line to add header
	// w.Header().Set("key","value")
	fmt.Fprintf(w, "response from resource IdDelete")
}

// resourceIdyet_anotherGet is the handler for GET /resources/{resourceId}/yet_another
// Yet another
func (api ResourcesAPI) resourceIdyet_anotherGet(w http.ResponseWriter, r *http.Request) {
	// uncomment below line to add header
	// w.Header().Set("key","value")
	fmt.Fprintf(w, "response from resourceIdyet_anotherGet")
}
