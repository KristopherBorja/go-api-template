package main

import (
	cfg "go-api-template/src/config"
	"go-api-template/src/endpoints"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Custom validator that wraps go-playground/validator
type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func init() {
	cfg.Red()

}

func main() {
	config := cfg.Cfg()
	e := echo.New()

	logLevel := cfg.GetLogLevel(config.LogLevel)

	e.Logger.SetLevel(logLevel)
	e.Use(middleware.RequestID())                                                            // 📌 Add unique ID to all logs/errors early
	e.Use(middleware.Logger())                                                               // 📝 Log every request (with RequestID)
	e.Use(middleware.Recover())                                                              // 🛑 Catch panics before they crash the server
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))                  // 🚦 Enforce before any work is done
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{Timeout: 10 * time.Second})) // ⏱ Enforce max duration
	e.Use(middleware.CSRF())                                                                 // 🛡 Security: CSRF protection
	e.Use(middleware.CORS())                                                                 // 🌍 Cross-origin access
	e.Use(middleware.Secure())                                                               // 🔐 Security headers
	// Register validator
	e.Validator = &CustomValidator{validator: validator.New()}

	endpoints.RegisterGreetingsRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
