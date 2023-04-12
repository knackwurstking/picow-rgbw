package sse

import (
	"net/http"

	"github.com/gookit/slog"
)

type Handler struct {
	connections map[string][]*Connection
}

func NewHandler() *Handler {
	return &Handler{
		connections: make(map[string][]*Connection),
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
	h.add(event, c)

	w.WriteHeader(http.StatusOK)
	return c, true
}

func (h *Handler) Close(event string, conn *Connection) {
	h.remove(event, conn)
}

func (h *Handler) Dispatch(event string, data any) {
	if conns, ok := h.connections[event]; ok {
		for _, c := range conns {
			go func(c *Connection) {
				if err := c.Write(event, data); err != nil {
					slog.Error(err.Error())
				}
			}(c)
		}
	}
}

func (h *Handler) add(event string, conn *Connection) {
	slog.Debug("add sse connection:", conn.Request.RemoteAddr)
	if conns, ok := h.connections[event]; ok {
		h.connections[event] = append(conns, conn)
	} else {
		h.connections[event] = []*Connection{conn}
	}
}

func (h *Handler) remove(event string, conn *Connection) {
	slog.Debug("remove sse connection:", conn.Request.RemoteAddr)
	if conns, ok := h.connections[event]; ok {
		for i, c2 := range conns {
			if c2 == conn {
				h.connections[event] = append(conns[:i], conns[i+1:]...)
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
