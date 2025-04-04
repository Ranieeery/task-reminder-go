package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	log.Fatal(app.Listen(":4000"))
}
