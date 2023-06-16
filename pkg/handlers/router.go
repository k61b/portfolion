package handlers

import (
	"github.com/kayraberktuncer/portfolion/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

type Handlers struct {
	listenAddr string
	store      models.Store
}

func NewHandlers(listenAddr string, store models.Store) *Handlers {
	return &Handlers{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (h *Handlers) Run() {
	app := fiber.New()

	app.Post("/session", h.Session)
	app.Get("/auth", h.AuthMiddleware, h.Auth)
	app.Get("/logout", h.AuthMiddleware, h.Logout)

	app.Listen(h.listenAddr)
}
