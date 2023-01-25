package types

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Type struct {
	ID   string `bson:"_id" json:"id"`
	Type string `json:"catagory"`
}

type getTypeFunc func(context.Context) ([]Type, error)

func (fn getTypeFunc) getTypeFunc(ctx context.Context) ([]Type, error) {
	return fn(ctx)
}

func GetTypeHandler(svg getTypeFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		results, err := svg.getTypeFunc(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, echo.Map{"data": results})
	}

}
