package commonpages

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/yunerou/oauth2-client/common/comctx"
)

type CallbackFromOauth2Provider struct {
	State *string `query:"state"`
	Code  *string `query:"code"`
	Err   *string `query:"error"`
}

func (d *commonPages) Callback(sStore *session.Store) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		ctx := comctx.GetInjectedCtx(c)

		// parse query param from FE
		input := new(CallbackFromOauth2Provider)
		if err := c.QueryParser(input); err != nil {
			return err
		}

		if input.Err != nil {
			return c.JSON(input)
		}

		oauth2PrdName := c.Params("oauth2Provider", "")
		if oauth2PrdName == "" {
			panic("call wrong callback url format")
		}
		oauth2Cfg, ok := d.oauth2Provider[oauth2PrdName]
		if !ok {
			panic(fmt.Sprintf("oauth2 provider is not support [%s]", oauth2PrdName))
		}

		// validate state from redis. Bypass

		// get code

		fmt.Printf("code = %s \n", *input.Code)

		if input.Code != nil {
			tokenRes, err := oauth2Cfg.Exchange(ctx, *input.Code)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "Unable to exchange code for token: %s\n", err)
			}
			// DEBUG code
			callGithubInfo(tokenRes.AccessToken)
			// END DEBUG
			return c.JSON(tokenRes)
		}
		return c.JSON("something go wrong")
	}
}

func callGithubInfo(accessTK string) {
	url := "https://api.github.com/user/emails"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessTK))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
