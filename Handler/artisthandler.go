package groupie

import (
	tools "goupie/Tools"
	"net/http"
	"strconv"
	"text/template"
)

var artistTmpl = template.Must(template.ParseFiles("Template/artist.html"))

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	pages, err := tools.FetchPage(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	for _, p := range pages {
		if p.ID == id {
			if err := artistTmpl.Execute(w, p); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}

	http.NotFound(w, r)
}
