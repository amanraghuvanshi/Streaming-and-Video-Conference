package chat

import (
	"time"

	"github.com/fasthttp/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

type Client struct {
	Hub        *Hub
	Connection *websocket.Conn
	Send       chan []byte
}

// we need three functions here
func (c *Client) readPump() {

}

func (c *Client) writePump() {

}

func PeerChatConn() {

}

// "github.com/fasthttp/websocket"
