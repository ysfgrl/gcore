package gcore

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ysfgrl/gcore/groute"
)

type IModule interface {
	Register(app *fiber.App)
	AddController(ctrl groute.IRoute)
	Init()
}
