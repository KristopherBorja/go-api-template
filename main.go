package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "OK",
		})
	})

	e.POST("/user", func(c echo.Context) error {
		u := new(User)

		if err := c.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
		}

		if err := c.Validate(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, u)
	})

	// func show(c echo.Context) error {
	// 	// Get team and member from the query string
	// 	team := c.QueryParam("team")
	// 	member := c.QueryParam("member")
	// 	return c.String(http.StatusOK, "team:" + team + ", member:" + member)
	// }

	e.Logger.Fatal(e.Start(":8080"))
}
