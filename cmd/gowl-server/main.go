package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"log"
)

func main() {

	await := <-ReadEnvFromYaml()

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Send to client")
	})

	// Concurrent listen function
	go func() {
		log.Fatal(app.Listen(3000))
	}()
}
