package sse

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/knackwurstking/picow-rgbw-web/pkg/log"
)

type Connection struct {
	Writer  http.ResponseWriter
	Flusher http.Flusher
	Request *http.Request

	mutex sync.RWMutex
}

func (c *Connection) Write(event string, data any) error {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	d, err := json.Marshal(data)
	if err != nil {
		return err
	}
	d = append(d, []byte("\n\n")...)

	_, err = c.Writer.Write([]byte("event: " + event + "\n"))
	if err != nil {
		return err
	}
	c.Flusher.Flush()

	_, err = c.Writer.Write(append([]byte("data: "), d...))
	if err != nil {
		return err
	}
	c.Flusher.Flush()

	return nil
}

type Connections struct {
	data map[string][]*Connection
}

func NewConnections() *Connections {
	return &Connections{
		data: make(map[string][]*Connection),
	}
}

func (cs *Connections) Add(e string, c *Connection) {
	log.Debug.Printf("Add sse connection for \"%s\": %s", e, c.Request.RemoteAddr)

	if conns, ok := cs.data[e]; ok {
		cs.data[e] = append(conns, c)
	} else {
		cs.data[e] = []*Connection{c}
	}
}

func (cs *Connections) Get(e string) []*Connection {
	if conns, ok := cs.data[e]; ok {
		return conns
	}

	return []*Connection{}
}

func (cs *Connections) Remove(e string, c *Connection) {
	log.Debug.Printf("remove sse connection for \"%s\": %s", e, c.Request.RemoteAddr)

	if conns, ok := cs.data[e]; ok {
		for i, c2 := range conns {
			if c2 == c {
				cs.data[e] = append(conns[:i], conns[i+1:]...)
				return
			}
		}
	}
}
