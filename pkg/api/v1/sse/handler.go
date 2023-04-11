package sse

import (
	"net/http"

	"github.com/knackwurstking/picow-rgbw-web/pkg/api/v1/pico"
)

type Connection struct {
	Request *http.Request
}

func NewConnection() *Connection {
	return &Connection{}
}

type Handler struct {
	connections []*Connection
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request, p *pico.Handler) (conn *Connection, ok bool) {
	h.headers(w)

	// TODO: create connection

	conn = NewConnection()
	h.connections = append(h.connections, conn)

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
