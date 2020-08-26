package handler

import (
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

var responseTemplate *template.Template

func init() {
	// set default path
	responseTemplate = template.Must(template.ParseGlob("./public/*.html"))
}

// Home ...
// redirect to home.html
func Home(res http.ResponseWriter, req *http.Request, param httprouter.Params) {
	responseTemplate.ExecuteTemplate(res, "home.html", nil)
}

// About ...
// redirect to about.html
func About(res http.ResponseWriter, req *http.Request, param httprouter.Params) {
	responseTemplate.ExecuteTemplate(res, "about.html", nil)
}
