package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/SamMotta/bookings-go/pkg/config"
	"github.com/SamMotta/bookings-go/pkg/handlers"
	"github.com/SamMotta/bookings-go/pkg/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = 8080

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.IsProductionEnv = false

	session = scs.New()

	session.Lifetime = 24 * time.Hour              // How long the session should last, one day.
	session.Cookie.Persist = true                  // Persist the session even after the browser window is closed.
	session.Cookie.SameSite = http.SameSiteLaxMode // SameSiteLaxMode is the default, but we're setting it here for clarity.
	session.Cookie.Secure = app.IsProductionEnv    // Set to true in production.

	app.Session = session

	// Create a template cache
	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatalln(err)
	}

	app.TemplateCache = tc
	app.UseCache = false
	render.NewTemplates(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	var addr = fmt.Sprintf("127.0.0.1:%d", portNumber)

	server := &http.Server{
		Addr:    addr,
		Handler: routes(),
	}

	fmt.Printf("Starting server on address %s\n", addr)
	err = server.ListenAndServe()

	if err != nil {
		log.Fatalln(err)
	}
}
