package gcore

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ysfgrl/gcore/groute"
)
import "github.com/ysfgrl/gcore/gconf"

type IModule interface {
	Register(app *fiber.App)
	SetConf(conf *gconf.Conf)
	GetConf() *gconf.Conf
	AddController(ctrl groute.IRoute)
	Init()
}
