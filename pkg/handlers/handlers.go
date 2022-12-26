package handlers

import (
	"net/http"

	"github.com/iMykhailychenko/golang-blog/pkg/config"
	"github.com/iMykhailychenko/golang-blog/pkg/render"
	"github.com/iMykhailychenko/golang-blog/pkg/types"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func InitRepository(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func ApplyRepository(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html", &types.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html", &types.TemplateData{})
}


func (m *Repository) Join(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "join.page.html", &types.TemplateData{})
}


func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "login.page.html", &types.TemplateData{})
}

func (m *Repository) NewPost(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "new-post.page.html", &types.TemplateData{})
}
