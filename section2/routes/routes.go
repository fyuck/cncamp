package routes

import (
	"io"
	"net/http"
	"os"
	"strings"
	
	"github.com/rs/zerolog/log"
)

// HomeHandler 
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Get env VERSION.
	// If VERSION is empty, value of response header will be "emptyVersion".
	version := os.Getenv("VERSION")
	if version == "" {
		w.Header().Add("VERSION", "emptyVersion")
		log.Error().Msg("No VERSION found.")
	}else{
		w.Header().Add("VERSION", version)
	}
	
	// Get request headers and write them to response headers.
	for k, v := range r.Header {
		log.Info().Msg(k)
		w.Header().Add(k, strings.Join(v, ","))
	}
	io.WriteString(w, "Request headers are written in response headers.\n")
	// Get source IP and record in the log.
	sourceIP := r.Header.Get("X-FORWARDED-FOR")
	if sourceIP == "" {
		sourceIP = "unknown"
	}
	log.Info().Msgf("source IP is %v, status code is %v", sourceIP, 200)
}

// HealthCheckHandler simply return ok which is used for health checking.
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok\n")
}