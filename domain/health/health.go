package health

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HealthCheckHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, echo.Map{"Status": "It's All Good, API is Working :)"})
	}
}
