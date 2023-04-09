package singleton

import (
	"context"
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/yunerou/oauth2-client/common/comctx"

	"github.com/yunerou/oauth2-client/providers/chat"
)

func GetLogger(ctx context.Context) *zerolog.Logger {
	lc := getZeroLog()
	if ctx != nil {
		requestTraceID := comctx.GetRequestTraceID(ctx)
		if requestTraceID != "" {
			*lc.Context = lc.Context.Str(comctx.RequestTraceIDStringKey, requestTraceID)
		}
	}
	lo := lc.Logger()
	return &lo
}

var (
	zerologSingleton *zerologWrap
)

type zerologWrap struct {
	*zerolog.Context
}

// NewZerolog ... is singleton.
func getZeroLog() *zerologWrap {
	if zerologSingleton == nil {
		zerologSingleton = NewZerolog(nil)
	}
	return zerologSingleton
}

type ZerologConfig struct {
	Writer io.Writer
	HookFn func(e *zerolog.Event, level zerolog.Level, message string)
}

func NewZerolog(config *ZerologConfig) *zerologWrap {
	if config == nil {
		// default config
		config = &ZerologConfig{
			Writer: os.Stderr,
		}
	}
	zerolog.SetGlobalLevel(
		zerolog.Level(
			GetViper().GetInt("LOG.ZEROLOG_LEVEL")))
	zerologL := zerolog.New(config.Writer)
	if GetViper().GetBool("LOG.BEAUTIFUL") {
		zerologL = zerologL.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	zerologC := zerologL.With().Timestamp().Caller()
	zerologIns := &zerologWrap{
		&zerologC,
	}
	// hooks.
	// if config != nil && config.HookFn != nil {
	// 	zerologIns.Logger().HookFn(config.HookFn)
	// }

	return zerologIns
}

type SeverityHook struct {
	chatBot chat.ChatProvider
}

func (h SeverityHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	switch level {
	case zerolog.ErrorLevel,
		zerolog.WarnLevel:
		_ = h.chatBot.SendText(msg)
	case zerolog.DebugLevel,
		zerolog.InfoLevel,
		zerolog.FatalLevel,
		zerolog.PanicLevel,
		zerolog.NoLevel,
		zerolog.Disabled,
		zerolog.TraceLevel:
	}
}
