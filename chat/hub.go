package chat

import (
	"fmt"
	"strconv"
	"time"
)

//credit for https://hoohoo.top/blog/20220320172715-go-websocket/

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan Message

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	fmt.Println("Hub started!")
	for {
		select {
		case client := <-h.register:
			_, ok := h.clients[client]
			if client.verified && !ok {

				clientId := client.ID
				for client := range h.clients {
					msg := Message{Sender: 777, SendAt: time.Now(), Text: "Some one connected: " + strconv.Itoa(clientId)}
					client.send <- msg
				}

				h.clients[client] = true
				fmt.Println(h.clients[client])
			}

		case client := <-h.unregister:
			clientId := client.ID
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
			for client := range h.clients {
				msg := Message{Sender: 777, SendAt: time.Now(), Text: "Some one leave: " + strconv.Itoa(clientId)}
				client.send <- msg
			}
		case userMessage := <-h.broadcast:

			for client := range h.clients {
				//prevent self receive the message
				fmt.Printf("Client: %v, Sender: %v", client.ID, userMessage.Sender)

				if client.ID == userMessage.Sender {
					continue
				}
				select {
				case client.send <- userMessage:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
