package main

import (
	//"github.com/accnameowl/gowl/cmd/gowl-server/router"
	"flag"

	"github.com/gofiber/fiber"
	"github.com/pkg/profile"
)

func main() {

	// ! Get configs from config.yml
	config := <-ReadEnvFromYaml()

	// ! runtime flags
	flag.BoolVar(&config.PPROF.Active, "pprof", false, "Run with PPROF tool")
	flag.BoolVar(&config.PPROF.CPU, "cpu", false, "Pprof<-cpu")
	flag.BoolVar(&config.PPROF.Mem, "mem", false, "Pprof<-mem")
	flag.BoolVar(&config.PPROF.Trace, "trace", false, "Pprof<-Trace")
	flag.IntVar(&config.PPROF.MemRate, "rate", 0, "Pprof<-mem<-MemProfilerate()")
	flag.Parse()

	// !PROFILING TOOL
	// * only running if flag pprof is enabled enabled
	// * Run `go tool pprof -http=:PORT FILE` to review profiling
	if config.PPROF.Active {
		if config.PPROF.CPU {
			defer profile.Start(profile.CPUProfile, profile.ProfilePath("./dumps")).Stop()
		}
		if config.PPROF.Mem {
			if config.PPROF.MemRate > 0 {
				defer profile.Start(profile.MemProfile, profile.MemProfileRate(config.PPROF.MemRate), profile.ProfilePath("./dumps")).Stop()
			} else {
				defer profile.Start(profile.MemProfile, profile.ProfilePath("./dumps")).Stop()
			}
		}
	}

	runtimeSettings := <-FetchFiberSettings(&config)
	app := fiber.New(&runtimeSettings)

	app.Static("/profiling", "./router/public/index.html")

	app.Listen(config.Server.Port)
}
