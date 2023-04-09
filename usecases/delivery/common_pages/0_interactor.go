package commonpages

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/oauth2"
)

type CommonPages interface {
	Home(sStore *session.Store) func(c *fiber.Ctx) error
	Callback(sStore *session.Store) func(*fiber.Ctx) error
	// ErrorPage(c *fiber.Ctx) error
}

type commonPages struct {
	// mapping name with Oauth2 provider config
	oauth2Provider map[string]*oauth2.Config
}

func NewCommonPages() CommonPages {
	inst := &commonPages{}
	inst.registerOauth2Provider()
	return inst
}
