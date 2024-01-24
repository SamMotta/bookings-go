package handlers

import "github.com/SamMotta/bookings-go/pkg/config"

iRendermport (
	"net/http"

	"github.com/SamMotta/bookings-go/pkg/config"
	"github.com/SamMotta/bookings-go/pkg/models"
	"github.com/SamMotta/bookings-go/pkg/render"
)

var Repo *Repository

// Repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo sets the repository for the handlers
func NewRepo(app *config.AppConfig) *Repository {
	return &Repository{
		App: app,
	}
}

// NewHandlers Sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr

	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again!"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page", &models.TemplateData{StringMap: stringMap})
}
