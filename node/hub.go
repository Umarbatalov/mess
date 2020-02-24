package node

import "github.com/Umarbatalov/mess/model"

type Hub struct {
	clients    map[int]*Client
	send       chan *model.Message
	register   chan *Client
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[int]*Client),
		send:       make(chan *model.Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client.id] = client
		case client := <-h.unregister:
			if _, ok := h.clients[client.id]; ok {
				close(client.send)
				delete(h.clients, client.id)
			}
		case message := <-h.send:
			if client, ok := h.clients[message.Receiver]; ok {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client.id)
				}
			}
		}
	}
}
