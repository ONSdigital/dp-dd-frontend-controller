package config

// BindAddr is the address to bind to when starting the server.
var BindAddr = ":20030"

// RendererURL is the address of the frontend renderer service.
var RendererURL = "http://localhost:20010"

// DiscoveryAPIURL is the address of the data discovery REST API service.
var DiscoveryAPIURL = "http://localhost:20099"

// JobAPIRURL is the address of the data discovery download job creation service
var JobAPIURL = "http://localhost:20100"

// ExternalURL is the base URL through which users are accessing the service.
var ExternalURL = "http://localhost:20000/dd"