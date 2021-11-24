package main

import (
	"net/http"

	"github.com/TheCyclistGoServer/handlers"
)

func router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/api/data/location", handlers.LocationDataHandler)
	return mux
}
