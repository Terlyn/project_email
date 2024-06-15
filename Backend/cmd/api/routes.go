package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// routes generates our routes and attaches them to handlers, using the chi router
func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)

	mux.Get("/", app.handlers.Home)
	mux.Get("/search", app.handlers.Search)

	return mux
}
