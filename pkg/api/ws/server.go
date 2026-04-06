package ws

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// Server handles real-time WebSocket connections.
type Server struct {
	mu      sync.Mutex
	clients map[*websocket.Conn]bool
}

// NewServer creates a new WebSocket server.
func NewServer() *Server {
	return &Server{
		clients: make(map[*websocket.Conn]bool),
	}
}

// Handler handles new WebSocket connections.
func (s *Server) Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	s.mu.Lock()
	s.clients[conn] = true
	s.mu.Unlock()

	defer func() {
		s.mu.Lock()
		delete(s.clients, conn)
		s.mu.Unlock()
	}()

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
	}
}

// Broadcast sends a message to all connected WebSocket clients.
func (s *Server) Broadcast(msg interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for client := range s.clients {
		client.WriteJSON(msg)
	}
}
