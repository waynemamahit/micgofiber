package main

import (
	"micgofiber/router"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
)

func main() {
	app := fiber.New()

	// Middleware
	app.Use(helmet.New())
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(limiter.New(limiter.Config{
		Max:        20,
		Expiration: 30 * time.Second,
	}))
	headerCsrf := "X-CSRF-Token"
	app.Use(csrf.New(csrf.Config{
		KeyLookup:         "header:" + headerCsrf,
		CookieName:        "__Secure-csrf_",
		CookieSameSite:    "Lax",
		CookieSecure:      true,
		CookieSessionOnly: true,
		Expiration:        1 * time.Hour,
		KeyGenerator:      utils.UUIDv4,
		ErrorHandler:      fiber.DefaultErrorHandler,
		Extractor:         csrf.CsrfFromHeader(headerCsrf),
		Session:           session.New(),
		SessionKey:        "fiber.csrf.token",
		HandlerContextKey: "fiber.csrf.handler",
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST",
	}))

	apiV1Prefix := "/api/v1"
	router.NewTodoRouter(app.Group(apiV1Prefix + "/todo"))

	app.Listen(":3002")
}
