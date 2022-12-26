package main

import (
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

func InitSession() {
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session
}
