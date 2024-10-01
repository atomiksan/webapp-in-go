package handlers

import (
	"net/http"

	"github.com/atomiksan/webapp-in-go/pkg/config"
	"github.com/atomiksan/webapp-in-go/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the epository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the homepage handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
