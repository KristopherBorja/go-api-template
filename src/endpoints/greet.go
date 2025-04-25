package endpoints

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

func RegisterRoutes(e *echo.Echo) {
	e.GET("/health", health)
	e.POST("/user", greetUser)

}
func health(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status": "OK",
	})
}

func greetUser(c echo.Context) error {
	u := new(User)

	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	if err := c.Validate(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, u)
}

// func show(c echo.Context) error {
// Get team and member from the query string
// 	team := c.QueryParam("team")
// 	member := c.QueryParam("member")
// 	return c.String(http.StatusOK, "team:" + team + ", member:" + member)
// }
