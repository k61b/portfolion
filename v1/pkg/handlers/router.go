package handlers

import (
	"time"

	"github.com/kayraberktuncer/portfolion/pkg/common/lib"
	"github.com/kayraberktuncer/portfolion/pkg/common/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	_ "github.com/kayraberktuncer/portfolion/docs"
	fiberSwagger "github.com/swaggo/fiber-swagger"
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
		Max:        20,
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

	app.Use(cors.New(cors.Config{
		AllowOrigins:     lib.GoDotEnvVariable("ALLOWED_ORIGINS"),
		AllowCredentials: true,
	}))

	api := app.Group("/api/v1")

	app.Get("/swagger/*", fiberSwagger.WrapHandler)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/swagger/index.html")
	})

	api.Post("/session", h.Session)
	api.Get("/auth", h.AuthMiddleware, h.Auth)
	api.Get("/logout", h.AuthMiddleware, h.Logout)
	api.Get("/bookmarks", h.AuthMiddleware, h.GetBookmarks)
	api.Post("/bookmarks", h.AuthMiddleware, h.CreateBookmark)
	api.Put("/bookmarks/:symbol", h.AuthMiddleware, h.UpdateBookmark)
	api.Delete("/bookmarks/:symbol", h.AuthMiddleware, h.DeleteBookmark)
	api.Get("/search/:symbol", cachingMiddleware, h.SearchSymbol)

	go h.UpdateSymbolValuesPeriodically(1 * time.Minute)

	app.Listen(h.listenAddr)
}

func (h *Handlers) UpdateSymbolValuesPeriodically(interval time.Duration) {
	for {
		h.UpdateSymbolValues()
		time.Sleep(interval)
	}
}
