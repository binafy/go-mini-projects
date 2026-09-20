package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Hub struct {
	clients map[*websocket.Conn]bool
	mu      sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[*websocket.Conn]bool),
	}
}

func main() {
	hub := NewHub()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("Error upgrading to websocket: %v", err)
			return
		}
		defer conn.Close()

		// Add client
		hub.AddClient(conn)

		// Remove client when disconnected
		defer hub.RemoveClient(conn)

		for {
			// Read messages
			msgType, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("client disconnected:", err)
				break
			}
			if msgType != websocket.TextMessage {
				continue
			}

			fmt.Printf("Received: %s\n", message)

			// Broadcast message to all clients
			hub.Broadcast(message)
		}
	})

	fmt.Printf("Listening on port 8080\n")
	http.ListenAndServe(":8080", nil)
}

func (h *Hub) AddClient(conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.clients[conn] = true
}

func (h *Hub) RemoveClient(conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()

	delete(h.clients, conn)
}

func (h *Hub) Broadcast(message []byte) {
	h.mu.Lock()
	defer h.mu.Unlock()

	for client := range h.clients {
		err := client.WriteMessage(websocket.TextMessage, message)

		if err != nil {
			log.Println("write error:", err)

			client.Close()
			delete(h.clients, client)
		}
	}
}
