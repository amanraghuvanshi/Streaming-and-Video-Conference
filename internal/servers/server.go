package servers

import (
	"flag"
	"os"
	"time"

	"github.com/amanraghuvanshi/videostreaming/internal/handlers"
	w "github.com/amanraghuvanshi/videostreaming/pkg/webrtc"

	"github.com/gofiber/fiber/template/html"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/websocket/v2"
)

var (
	addr = flag.String("addr", ":", os.Getenv("PORT"))
	cert = flag.String("cert", "", "")
	key  = flag.String("key", "", "")
)

func Run() error {
	flag.Parse()

	if *addr == ":" {
		*addr = ":8080"
	}

	//  <-------------------------------------------------->
	// here we are creating the basic HTML files to create the view of the application
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine})
	app.Use(logger.New())
	app.Use(cors.New())

	//  <-------------------------------------------------->
	// ROUTES
	app.Get("/", handlers.Welcome)
	app.Get("/room/create", handlers.RoomCreate)
	app.Get("/room/:uuid", handlers.Room)
	app.Get("/room/:uuid/websocket", websocket.New(handlers.Roomwebsocket, websocket.Config{
		HandshakeTimeout: 10 * time.Second}))
	app.Get("/room/:uuid/chat", handler.RoomChat)
	app.Get("/room/:uuid/websocket", websocket.New(handlers.RoomChatWebsocket))
	app.Get("/room/:uuid/viewer/websocket", websocket.New(handlers.NewViewerWebSocket))
	app.Get("/stream/:ssuid", handlers.Stream)
	app.Get("/stream/:ssuid/websocket", websocket.New(handlers.StreamWebsocket, websocket.Config{
		HandshakeTimeout: 10 * time.Second,
	}))
	app.Get("/stream/:ssuid/chat/websocket", websocket.New(handlers.StreamChatWebsocket))
	app.Get("/stream/:suuid/viewer/websocket", websocket.New(handlers.StreamViewerWebsocket))
	app.Static("/", "./assests")

	w.Rooms = make(map[string]*w.room)
	w.Streams = make(map[string]*w.room)

	if *cert != "" {
		return app.ListenTLS(*addr, *cert, *key)
	}
	return app.Listen(*addr)

	// <------------------------------------------>
	go DispatchKeyFrames()

}

func DispatchKeyFrames() {
	for range time.NewTicker(time.Second * 3).C {
		for _, room := range w.Rooms {
			room.Peers.DispatchKeyFrames()
		}
	}
}
