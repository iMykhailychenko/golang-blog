package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/iMykhailychenko/golang-blog/pkg/config"
	"github.com/iMykhailychenko/golang-blog/pkg/handlers"
	"github.com/iMykhailychenko/golang-blog/pkg/render"
)

const port = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	app.InProduction = false // change this value if you are in production mode

	c := render.NewHtmlCache()
	app.HtmlCache = c

	render.ApplyCacheConfig(&app)
	r := handlers.InitRepository(&app)
	handlers.ApplyRepository(r)

	// set up the session
	InitSession()

	// run server
	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	fmt.Println("Startin server on http://localhost" + port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
