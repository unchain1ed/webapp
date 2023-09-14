package main

import (
	"log"

	"github.com/unchain1ed/server-app/controller"
)

func main() {
	log.Println("Start App...")
	controller.GetRouter()
}