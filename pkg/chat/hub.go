package chat

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

// return a new instance of hub
func NewHubinstance() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

// this function is responsible for registering, unregistering and broadcasting
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.Send)
			}
		case messgae := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.Send <- messgae:
				default:
					close(client.Send)
					delete(h.clients, client)
				}
			}
		}

	}
}
