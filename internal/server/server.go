package server

import "net/http"

func NewHandler() *http.ServeMux {
	mux := http.NewServeMux()

	// TODO: Add handler (groups?: "/api/v1", "/api/v1/devices", "/api/v1/events")

	return mux
}

func New(addr string) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: NewHandler(),
	}
}
