package main

import (
	"context"
	cfg "go-api-template/src/config"
	"go-api-template/src/endpoints"
	"go-api-template/src/logs"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// func init() {
// 	cfg.Red()

// }

func main() {
	config := cfg.Cfg()
	e := echo.New()

	logLevel := cfg.GetLogLevel(config.LogLevel)
	logs.InitLogger(slog.Level(logLevel))

	e.Use(middleware.RequestID()) // 📌 Add unique ID to all logs/errors early
	e.Use(logs.SlogMiddleware())
	e.Use(middleware.Recover())                                                              // 🛑 Catch panics before they crash the server
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))                  // 🚦 Enforce before any work is done
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{Timeout: 10 * time.Second})) // ⏱ Enforce max duration
	e.Use(middleware.CSRF())                                                                 // 🛡 Security: CSRF protection
	e.Use(middleware.CORS())                                                                 // 🌍 Cross-origin access
	e.Use(middleware.Secure())

	e.Validator = &CustomValidator{validator: validator.New()}

	endpoints.RegisterGreetingsRoutes(e)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	logs.Logger.Info("Staring application", slog.String("env", "production"))
	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			logs.Logger.Error("Server failed to start", slog.String("error", err.Error()))
			quit <- os.Interrupt
		}
	}()

	<-quit

	logs.Logger.Info("Shutdown signal received, exiting gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		logs.Logger.Error("Failed to gracefully shutdown server", slog.String("error", err.Error()))
	} else {
		logs.Logger.Info("Server shutdown completed OK")
	}
}
