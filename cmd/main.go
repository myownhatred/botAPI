package main

import (
	"log"

	"github.com/myownhatred/botAPI"
)

func main() {
	srv := new(botAPI.Server)
	if err := srv.Run("6666"); err != nil {
		log.Fatalf("Server run error: %s", err.Error())
	}
}
