package server

import (
	"context"
	"net/http"

	"github.com/knackwurstking/picow-rgbw-web/pkg/api/v1/pico"
	"github.com/knackwurstking/picow-rgbw-web/pkg/middleware"
)

func New(addr string, picoHandler *pico.Handler) *http.Server {
	UseMiddleware(middleware.NewLogger)
	UseMiddleware(middleware.NewCORS)

	// TODO: use custom type as key (replace "pico")
	ctx := context.WithValue(context.Background(), "pico", picoHandler)
	return &http.Server{
		Addr:    addr,
		Handler: NewRegExHandler(ctx),
	}
}
