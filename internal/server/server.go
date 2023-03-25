package server

import "net/http"

func NewHandler() *http.ServeMux {
	mux := http.NewServeMux()

	// Add handler
	{ // Group: "/api/v1"
		// TODO: "/devices"...
		// ...

		// TODO: "/events"...
		// ...
	}

	return mux
}

func New(addr string) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: NewHandler(),
	}
}
