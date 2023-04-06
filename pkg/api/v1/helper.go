package v1

import (
	"context"
	"net/http"
	"strings"

	"github.com/gookit/slog"
	"github.com/knackwurstking/picow-rgbw-web/pkg/api/v1/pico"
)

func getHandler(w http.ResponseWriter, ctx context.Context) (handler *pico.Handler, ok bool) {
	h := ctx.Value("pico")
	if h == nil {
		slog.Fatal("Server context value for \"pico\" is missing (*pico.Handler)")
		http.Error(w, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return nil, false
	}
	return h.(*pico.Handler), true
}

func hasJSONContent(w http.ResponseWriter, r *http.Request) bool {
	if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
		http.Error(w, http.StatusText(http.StatusBadRequest),
			http.StatusBadRequest)
		return false
	}

	return true
}
