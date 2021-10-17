package main

import (
	"flag"
	"net/http"

	"cncamp/section2/routes"
	
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"

)

func main() {
	flag.Parse()
	startServer()
}

func startServer() {
	r := mux.NewRouter()

	r.HandleFunc("/healthz", routes.HealthCheckHandler)
	r.PathPrefix("/").HandlerFunc(routes.HomeHandler)

	log.Info().Msg("Service is starting to listen on http://localhost/")
	err := http.ListenAndServe("0.0.0.0:80", r)
	if err != nil {
		log.Info().Msg("Start server failed...")
	}
}
