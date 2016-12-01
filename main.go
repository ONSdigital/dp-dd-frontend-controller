package main

import (
	"net/http"
	"os"
	"time"

	"github.com/ONSdigital/dp-dd-frontend-controller/config"
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

	log.Namespace = "dp-dd-frontend-controller"

	router := pat.New()
	aliceHandler := alice.New(
		timeout.Handler(10*time.Second),
		log.Handler,
		requestID.Handler(16),
	).Then(router)

	router.HandleFunc("/dd", homepage.Handler)
	router.HandleFunc("/dd/", homepage.Handler)

	log.Debug("Starting server", log.Data{
		"bind_addr": config.BindAddr,
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
