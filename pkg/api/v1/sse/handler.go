package sse

type Connection struct {
}

func NewConnection() Connection {
	return Connection{}
}

type Handler struct {
	connections []Connection
}

func NewHandler() *Handler {
	return &Handler{}
}
