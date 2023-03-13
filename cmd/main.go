package main

import (
	"log"

	"github.com/amanraghuvanshi/videostreaming/internal/servers"
)

func main() {
	if err := servers.Run(); err != nil {
		log.Fatalln(err.Error())
	}
}
