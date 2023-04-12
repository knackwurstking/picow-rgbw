package sse

import (
	"encoding/json"
	"net/http"

	"github.com/gookit/slog"
)

type Connection struct {
	Writer  http.ResponseWriter
	Flusher http.Flusher
	Request *http.Request
}

func (c *Connection) Write(event string, data any) error {
	d, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = c.Writer.Write([]byte("event: " + event + "\n"))
	if err != nil {
		return err
	}

	_, err = c.Writer.Write(append([]byte("data: "), d...))

	return err
}

type Connections map[string][]*Connection

func (cs Connections) Add(e string, c *Connection) {
	slog.Debug("add sse connection:", c.Request.RemoteAddr)

	if conns, ok := cs[e]; ok {
		cs[e] = append(conns, c)
	} else {
		cs[e] = []*Connection{c}
	}
}

func (cs Connections) Get(e string) []*Connection {
	if conns, ok := cs[e]; ok {
		return conns
	}

	return []*Connection{}
}

func (cs Connections) Remove(e string, c *Connection) {
	slog.Debug("remove sse connection:", c.Request.RemoteAddr)

	if conns, ok := cs[e]; ok {
		for i, c2 := range conns {
			if c2 == c {
				cs[e] = append(conns[:i], conns[i+1:]...)
				return
			}
		}
	}
}
