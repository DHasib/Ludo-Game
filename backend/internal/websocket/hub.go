package websocket

import (
	"log"
	sync "sync"

	"github.com/gorilla/websocket"
)

type Hub struct {
	clients map[*websocket.Conn]bool
	mutex   sync.Mutex
}

var hub = Hub{clients: make(map[*websocket.Conn]bool)}

// Broadcast sends message to all connected clients.
func Broadcast(message []byte) {
	hub.mutex.Lock()
	defer hub.mutex.Unlock()
	for c := range hub.clients {
		if err := c.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Println("write error", err)
			c.Close()
			delete(hub.clients, c)
		}
	}
}
