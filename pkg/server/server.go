package server

import (
	"net/http"

	"github.com/knackwurstking/picow-rgbw-web/pkg/api/v1/pico"
	"github.com/knackwurstking/picow-rgbw-web/pkg/middleware"
)

func New(addr string, p *pico.Handler) *http.Server {
	UseMiddleware(middleware.NewLogger)
	UseMiddleware(middleware.NewCORS)

	return &http.Server{
		Addr:    addr,
		Handler: NewRegExHandler(p),
	}
}
