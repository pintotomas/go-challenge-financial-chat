package main

import (
	"fmt"
	"go-challenge-financial-chat/app/config"
	"go-challenge-financial-chat/app/db"
	"go-challenge-financial-chat/app/router"
	"log"
)

func main() {
	config.Load()
	db.Connect()

	r := router.Setup()
	addr := fmt.Sprintf(":%s", config.ENV.Port)

	if err := r.Run(addr); err != nil {
		log.Fatalln("Failed to start the application.", err)
	}
}
