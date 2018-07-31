package api

import (
	"log"
	"net/http"
	"time"
)

func (s Server) timeoutHandler(h http.Handler) http.Handler {
	return http.TimeoutHandler(h, 2*time.Second, "timed out")
}

func (s Server) loggingHandler(next http.Handler) http.Handler {

	// Define the http.HandlerFunc
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Do middleware things
		start := time.Now()
		defer func() { log.Printf("path=%s time=%s", r.URL.Path, time.Since(start)) }()
		next.ServeHTTP(w, r)
	})
}
