package handlers

import (
	"fmt"
	"net/http"
)

// Handlers struct will have methods for handling various routes
type Handlers struct {
	// Add any dependencies your handlers might need, e.g. a database connection
}

// NewHandlers returns a new Handlers struct
func NewHandlers() *Handlers {
	return &Handlers{}
}

// Home is the handler for the default route
func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the home page!")
}

// Search is the handler for the /search route
func (h *Handlers) Search(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Search functionality goes here.")
}
