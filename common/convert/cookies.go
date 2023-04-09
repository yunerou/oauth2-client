package convert

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func FiberCookie2HttpCookie(in *fiber.Cookie) (out *http.Cookie) {
	if in == nil {
		return nil
	}

	out = &http.Cookie{
		Name:     in.Name,
		Value:    in.Value,
		MaxAge:   in.MaxAge,
		Secure:   in.Secure,
		HttpOnly: in.HTTPOnly,
	}

	if in.Path != "" {
		out.Path = in.Path
	}
	if in.Domain != "" {
		out.Domain = in.Domain
	}
	if (in.Expires != time.Time{}) {
		out.Expires = in.Expires
	}

	switch utils.ToLower(in.SameSite) {
	case fiber.CookieSameSiteStrictMode:
		out.SameSite = http.SameSiteStrictMode
	case fiber.CookieSameSiteNoneMode:
		out.SameSite = http.SameSiteNoneMode
	case fiber.CookieSameSiteLaxMode:
		out.SameSite = http.SameSiteLaxMode
	case fiber.CookieSameSiteDisabled:
		out.SameSite = http.SameSiteDefaultMode
	default:
		out.SameSite = http.SameSiteLaxMode
	}
	return
}

func HttpCookie2FiberCookie(in *http.Cookie) (out *fiber.Cookie) {
	if in == nil {
		return nil
	}

	out = &fiber.Cookie{
		Name:     in.Name,
		Value:    in.Value,
		MaxAge:   in.MaxAge,
		Secure:   in.Secure,
		HTTPOnly: in.HttpOnly,
	}
	if in.Path != "" {
		out.Path = in.Path
	}
	if in.Domain != "" {
		out.Domain = in.Domain
	}
	if (in.Expires != time.Time{}) {
		out.Expires = in.Expires
	}

	switch in.SameSite {
	case http.SameSiteStrictMode:
		out.SameSite = fiber.CookieSameSiteStrictMode
	case http.SameSiteNoneMode:
		out.SameSite = fiber.CookieSameSiteNoneMode
	case http.SameSiteLaxMode:
		out.SameSite = fiber.CookieSameSiteLaxMode
	case http.SameSiteDefaultMode:
		out.SameSite = fiber.CookieSameSiteDisabled
	default:
		out.SameSite = fiber.CookieSameSiteLaxMode
	}
	return
}
