package main

import (
	"github.com/unchain1ed/server-app/controller"
	"log"
)

func main() {
	log.Println("Start App...")
	controller.GetRouter()
	// router := controller.GetRouter()
	// router.Run(":8080")
}