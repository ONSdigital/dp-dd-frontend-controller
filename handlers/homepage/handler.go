package homepage

import (
	"github.com/ONSdigital/dp-dd-frontend-controller/config"
	"github.com/ONSdigital/dp-dd-frontend-controller/discovery"
	"github.com/ONSdigital/dp-dd-frontend-controller/renderer"
	"github.com/ONSdigital/dp-frontend-models/model/dd/homepage"
	"github.com/ONSdigital/go-ns/log"
	"net/http"
)

// Handler handles requests to the homepage
func Handler(w http.ResponseWriter, req *http.Request) {

	// Call into DD API to get a list of all datasets
	datasets, err := discovery.ListDatasets()
	if err != nil {
		log.Error(err, nil)
		respond(w, http.StatusInternalServerError, []byte(err.Error()))
		return
	}

	// Rewrite the URLs in the datasets to point to our own address
	for _, dataset := range datasets.Items {
		dataset.URL = config.ExternalURL + "/dataset/" + dataset.ID
	}

	page := homepage.Homepage{
		Datasets: datasets,
	}

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
