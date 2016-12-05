package main

import (
	"net/http"
	"os"
	"time"

	"github.com/ONSdigital/dp-dd-frontend-controller/config"
	"github.com/ONSdigital/dp-dd-frontend-controller/handlers/dataset"
	"github.com/ONSdigital/dp-dd-frontend-controller/handlers/homepage"
	"github.com/ONSdigital/go-ns/handlers/requestID"
	"github.com/ONSdigital/go-ns/handlers/timeout"
	"github.com/ONSdigital/go-ns/log"
	"github.com/gorilla/pat"
	"github.com/justinas/alice"
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

	router.HandleFunc("/dd", homepage.Handler)
	router.HandleFunc("/dd/", homepage.Handler)
	router.Get("/dd/dataset/{id}", dataset.Handler)

	log.Debug("Starting server", log.Data{
		"bind_addr":         config.BindAddr,
		"renderer_url":      config.RendererURL,
		"discovery_api_url": config.DiscoveryAPIURL,
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
