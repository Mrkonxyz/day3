package catalog

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Catalog struct {
	ID      string `bson:"_id" json:"id"`
	Catalog string `json:"catalog"`
}

type getCatalogFunc func(context.Context) ([]Catalog, error)

func (fn getCatalogFunc) getCatalogFunc(ctx context.Context) ([]Catalog, error) {
	return fn(ctx)
}

func GetCatalogHandler(svg getCatalogFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		results, err := svg.getCatalogFunc(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, echo.Map{"data": results})
	}

}
