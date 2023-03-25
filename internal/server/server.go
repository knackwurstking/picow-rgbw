package server

import "net/http"

func NewHandler() *http.ServeMux {
	mux := http.NewServeMux()

	// TODO: Add handler

	return mux
}

func New(addr string) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: NewHandler(),
	}
}
