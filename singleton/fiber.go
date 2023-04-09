package singleton

import (
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

var (
	onceFiber      sync.Once
	fiberSingleton *fiberWrap
)

type fiberWrap struct {
	*fiber.App
}

func GetFiber() *fiberWrap {
	onceFiber.Do(func() {
		engine := html.New(GetViper().GetString("VIEWS_DIRECTORY"), ".html")
		app := fiber.New(fiber.Config{
			Views: engine,
		})
		fiberSingleton = &fiberWrap{
			app,
		}
	})
	return fiberSingleton

}
