package unit

// Хаб - регистрирует в себе клиентов
type Hub struct {
	clients    map[int]*Client
	send       chan *Message
	register   chan *Client
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[int]*Client),
		send:       make(chan *Message),
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
			if client, ok := h.clients[message.Client_id]; ok {
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
