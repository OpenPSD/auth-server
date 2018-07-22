package api

import (
	"net/http"
	"time"
)

func (s Server) timeoutHandler(h http.Handler) http.Handler {
	return http.TimeoutHandler(h, 2*time.Second, "timed out")
}
