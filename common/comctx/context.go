package comctx

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

type key int

const (
	RequestTraceIDKey key = iota
	QueueMsgIDKey

	InjectedCtx
)
const (
	RequestTraceIDStringKey = "requestTraceID"
)

func GetRequestTraceID(ctx context.Context) string {
	traceID, ok := ctx.Value(RequestTraceIDKey).(string)
	if ok {
		return traceID
	}
	return ""
}

func SetRequestTraceID(ctx context.Context, requestTraceID string) (outCtx context.Context) {
	outCtx = context.WithValue(ctx, RequestTraceIDKey, requestTraceID)
	return outCtx
}

func GetInjectedCtx(fiberCtx *fiber.Ctx) context.Context {
	ctx, ok := fiberCtx.Locals(InjectedCtx).(context.Context)
	if !ok {
		ctx = InjectCtx(fiberCtx)
	}
	return ctx
}

func InjectCtx(fiberCtx *fiber.Ctx) context.Context {
	newCtx := context.Background()
	requestTraceID, ok1 := fiberCtx.Locals(RequestTraceIDStringKey).(string)
	if !ok1 {
		return newCtx
	}
	newCtx = SetRequestTraceID(newCtx, requestTraceID)

	_ = fiberCtx.Locals(InjectedCtx, newCtx)
	return newCtx
}
