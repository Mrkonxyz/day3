package catagory

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Catagory struct {
	ID       string `bson:"_id" json:"id"`
	Category string `json:"catagory"`
}

type getCatagoryFunc func(context.Context) ([]Catagory, error)

func (fn getCatagoryFunc) getCatagoryFunc(ctx context.Context) ([]Catagory, error) {
	return fn(ctx)
}

func GetCatagoryHandler(svg getCatagoryFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		results, err := svg.getCatagoryFunc(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, echo.Map{"data": results})
	}

}
