package main

import (
	//"github.com/accnameowl/gowl/cmd/gowl-server/router"
	"flag"

	"github.com/gofiber/fiber"
	"github.com/pkg/profile"
)

func main() {

	config := <-ReadEnvFromYaml()

	flag.BoolVar(&config.Build.Pprof, "pprof", false, "Run gotool pprof")
	flag.Parse()

	// !PROFILING TOOL
	// * only running if flag pprof is enabled enabled
	if config.Build.Pprof {
		defer profile.Start(profile.CPUProfile, profile.ProfilePath("./")).Stop()
	}

	runtimeSettings := <-FetchFiberSettings(&config)
	app := fiber.New(&runtimeSettings)

	app.Static("/profiling", "./router/public/profiletest.html")

	app.Listen(config.Server.Port)
}
