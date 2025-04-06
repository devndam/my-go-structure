package main

import (
	"log"

	"github.com/devndam/truepeer/app/config"
	"github.com/gofiber/fiber/v2"
)

func init() {
	config.LoadEnVariables()
}

func main() {
	app := fiber.New()

	log.Fatal(app.Listen(":3000"))
}
