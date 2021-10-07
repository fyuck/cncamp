package routes

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/golang/glog"
)

// HomeHandler 
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Get env VERSION.
	// If VERSION is empty, value of response header will be "emptyVersion".
	version := os.Getenv("VERSION")
	if version == "" {
		w.Header().Add("VERSION", "emptyVersion")
		io.WriteString(w, "No VERSION found.\n")
	}else{
		w.Header().Add("VERSION", version)
	}
	
	// Get request headers and write them to response headers.
	for k,v := range r.Header {
		w.Header().Add(k, strings.Join(v, ","))
	}
	io.WriteString(w, "Request headers are written in response headers.\n")
	
	// Get source IP and record in the log.
	sourceIP := r.Header.Get("X-FORWARDED-FOR")
	glog.Infof("source IP is %v, status code is %v", sourceIP, 200)
}

// HealthCheckHandler simply return ok which is used for health checking.
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok\n")
}