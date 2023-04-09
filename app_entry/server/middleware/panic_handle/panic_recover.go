package panichandle

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gofiber/fiber/v2"
	"github.com/yunerou/oauth2-client/common/comctx"
	"github.com/yunerou/oauth2-client/singleton"
)

// PanicRecover ..
func PanicRecover() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) (err error) {
		// Catch panics
		defer func() {
			if r := recover(); r != nil {
				internalServerErrorResponseStackTraceHandler(c, r)
			}
		}()

		// Return err if exist, else move to next handler
		return c.Next()
	}
}

func internalServerErrorResponseStackTraceHandler(c *fiber.Ctx, e interface{}) {
	ctx := comctx.GetInjectedCtx(c)
	singleton.GetLogger(ctx).Info().
		Str("panic", fmt.Sprintf("%v\n%s", e, string(debug.Stack()))).Send()

	_ = c.Status(http.StatusInternalServerError).JSON(fiber.Map{
		"message": "golang panic",
		"panic":   fmt.Sprintf("%v", e),
	})
}
