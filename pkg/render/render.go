package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/iMykhailychenko/golang-blog/pkg/config"
	"github.com/iMykhailychenko/golang-blog/pkg/types"
)

var app *config.AppConfig

func ApplyCacheConfig(a *config.AppConfig) {
	app = a
}

func NewHtmlCache() types.HtmlCacheType {
	c := types.HtmlCacheType{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		log.Fatal(err)
	}

	for _, page := range pages {
		name := filepath.Base(page)

		tmpl, err := template.New(name).ParseFiles(page)
		if err != nil {
			log.Fatal(err)
		}

		layouts, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			log.Fatal(err)
		}

		if len(layouts) > 0 {
			tmpl, err = tmpl.ParseGlob("./templates/*.layout.html")

			if err != nil {
				log.Fatal(err)
			}
		}

		c[name] = tmpl
	}

	return c
}

func getDefaultData(td *types.TemplateData) *types.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *types.TemplateData) {
	var tt types.HtmlCacheType

	if app.InProduction {
		tt = app.HtmlCache
	} else {
		tt = NewHtmlCache()
	}

	t, ok := tt[tmpl]
	if !ok {
		log.Fatalf(fmt.Sprintf("Template %s not found\n", tmpl))
	}

	buf := new(bytes.Buffer)
	err := t.Execute(buf, getDefaultData(td))
	if err != nil {
		log.Fatal(err)
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Fatal(err)
	}
}
