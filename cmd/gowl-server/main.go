package main

import (
	//"github.com/accnameowl/gowl/cmd/gowl-server/router"
	"flag"

	"github.com/gofiber/fiber"
	"github.com/gofiber/pprof"
)

func main() {

	// ! Get configs from config.yml
	config := ReadEnvFromYaml()

	// ! runtime flags
	flag.BoolVar(&config.PPROF.Active, "pprof", false, "Run with PPROF tool")
	flag.Parse()

	runtimeSettings := FetchFiberSettings(&config)
	app := fiber.New(runtimeSettings)

	//! pprof tool
	if config.PPROF.Active {
		app.Use(pprof.New())
	}

	app.Static("/profiling", "./router/public/index.html")

	app.Listen(config.Server.Port)
}
