package handlers

import (
	"time"

	"github.com/kayraberktuncer/portfolion/pkg/common/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/limiter"
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

	limit := limiter.New(limiter.Config{
		Max:        10,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"message": "Too many requests",
			})
		},
	})

	cachingMiddleware := cache.New(cache.Config{
		Expiration:   1 * time.Minute,
		CacheControl: true,
	})

	app.Use(limit)

	app.Post("/session", h.Session)
	app.Get("/auth", h.AuthMiddleware, h.Auth)
	app.Get("/logout", h.AuthMiddleware, h.Logout)
	app.Get("/bookmarks", cachingMiddleware, h.AuthMiddleware, h.GetBookmarks)
	app.Post("/bookmarks", h.AuthMiddleware, h.CreateBookmark)
	app.Put("/bookmarks/:symbol", h.AuthMiddleware, h.UpdateBookmark)
	app.Delete("/bookmarks/:symbol", h.AuthMiddleware, h.DeleteBookmark)
	app.Get("/search/:symbol", cachingMiddleware, h.AuthMiddleware, h.SearchSymbol)

	app.Listen(h.listenAddr)
}
