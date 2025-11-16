package gcore

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ysfgrl/gcore/gconf"
	"github.com/ysfgrl/gcore/groute"
)

type BaseModule struct {
	conf   *gconf.Conf
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
func (b *BaseModule) SetConf(conf *gconf.Conf) {
	b.conf = conf
}
func (b *BaseModule) GetConf() *gconf.Conf {
	return b.conf
}
