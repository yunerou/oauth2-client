package middleware

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yunerou/oauth2-client/common/comctx"
	"github.com/yunerou/oauth2-client/singleton"
)

// LogStruct - logger structure.
type LogStruct struct {
	IP        string `json:"ip"`
	Method    string `json:"method"`
	URL       string `json:"url"`
	StartTime string `json:"start_time"`
	Duration  string `json:"duration"`
	Agent     string `json:"agent"`
	Status    int    `json:"status"`

	BodySize  int
	AllParams map[string]string
}

// Log - logger will print JSON formatted logs onto STDOUT.
func Log(ctx *fiber.Ctx) error {
	t := time.Now()
	logger := LogStruct{
		IP:        ctx.IP(),
		URL:       ctx.OriginalURL(),
		StartTime: t.Format(time.RFC3339),
		Method:    string(ctx.Context().Method()),
		Agent:     string(ctx.Context().UserAgent()),
		AllParams: ctx.AllParams(),
	}
	_ = ctx.Next()
	logger.Status = ctx.Context().Response.StatusCode()
	logger.Duration = time.Since(t).String()
	logStr, _ := json.Marshal(logger)

	lg := singleton.GetLogger(comctx.GetInjectedCtx(ctx)).Info()

	lg.RawJSON("http", logStr).Send()
	return nil
}
