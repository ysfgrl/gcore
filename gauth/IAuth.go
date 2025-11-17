package gauth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ysfgrl/gcore/gerror"
)

type IAuth interface {
	GetToken(ctx *fiber.Ctx) (string, *gerror.Error)
	GetUser(ctx *fiber.Ctx) (*Claims, *gerror.Error)
	RoleRequire(ctx *fiber.Ctx, roles []string) error
	Require(ctx *fiber.Ctx) error
	CreateToken(payload Claims) (string, *gerror.Error)
	Init()
}
