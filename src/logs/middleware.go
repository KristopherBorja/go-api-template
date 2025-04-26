package logs

import (
	"context"
	"time"

	"log/slog"

	"github.com/labstack/echo/v4"
)

func SlogMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			req := c.Request()
			res := c.Response()

			ctx := req.Context()

			ctx = context.WithValue(ctx, ContextKeyRequestID, res.Header().Get(echo.HeaderXRequestID))
			ctx = context.WithValue(ctx, ContextKeyRemoteIP, c.RealIP())
			ctx = context.WithValue(ctx, ContextKeyUserAgent, req.UserAgent())
			ctx = context.WithValue(ctx, ContextKeyURI, req.RequestURI)
			ctx = context.WithValue(ctx, ContextKeyMethod, req.Method)
			ctx = context.WithValue(ctx, ContextKeyHost, req.Host)

			c.SetRequest(req.WithContext(ctx))

			err := next(c)

			stop := time.Now()

			requestID, _ := ctx.Value(ContextKeyRequestID).(string)
			remoteIP, _ := ctx.Value(ContextKeyRemoteIP).(string)
			userAgent, _ := ctx.Value(ContextKeyUserAgent).(string)
			uri, _ := ctx.Value(ContextKeyURI).(string)
			method, _ := ctx.Value(ContextKeyMethod).(string)
			host, _ := ctx.Value(ContextKeyHost).(string)

			Logger.InfoContext(ctx, "Handled request",
				slog.String("time", time.Now().Format(time.RFC3339Nano)),
				slog.String("id", requestID),
				slog.String("remote_ip", remoteIP),
				slog.String("host", host),
				slog.String("method", method),
				slog.String("uri", uri),
				slog.String("user_agent", userAgent),
				slog.Int("status", res.Status),
				slog.Int64("bytes_in", req.ContentLength),
				slog.Int64("bytes_out", int64(res.Size)),
				slog.Duration("latency", stop.Sub(start)),
				slog.String("latency_human", stop.Sub(start).String()),
			)

			return err
		}
	}
}
