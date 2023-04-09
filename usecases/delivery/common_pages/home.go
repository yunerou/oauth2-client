package commonpages

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/oauth2"
)

func (d *commonPages) Home(sStore *session.Store) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// ctx := comctx.GetInjectedCtx(c)
		// sess, err := sStore.Get(c)
		// sess.Set("name", "xyz")
		// // Save session
		// if err = sess.Save(); err != nil {
		// 	panic(err)
		// }
		// if err != nil {
		// 	singleton.GetLogger(ctx).Error().Err(err).Msg("config session fail")
		// }

		state := fmt.Sprintf("%s%d", "state", time.Now().Unix())

		nonce := fmt.Sprintf("%s%d", "nonce", time.Now().Unix())

		ghAuthCodeURL := d.oauth2Provider["github"].AuthCodeURL(
			state,
			// oauth2.SetAuthURLParam("audience", strings.Join(audience, "+")),
			oauth2.SetAuthURLParam("nonce", nonce),
			// oauth2.SetAuthURLParam("prompt", strings.Join(prompt, "+")),
			oauth2.SetAuthURLParam("max_age", "300"),
		)

		ggAuthCodeURL := d.oauth2Provider["google"].AuthCodeURL(
			state,
			// oauth2.SetAuthURLParam("audience", strings.Join(audience, "+")),
			oauth2.SetAuthURLParam("nonce", nonce),
			// oauth2.SetAuthURLParam("prompt", strings.Join(prompt, "+")),
			oauth2.SetAuthURLParam("max_age", "300"),
		)

		dataMap := map[string]interface{}{
			"Title":             "Homepage",
			"GitubAuthCodeURL":  ghAuthCodeURL,
			"GoogleAuthCodeURL": ggAuthCodeURL,
		}
		return c.Render("home", dataMap)
	}
}
