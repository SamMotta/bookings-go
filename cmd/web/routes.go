package main

import (
	"github.com/SamMotta/bookings-go/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func routes() http.Handler {
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	dir := "./static"

	fileServer := http.FileServer(
		http.Dir(dir),
	)

	mux.Handle(
		"/static/*",
		http.StripPrefix("/static", fileServer),
	)

	return mux
}
