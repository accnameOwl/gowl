package router

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber"
)

// NewStatic ...
// Attach new static to a fiber.app
//
// params: standard,
func NewStatic(app *fiber.App, prefix string, dir string, params string) {
	if prefix == "" || dir == "" {
		fmt.Println("Could not find prefix or dir argument")
	} else {
		if len(params) > 0 {
			strList := strings.Split(params, ", ")
			for _, v := range strList {
				if v == "standard" {
					InitStandardStatics(app)
				}
				// insert other params below
			}
		} else {
			app.Static(prefix, dir)
		}
	}
}

// InitStandardStatics ...
func InitStandardStatics(app *fiber.App) {
	app.Static("/", "./public/index.html")
}
