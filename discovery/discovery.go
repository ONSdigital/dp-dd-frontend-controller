package discovery

import (
	"encoding/json"
	"fmt"
	"github.com/ONSdigital/dp-dd-frontend-controller/config"
	"github.com/ONSdigital/dp-frontend-models/model/dd"
	"github.com/ONSdigital/go-ns/log"
	"io"
	"io/ioutil"
	"net/http"
)

// ListDatasets lists the available datasets by querying the DD API.
func ListDatasets() (datasets *dd.Datasets, err error) {
	request, err := http.NewRequest("GET", config.DiscoveryAPIURL+"/datasets", nil)
	if err != nil {
		log.Error(err, nil)
		return
	}

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		log.ErrorR(request, err, nil)
		return
	}
	defer checkClose(res.Body)

	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("discovery.ListDatasets: unexpected status code from API: %d", res.StatusCode)
		return
	}

	datasetsJSON, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.ErrorR(request, err, nil)
		return
	}

	if err = json.Unmarshal(datasetsJSON, &datasets); err != nil {
		log.ErrorR(request, err, nil)
		return
	}

	return
}

// GetDataset retrieves metadata and dimension info for a dataset.
func GetDataset(id string) (dataset *dd.Dataset, err error) {
	request, err := http.NewRequest("GET", config.DiscoveryAPIURL+"/datasets/"+id, nil)
	if err != nil {
		log.Error(err, nil)
		return
	}

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		log.ErrorR(request, err, nil)
		return
	}
	defer checkClose(res.Body)

	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("discovery.GetDataset: unexpected status code from API: %d", res.StatusCode)
		return
	}

	datasetJSON, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.ErrorR(request, err, nil)
		return
	}

	if err = json.Unmarshal(datasetJSON, &dataset); err != nil {
		log.ErrorR(request, err, nil)
		return
	}

	return

}

func checkClose(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Error(err, nil)
	}
}
