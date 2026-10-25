package groupie

import (
	tools "goupie/Tools"
	"net/http"
	"text/template"
)

const (
	url = "https://groupietrackers.herokuapp.com/api"
)

var tmpl = template.Must(template.ParseFiles("Template/index.html"))

func GroupieHandler(w http.ResponseWriter, r *http.Request) {
	page, err := tools.FetchPage(url)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	if err := tmpl.Execute(w, page); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
