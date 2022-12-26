package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/iMykhailychenko/golang-blog/pkg/config"
	"github.com/iMykhailychenko/golang-blog/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Use(NoSurf)
	r.Use(SessionLoad)

	r.Get("/", handlers.Repo.Home)
	r.Get("/about", handlers.Repo.About)
	r.Get("/join", handlers.Repo.Join)
	r.Get("/login", handlers.Repo.Login)
	r.Get("/new-post", handlers.Repo.NewPost)

	fs := http.FileServer(http.Dir("./static/"))
	r.Handle("/static/*", http.StripPrefix("/static", fs))

	return r
}
