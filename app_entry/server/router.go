package server

import (
	"net/http"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/yunerou/oauth2-client/app_entry/server/routename"
	"github.com/yunerou/oauth2-client/singleton"
)

var (
	once      sync.Once
	allRoutes *routesCollection
)

func routers() *routesCollection {
	once.Do(func() {
		allRoutes = initRoutesCollection()
	})
	return allRoutes
}

func healthcheckRegister(
	parentPath fiber.Router,
) {
	// health check
	parentPath.Get("/healthcheckz", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": "ok",
		})
	}).Name(routename.Healthcheck)
	// version check
	parentPath.Get("/versionz", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"version": singleton.GetViper().GetString("VERSION"),
		})
	}).Name(routename.Version)
	// panic test
	parentPath.Get("/panicz", func(c *fiber.Ctx) error {
		panic("test-test-test-something wrong")
	})
}

func homepageRegister(
	parentPath fiber.Router,
	sessionStore *session.Store,
) {
	parentPath.Get(
		"/",
		routers().commonpages.Home(sessionStore),
	).Name(routename.Homepage)

	parentPath.Get(
		":oauth2Provider/callback",
		routers().commonpages.Callback(sessionStore),
	).Name(routename.Callback)
}
