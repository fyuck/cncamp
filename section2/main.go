package main

import (
	"flag"
	// "fmt"
	"net/http"

	"cncamp/section2/routes"
	"github.com/gorilla/mux"
	"github.com/golang/glog"
)

func main() {
	setupLogger()
	startServer()
}

func startServer() {
	r := mux.NewRouter()

	r.HandleFunc("/healthz", routes.HealthCheckHandler)
	r.PathPrefix("/").HandlerFunc(routes.HomeHandler)

	err := http.ListenAndServe("0.0.0.0:80", r)
	if err != nil {
		//fmt.Println("Failed to start server: %v", err)
	}
}

func setupLogger() {
	flag.Set("v", "4")
	glog.V(2).Info("starting http server...")
	defer glog.Flush()
}