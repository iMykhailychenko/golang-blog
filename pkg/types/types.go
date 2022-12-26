package types

import "html/template"

type HtmlCacheType = map[string]*template.Template

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	CSRFToken string
	Message   string
	Error     string
	Warning   string
}
