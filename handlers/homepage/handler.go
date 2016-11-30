package homepage

import (
	"net/http"
	"github.com/ONSdigital/dp-dd-frontend-controller/renderer"
	"github.com/ONSdigital/go-ns/log"
	"github.com/ONSdigital/dp-frontend-models/model/dd/homepage"
)

// Handler handles requests to the homepage
func Handler(w http.ResponseWriter, req *http.Request) {
	page := homepage.Homepage{}

	body, err := renderer.Render(page, "dd/homepage")
	if err != nil {
		log.ErrorR(req, err, nil)
		respond(w, http.StatusInternalServerError, []byte(err.Error()))
		return
	}

	respond(w, http.StatusOK, body)
}

func respond(w http.ResponseWriter, status int, body []byte) {
	w.WriteHeader(status)
	if _, err := w.Write(body); err != nil {
		log.Error(err, nil)
	}
}