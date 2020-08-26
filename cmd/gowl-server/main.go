package main

import (

	//"github.com/accnameowl/gowl/cmd/gowl-server/router"
	"github.com/gofiber/fiber"
)

func main() {

	config := <-ReadEnvFromYaml()

	runtimeSettings := <-FetchFiberSettings(&config)
	app := fiber.New(&runtimeSettings)

	app.Listen(config.Server.Port)
}
