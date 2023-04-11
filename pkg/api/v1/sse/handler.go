package sse

import (
	"net/http"
)

type Connection struct {
	Writer  http.ResponseWriter
	Flusher http.Flusher
	Request *http.Request
}

type Handler struct {
	//connections string[string]*Connection
	connections []*Connection
}

func NewHandler() *Handler {
	return &Handler{}
}

// Add a connection to handle
// type: "devices-update", "device-update"
func (h *Handler) Add(t string, w http.ResponseWriter, r *http.Request) (conn *Connection, ok bool) {
	// TODO: check for if event type exists (panic if not)

	h.headers(w)
	f, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, http.StatusText(http.StatusBadRequest),
			http.StatusBadRequest)
		return
	}

	conn = &Connection{
		Writer:  w,
		Flusher: f,
		Request: r,
	}

	// TODO: use map[<event-type>]connection
	h.connections = append(h.connections, conn)

	w.WriteHeader(http.StatusOK)
	return conn, true
}

func (h *Handler) Close(conn *Connection) {
	for i, c := range h.connections {
		if c == conn {
			h.connections = append(h.connections[:i], h.connections[i+1:]...)
			return
		}
	}
}

func (h *Handler) headers(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
