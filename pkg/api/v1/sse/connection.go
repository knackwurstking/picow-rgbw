package sse

import (
	"encoding/json"
	"net/http"
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
