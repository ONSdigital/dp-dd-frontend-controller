package discovery

import (
	"github.com/ONSdigital/dp-dd-frontend-controller/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ONSdigital/dp-frontend-models/model/dd"
	. "github.com/smartystreets/goconvey/convey"
)

func TestListDatasets(t *testing.T) {
	var json []byte
	statusCode := http.StatusOK

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		w.Write(json)
	}))
	defer ts.Close()

	config.DiscoveryAPIURL = ts.URL

	Convey("ListDatasets returns empty list if no datasets registered", t, func() {
		statusCode = http.StatusOK
		json = []byte(`{"items":[],"total":0}`)
		datasets, err := ListDatasets()
		So(err, ShouldBeNil)
		So(datasets, ShouldNotBeNil)
		So(datasets.Items, ShouldBeEmpty)
	})

	Convey("ListDatasets returns accurate dataset metadata", t, func() {
		statusCode = http.StatusOK
		json = []byte(`{"items":[{"datasetID":"one","title":"title one","url":"url1", "metadata":{"description":"test description"}}],"total":1}`)
		datasets, err := ListDatasets()
		So(err, ShouldBeNil)
		So(datasets, ShouldNotBeNil)
		So(datasets.Items, ShouldHaveLength, 1)
		So(datasets.Items[0], ShouldResemble, &dd.Dataset{
			ID:    "one",
			Title: "title one",
			URL:   "url1",
			Metadata: &dd.Metadata{
				Description: "test description",
			},
		})
	})

	Convey("ListDatasets returns error if API is unavailable", t, func() {
		statusCode = http.StatusServiceUnavailable
		_, err := ListDatasets()
		So(err, ShouldNotBeNil)
	})

	Convey("ListDatasets returns error if API is unreachable", t, func() {
		config.DiscoveryAPIURL = "an-unreachable-address"
		defer func() { config.DiscoveryAPIURL = ts.URL }()
		_, err := ListDatasets()
		So(err, ShouldNotBeNil)
	})
}

func TestGetDataset(t *testing.T) {
	var json []byte
	statusCode := http.StatusOK

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		w.Write(json)
	}))
	defer ts.Close()

	config.DiscoveryAPIURL = ts.URL

	Convey("GetDataset returns error if dataset does not exist", t, func() {
		statusCode = http.StatusNotFound
		_, err := GetDataset("foo")
		So(err, ShouldNotBeNil)
	})

	Convey("GetDataset returns accurate dataset information", t, func() {
		statusCode = http.StatusOK
		json = []byte(`{"datasetID":"one","title":"title one","url":"url1","metadata":{"description":"test description"},"dimensions":[{"id":"a","name":"A"}]}`)
		dataset, err := GetDataset("one")
		So(err, ShouldBeNil)
		So(dataset, ShouldNotBeNil)
		var dimensions []*dd.Dimension
		dimensions = make([]*dd.Dimension, 1)
		dimensions[0] = &dd.Dimension{
			ID:   "a",
			Name: "A",
		}
		So(dataset, ShouldResemble, &dd.Dataset{
			ID:    "one",
			Title: "title one",
			URL:   "url1",
			Metadata: &dd.Metadata{
				Description: "test description",
			},
			Dimensions: dimensions,
		})
	})

	Convey("GetDataset returns error if API is unavailable", t, func() {
		statusCode = http.StatusServiceUnavailable
		_, err := GetDataset("test")
		So(err, ShouldNotBeNil)
	})

	Convey("ListDatasets returns error if API is unreachable", t, func() {
		config.DiscoveryAPIURL = "an-unreachable-address"
		defer func() { config.DiscoveryAPIURL = ts.URL }()
		_, err := GetDataset("test")
		So(err, ShouldNotBeNil)
	})
}
