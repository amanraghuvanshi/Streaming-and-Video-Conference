package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	gguid "github.com/google/uuid"
)

// this will be the function that will create or get the room.
func createOrGetRoom(uuid string) (string, string, room) {

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

}
