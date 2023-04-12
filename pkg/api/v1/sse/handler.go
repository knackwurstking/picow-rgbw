package sse

import (
	"net/http"

	"github.com/knackwurstking/picow-rgbw-web/pkg/log"
)

type Handler struct {
	Connections *Connections
}

func NewHandler() *Handler {
	return &Handler{
		Connections: NewConnections(),
	}
}

// Add a connection to handle
// type: "devices-update", "device-update"
func (h *Handler) Add(event string, w http.ResponseWriter, r *http.Request) (*Connection, bool) {
	h.headers(w)
	f, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, http.StatusText(http.StatusBadRequest),
			http.StatusBadRequest)
		return nil, false
	}

	c := &Connection{
		Writer:  w,
		Flusher: f,
		Request: r,
	}
	h.Connections.Add(event, c)

	w.WriteHeader(http.StatusOK)
	return c, true
}

func (h *Handler) Close(event string, conn *Connection) {
	h.Connections.Remove(event, conn)
}

func (h *Handler) Dispatch(event string, eventType string, data any) {
	for _, c := range h.Connections.Get(event) {
		go func(c *Connection) {
			if err := c.Write(eventType, data); err != nil {
				log.Error.Println(err.Error())
			}
		}(c)
	}
}

func (h *Handler) headers(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
