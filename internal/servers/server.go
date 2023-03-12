package servers

import (
	"flag"
	"os"
	"time"

	"github.com/amanraghuvanshi/videostreaming/internal/handlers"
	"gorm.io/gorm/logger"

	"github.com/gofiber/fiber/template/html"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket"
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
	app.Get("/room/:uuid/websocket", websocket.New(handlers.Roomwebsocket, websocket.Config{HandshakeTimeout: 10 * time.Second}))
	app.Get("/room/:uuid/chat", handler.RoomChat)
	app.Get("/room/:uuid/websocket", websocket.New(handlers.RoomChatWebsocket))
	app.Get("/room/:uuid/viewer/websocket", websocket.New(handlers.NewViewerWebSocket))
	app.Get("/stream/:ssuid", handlers.Stream)
	app.Get("/stream/:ssuid/websocket")
	app.Get("/stream/:ssuid/chat/websocket")
	app.Get("/stream/:suuid/viewer/websocket")
}
