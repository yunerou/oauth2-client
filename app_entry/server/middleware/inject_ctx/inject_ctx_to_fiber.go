package panichandle

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yunerou/oauth2-client/common/comctx"
)

// InjectCtx ..
func InjectCtx() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) (err error) {

		_ = comctx.InjectCtx(c)

		// Return err if exist, else move to next handler
		return c.Next()
	}
}
