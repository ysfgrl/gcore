package groute

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/ysfgrl/gcore/gerror"
	"github.com/ysfgrl/gcore/gmodel"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseRoute struct{}

func (ctrl *BaseRoute) GetIdParams(c *fiber.Ctx, key string) (primitive.ObjectID, *gerror.Error) {
	value := c.Params(key, key)
	if strings.EqualFold(value, key) {
		return primitive.ObjectID{}, gerror.PathParamRequired
	}
	id, err := primitive.ObjectIDFromHex(value)
	if err != nil {
		return primitive.ObjectID{}, gerror.PathParamRequired
	}
	return id, nil
}
func (ctrl *BaseRoute) GetParams(c *fiber.Ctx, key string) (string, *gerror.Error) {
	value := c.Params(key, key)
	if strings.EqualFold(value, key) {
		return key, gerror.PathParamRequired
	}
	return value, nil
}

func (ctrl *BaseRoute) Created(c *fiber.Ctx, content any) error {
	return c.Status(fiber.StatusCreated).JSON(gmodel.Response{
		Code:    fiber.StatusCreated,
		Content: content,
		Error:   nil,
	})
}
func (ctrl *BaseRoute) Ok(c *fiber.Ctx, content any) error {
	return c.Status(fiber.StatusOK).JSON(gmodel.Response{
		Code:    fiber.StatusOK,
		Content: content,
		Error:   nil,
	})
}
func (ctrl *BaseRoute) NotImplemented(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).JSON(gmodel.Response{
		Code:    fiber.StatusNotImplemented,
		Content: nil,
		Error:   gerror.NotImplementedError,
	})
}

func (ctrl *BaseRoute) InternalServerError(c *fiber.Ctx, err *gerror.Error) error {
	return c.Status(fiber.StatusInternalServerError).JSON(gmodel.Response{
		Code:    fiber.StatusInternalServerError,
		Content: nil,
		Error:   err,
	})
}

func (ctrl *BaseRoute) NotFound(c *fiber.Ctx, err *gerror.Error) error {
	return c.Status(fiber.StatusNotFound).JSON(gmodel.Response{
		Code:  fiber.StatusNotFound,
		Error: err,
	})
}

func (ctrl *BaseRoute) Unauthorized(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(gmodel.Response{
		Code:  fiber.StatusUnauthorized,
		Error: gerror.PermissionDenied,
	})
}

func (ctrl *BaseRoute) Forbidden(c *fiber.Ctx) error {
	return c.Status(fiber.StatusForbidden).JSON(gmodel.Response{
		Code:  fiber.StatusForbidden,
		Error: gerror.ForbiddenError,
	})
}

func (ctrl *BaseRoute) BadRequest(c *fiber.Ctx, err *gerror.Error) error {
	return c.Status(fiber.StatusBadRequest).JSON(gmodel.Response{
		Code:  fiber.StatusBadRequest,
		Error: err,
	})
}
