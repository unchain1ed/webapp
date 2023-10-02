package main

import (
	"log"

	"github.com/unchain1ed/webapp/controller"
)

func main() {
	log.Println("Start App...")
	controller.GetRouter()
}