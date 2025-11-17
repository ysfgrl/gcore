package gcore

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ysfgrl/gcore/groute"
)

type BaseModule struct {
	Routes []groute.IRoute
}

func (b *BaseModule) Register(app *fiber.App) {
	for _, controller := range b.Routes {
		controller.Register(app)
	}
}
func (b *BaseModule) AddController(ctrl groute.IRoute) {
	b.Routes = append(b.Routes, ctrl)
}
