package dataset

import (
	"net/http"

	"github.com/ONSdigital/dp-dd-frontend-controller/config"
	"github.com/ONSdigital/dp-dd-frontend-controller/discovery"
	"github.com/ONSdigital/dp-dd-frontend-controller/renderer"
	"github.com/ONSdigital/dp-frontend-models/model/dd/dataset"
	"github.com/ONSdigital/go-ns/log"
)

// Handler handles requests to the homepage
func Handler(w http.ResponseWriter, req *http.Request) {

	id := req.URL.Query().Get(":id")

	// Call into DD API to get the dataset information
	datasetModel, err := discovery.GetDataset(id)
	if err != nil {
		log.Error(err, nil)
		respond(w, http.StatusNotFound, []byte(err.Error()))
		return
	}
	log.DebugR(req, `Got response from API`, log.Data{"datasetModel": datasetModel})

	// Rewrite the URLs in the datasets to point to our own address
	datasetModel.URL = config.ExternalURL + "/versions/" + datasetModel.ID

	page := dataset.Page{
		Dataset: datasetModel,
	}

	body, err := renderer.Render(page, "dd/dataset")
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
