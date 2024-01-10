package lib

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
)

type AppConfig struct {
	App        *fiber.App
	ApiV1      string
	Port       string
	CsrfHeader string
	CsrfCookie string
}

func NewApp() *AppConfig {
	app := fiber.New()

	// Global Config
	csrfHeader := "X-Csrf-Token"
	csrfCookie := "__Secure-csrf_"

	// Middleware
	app.Use(helmet.New())
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(limiter.New(limiter.Config{
		Max:        20,
		Expiration: 30 * time.Second,
	}))
	app.Use(csrf.New(csrf.Config{
		KeyLookup:         "header:" + csrfHeader,
		CookieName:        csrfCookie,
		CookieSameSite:    "Lax",
		CookieSecure:      true,
		CookieSessionOnly: true,
		Expiration:        1 * time.Hour,
		KeyGenerator:      utils.UUIDv4,
		ErrorHandler:      fiber.DefaultErrorHandler,
		Extractor:         csrf.CsrfFromHeader(csrfHeader),
		Session:           session.New(),
		SessionKey:        "fiber.csrf.token",
		HandlerContextKey: "fiber.csrf.handler",
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST",
	}))
	app.Use("/static", filesystem.New(filesystem.Config{
		Root: http.Dir("./storage"),
	}))

	return &AppConfig{
		App:        app,
		ApiV1:      "/api/v1",
		Port:       "3002",
		CsrfHeader: csrfHeader,
		CsrfCookie: csrfCookie,
	}
}

func (app *AppConfig) GetRouterV1(url string) fiber.Router {
	return app.App.Group(app.ApiV1 + url)
}
