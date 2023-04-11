package v1

import (
	"net/http"
	"strings"
)

func hasJSONContent(w http.ResponseWriter, r *http.Request) bool {
	if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
		http.Error(w, http.StatusText(http.StatusBadRequest),
			http.StatusBadRequest)
		return false
	}

	return true
}
