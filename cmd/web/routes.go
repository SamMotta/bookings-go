package main

import (
	"net/http"
	"runtime"

	"github.com/SamMotta/bookings-go/pkg/config"
	"github.com/SamMotta/bookings-go/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	var dir string
	if runtime.GOOS == "windows" {
		dir = "./static"
	} else {
		dir = "../../static"
	}

	fileServer := http.FileServer(
		http.Dir(dir),
	)

	mux.Handle(
		"/static/*",
		http.StripPrefix("/static", fileServer),
	)

	return mux
}
