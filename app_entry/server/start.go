package server

import (
	"encoding/gob"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
	"github.com/yunerou/oauth2-client/app_entry/server/middleware"
	panichandle "github.com/yunerou/oauth2-client/app_entry/server/middleware/panic_handle"
	"github.com/yunerou/oauth2-client/common/comctx"
	"github.com/yunerou/oauth2-client/singleton"
)

func StartServer() {
	gob.Register(http.Cookie{})

	app := singleton.GetFiber()
	redisStore := redis.New(redis.Config{
		Host:     singleton.GetViper().GetString("REDIS.WRITER"),
		Port:     singleton.GetViper().GetInt("REDIS.PORT"),
		Password: singleton.GetViper().GetString("REDIS.PASSWORD"),
		Database: singleton.GetViper().GetInt("REDIS.DATABASE"),
	})
	// session store
	loginSessStore := session.New(session.Config{
		Storage:        redisStore,
		Expiration:     singleton.GetViper().GetDuration("SESSION.TTL"),
		CookieHTTPOnly: true,
		KeyLookup: "cookie:" + singleton.GetViper().
			GetString("SESSION.COOKIENAME_PREFIX") +
			"auth",
	})
	// crsfStore := session.New(session.Config{
	//	Storage:        redisStore,
	//	Expiration:     singleton.GetViper().GetDuration("SESSION.TTL"),
	//	CookieHTTPOnly: true,
	//	KeyLookup: "cookie:" + singleton.GetViper().
	//		GetString("SESSION.COOKIENAME_PREFIX") +
	//		"_csrf_token",
	// })

	for _, mdw := range globalAppMiddlerwares() {
		app.Use(mdw)
	}
	// router
	healthcheckRegister(app)
	homepageRegister(app, loginSessStore)
	// Start
	log.Fatal(app.Listen(":8080"))
}

func globalAppMiddlerwares() []func(*fiber.Ctx) error {
	return []func(*fiber.Ctx) error{
		// panic recover
		panichandle.PanicRecover(),
		// request trace id
		requestid.New(
			requestid.Config{
				Header:     singleton.GetViper().GetString("TRACE_ID_HEADER"),
				ContextKey: comctx.RequestTraceIDStringKey,
			}),
		// log
		middleware.Log,
	}
}
