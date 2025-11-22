package gcore

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ysfgrl/gcore/gerror"
	"github.com/ysfgrl/gcore/gmodel"
)

type Server struct {
	app         *fiber.App
	isListening bool
	modules     []IModule
	middlewares []fiber.Handler
	host        string
	port        int
}

func NewServer(host string, port int) *Server {
	return &Server{
		modules:     make([]IModule, 0),
		middlewares: make([]fiber.Handler, 0),
		host:        host,
		port:        port,
	}
}

func (server *Server) ListenAndServe() error {
	if server.isListening {
		return nil
	}
	if server.app == nil {
		server.app = fiber.New(fiber.Config{
			ErrorHandler:  server.errorHandler,
			StrictRouting: true,
		})
	}
	for _, middleware := range server.middlewares {
		server.app.Use(middleware)
	}
	for _, mod := range server.modules {
		mod.Register(server.app)
		mod.Init()
	}
	host := fmt.Sprintf("%s:%d", server.host, server.port)
	fmt.Println(host)
	return server.app.Listen(host)
}

func (server *Server) AddModule(module IModule) {
	if module == nil {
		return
	}
	server.modules = append(server.modules, module)
}

func (server *Server) IsListening() bool {
	return server.isListening
}
func (server *Server) Use(middle fiber.Handler) {
	server.middlewares = append(server.middlewares, middle)
}

func (server *Server) errorHandler(ctx *fiber.Ctx, err error) error {
	var e *fiber.Error
	if errors.As(err, &e) {
		return ctx.Status(e.Code).JSON(gmodel.Response{
			Code: e.Code,
			Error: &gerror.Error{
				Code:   "code.error" + strconv.Itoa(e.Code),
				Level:  gerror.LevelError,
				Detail: e.Message,
			},
		})
	}
	return ctx.Status(fiber.StatusInternalServerError).JSON(gmodel.Response{
		Code: fiber.StatusInternalServerError,
		Error: &gerror.Error{
			Code:   "internalServerError",
			Level:  gerror.LevelFatal,
			Detail: err.Error(),
			Err:    err,
		},
	})
}
