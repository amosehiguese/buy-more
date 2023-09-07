package main

import (
	"log"

	"github.com/amosehiguese/buy-more/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatalln(err)
	}
}
