package main

import (
	"github.com/unchain1ed/server-app/controller"
	"log"
)

func main() {
	router := controller.GetRouter()
	log.Println("Start App...")
	router.Run(":8080")
}