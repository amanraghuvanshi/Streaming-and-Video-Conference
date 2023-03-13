package chat

import (
	"bytes"
	"log"
	"time"

	"github.com/fasthttp/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

var (
	newLine = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.FastHTTPUpgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	Hub        *Hub
	Connection *websocket.Conn
	Send       chan []byte
}

// we need three functions here
func (c *Client) ReadPump() {
	defer func() {
		c.Hub.unregister <- c
		c.Connection.Close()
	}()

	c.Connection.SetReadLimit(maxMessageSize)
	c.Connection.SetReadDeadline(time.Now().Add(pongWait))
	c.Connection.SetPongHandler(func(appData string) error {
		c.Connection.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	for {
		_, message, err := c.Connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseAbnormalClosure) {
				log.Printf("Error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newLine, space, -1))
		c.Hub.broadcast <- message
	}
}

func (c *Client) WritePump() {

}

func PeerChatConn(c *websocket.Conn, hub *Hub) {
	client := &Client{
		Hub:        hub,
		Connection: c,
		Send:       make(chan []byte, 256),
	}
	client.Hub.register <- client

	go client.WritePump()
	client.ReadPump()
}

// "github.com/fasthttp/websocket"
