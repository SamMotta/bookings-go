package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/SamMotta/bookings-go/pkg/config"
	"github.com/SamMotta/bookings-go/pkg/models"
)

var app *config.AppConfig

// Sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	// Get the template cache
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// Get request template from cache
	t, ok := tc[tmpl]

	if !ok {
		log.Fatalln("Could not get template from cache.")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	err := t.Execute(buf, td)

	if err != nil {
		log.Fatalln(err)
	}

	// Render the template
	_, err = buf.WriteTo(w)

	log.Printf("%s Template rendered!", tmpl)

	if err != nil {
		log.Fatalln(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// Get all files that end with .page.gohtml from templates folder
	var dir string

	if runtime.GOOS == "windows" {
		dir = "templates"
	} else {
		dir = "../../templates"
	}

	pages, err := filepath.Glob(dir + "/*.page.gohtml")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob(dir + "/*.layout.gohtml")

		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(dir + "/*.layout.gohtml")

			if err != nil {
				return myCache, err
			}
		}

		myCache[strings.TrimSuffix(name, ".gohtml")] = ts
	}

	return myCache, nil
}
