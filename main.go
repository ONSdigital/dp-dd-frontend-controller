package main

import (
	"net/http"
	"os"
	"time"

	"github.com/ONSdigital/dp-dd-frontend-controller/config"
	"github.com/ONSdigital/dp-dd-frontend-controller/handlers/dataset"
	"github.com/ONSdigital/dp-dd-frontend-controller/handlers/datasetList"
	"github.com/ONSdigital/go-ns/handlers/requestID"
	"github.com/ONSdigital/go-ns/handlers/reverseProxy"
	"github.com/ONSdigital/go-ns/handlers/timeout"
	"github.com/ONSdigital/go-ns/log"
	"github.com/gorilla/pat"
	"github.com/justinas/alice"
	"net/url"
	"strings"
)

func main() {
	if v := os.Getenv("BIND_ADDR"); len(v) > 0 {
		config.BindAddr = v
	}

	if v := os.Getenv("RENDERER_URL"); len(v) > 0 {
		config.RendererURL = v
	}

	if v := os.Getenv("DISCOVERY_API_URL"); len(v) > 0 {
		config.DiscoveryAPIURL = v
	}

	if v := os.Getenv("JOB_API_URL"); len(v) > 0 {
		config.JobAPIURL = v
	}

	if v := os.Getenv("EXTERNAL_URL"); len(v) > 0 {
		config.ExternalURL = v
	}

	log.Namespace = "dp-dd-frontend-controller"

	router := pat.New()
	aliceHandler := alice.New(
		timeout.Handler(10*time.Second),
		log.Handler,
		requestID.Handler(16),
	).Then(router)

	apiURL, err := url.Parse(config.DiscoveryAPIURL)
	if err != nil {
		log.Error(err, nil)
		os.Exit(1)
	}

	apiProxy := reverseProxy.Create(apiURL, func(req *http.Request) {
		req.URL.Path = strings.TrimPrefix(req.URL.Path, `/dd/api`)
	})

	jobApiUrl, err := url.Parse(config.JobAPIURL)
	if err != nil {
		log.Error(err, nil)
		os.Exit(1)
	}

	jobApiProxy := reverseProxy.Create(jobApiUrl, func(req *http.Request) {
		log.Debug(`Job api requested`, nil)
		log.Debug(req.URL.Path, nil)
		req.URL.Path = strings.TrimPrefix(req.URL.Path, `/dd/api/jobs`)
	})

	router.HandleFunc("/dd", datasetList.Handler)
	router.HandleFunc("/dd/", datasetList.Handler)
	router.Handle("/dd/api/jobs{uri:(|/.*)}", jobApiProxy)
	router.Handle("/dd/api{uri:(|/.*)}", apiProxy)
	router.Get("/dd/datasets/{id}", dataset.Handler)

	log.Debug("Starting server", log.Data{
		"bind_addr":             config.BindAddr,
		"renderer_url":          config.RendererURL,
		"discovery_api_url":     config.DiscoveryAPIURL,
		"discovery_job_api_url": config.JobAPIURL,
	})

	server := &http.Server{
		Addr:         config.BindAddr,
		Handler:      aliceHandler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Error(err, nil)
		os.Exit(2)
	}
}
