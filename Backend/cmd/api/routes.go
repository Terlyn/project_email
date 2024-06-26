package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// routes genera nuestras rutas y las adjunta a los manejadores, utilizando el enrutador chi
func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Get("/", app.handlers.Home)
	mux.Get("/search", app.handlers.Search)
	mux.Get("/emails", app.handlers.GetEmails)
	// mux.Post("/emails", )
	mux.Post("/search", app.handlers.SearchEmails) // Nueva ruta POST para buscar correos electrónicos

	return mux
}
