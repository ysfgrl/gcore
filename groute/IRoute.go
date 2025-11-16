package groute

import (
	"github.com/gofiber/fiber/v2"
)

type IRoute interface {
	Register(app *fiber.App)
}
