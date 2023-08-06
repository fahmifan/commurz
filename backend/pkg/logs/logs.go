package logs

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/xid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const RequestIDHeaderKey string = "X-Request-ID"

type ContextKey int

const reqIdCtxKey ContextKey = iota + 1

func GetRequestID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	val, _ := ctx.Value(reqIdCtxKey).(string)
	return val
}

func EchoRequestID() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ec echo.Context) error {
			rid := ec.Request().Header.Get(RequestIDHeaderKey)
			if rid == "" {
				rid = xid.New().String()
			}

			ctx := context.WithValue(ec.Request().Context(), reqIdCtxKey, rid)
			ec.SetRequest(ec.Request().Clone(ctx))
			ec.Request().Header.Set(RequestIDHeaderKey, rid)
			ec.Response().Header().Set(RequestIDHeaderKey, rid)

			return next(ec)
		}
	}
}

func EchoRequestLogger(debug bool) func(c echo.Context, v middleware.RequestLoggerValues) error {
	return func(c echo.Context, val middleware.RequestLoggerValues) error {
		reqID := c.Request().Header.Get(RequestIDHeaderKey)

		logger := zerolog.
			New(os.Stdout).
			With().
			Str(string(RequestIDHeaderKey), reqID).
			Str("method", c.Request().Method).
			Str("path", c.Request().URL.Path).
			Dur("latency", val.Latency).
			Int("status", c.Response().Status).
			Str("ip", c.RealIP()).
			Str("user-agent", c.Request().UserAgent()).
			AnErr("error", val.Error).
			Logger()

		if debug {
			logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stdout})
		}

		statusCode := c.Response().Status
		msg := "Request"

		switch {
		case statusCode >= http.StatusBadRequest && statusCode < http.StatusInternalServerError:
			logger.WithLevel(zerolog.ErrorLevel).Msg(msg)
		case statusCode >= http.StatusInternalServerError:
			logger.WithLevel(zerolog.FatalLevel).Msg(msg)
		default:
			logger.WithLevel(zerolog.InfoLevel).Msg(msg)
		}

		return nil
	}
}

func ErrCtx(ctx context.Context, err error, label string, msg ...string) {
	logger := log.Error().Str("label", label).Err(err)
	if reqID := GetRequestID(ctx); reqID != "" {
		logger = logger.Str(string(RequestIDHeaderKey), reqID)
	}
	logger.Msg(strings.Join(msg, ". "))
}

func ErrWrapCtx(ctx context.Context, err error, label string, msg ...string) error {
	ErrCtx(ctx, err, label, msg...)
	return fmt.Errorf("%s: %w", label, err)
}

func ErrWrap(err error, label, msg string) error {
	log.Error().Err(err).Str("label", label).Msg(msg)
	return fmt.Errorf("%s: %w", label, err)
}

func ErrLoggerWrapCtx(ctx context.Context, label string) *zerolog.Event {
	logger := log.Error().Str("label", label)
	if reqID := GetRequestID(ctx); reqID != "" {
		logger = logger.Str(string(RequestIDHeaderKey), reqID)
	}
	return logger
}

func InfoCtx(ctx context.Context, label string, msg ...string) {
	logger := log.Info().Str("label", label)
	if reqID := GetRequestID(ctx); reqID != "" {
		logger = logger.Str(string(RequestIDHeaderKey), reqID)
	}
	logger.Msg(strings.Join(msg, ". "))
}
