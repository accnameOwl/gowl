package router

import (
	"github.com/gofiber/fiber"
)

// InitStandardStatics ...
func InitStandardStatics(app *fiber.App) {
	app.Static("/", "./public/index.html")
}
