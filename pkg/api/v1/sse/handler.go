package sse

import (
	"net/http"

	"github.com/gookit/slog"
)

type Connection struct {
	Writer  http.ResponseWriter
	Flusher http.Flusher
	Request *http.Request
}

type Handler struct {
	//connections string[string]*Connection
	connections map[string][]*Connection
}

func NewHandler() *Handler {
	return &Handler{}
}

// Add a connection to handle
// type: "devices-update", "device-update"
func (h *Handler) Add(t string, w http.ResponseWriter, r *http.Request) (conn *Connection, ok bool) {
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
	h.add(t, conn)

	w.WriteHeader(http.StatusOK)
	return conn, true
}

func (h *Handler) Close(t string, c *Connection) {
	h.remove(t, c)
}

func (h *Handler) Dispatch(t string, data any) {
	// TODO: iter connection for event [t]ype and write to connection
}

func (h *Handler) add(t string, c *Connection) {
	slog.Debug("add sse connection:", c.Request.RemoteAddr)
	if conns, ok := h.connections[t]; ok {
		h.connections[t] = append(conns, c)
	} else {
		h.connections[t] = []*Connection{c}
	}
}

func (h *Handler) remove(t string, c *Connection) {
	slog.Debug("remove sse connection:", c.Request.RemoteAddr)
	if conns, ok := h.connections[t]; ok {
		for i, c2 := range conns {
			if c2 == c {
				h.connections[t] = append(conns[:i], conns[i+1:]...)
				return
			}
		}
	}
}

func (h *Handler) headers(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
