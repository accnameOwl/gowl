package handler

import (
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

func init() {
	// set default path
	htmltmpl = template.Must(template.ParseGlob("./public/*.html"))
}

func Home(response http.ResponseWriter, request *http.Request, param httprouter.Params) {

}
