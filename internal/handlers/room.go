package handlers

import (
	"fmt"

	w "github.com/amanraghuvanshi/videostreaming/pkg/webrtc"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	gguid "github.com/google/uuid"
)

type websocketMessage struct {
	Event string `json:"event"`
	Data  string `json:"data"`
}

// this will be the function that will create or get the room.
func createOrGetRoom(uuid string) (string, string, *w.Room) {

}

func RomViewerWebsocket(c *websocket.Conn) {

}

func roomViewerConn(c *websocket.Conn, p *w.Peers) {

}

// this will get the room ID and connection string
func RoomCreate(c *fiber.Ctx) error {
	return c.Redirect(fmt.Sprintf("/room/%s", gguid.New().String()))
}

// this function will be getting the room if it exists, otherwise it will create the room
func Room(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		c.Status(400)
		return nil
	}

	uuid, suuid, _ := createOrGetRoom(uuid)
}

// this will create the room with uuid
func Roomwebsocket(c *websocket.Conn) {
	uuid := c.Params("uuid")
	if uuid == "" {
		return
	}
	_, _, room := createOrGetRoom(uuid)
	w.RoomConn(c, room.Peers)
}
