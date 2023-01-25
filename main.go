package main

import (
	"context"

	"time"

	"github.com/labstack/echo/v4"
	"github.com/mrkonz/domain/catagory"
	"github.com/mrkonz/domain/catalog"
	"github.com/mrkonz/domain/health"
	"github.com/mrkonz/domain/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {

	ctx, cancell := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancell()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:password@localhost:27017"))
	err = client.Ping(ctx, readpref.Primary())

	mongodb := client.Database("sxexpo")

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	if err != nil {
		panic(err)
	}
	e := echo.New()

	e.GET("/health", health.HealthCheckHandler())
	e.GET("/catalogs", catalog.GetCatalogHandler(catalog.GetCatalog(mongodb)))
	e.GET("/catagories", catagory.GetCatagoryHandler(catagory.GetCatagory(mongodb)))
	e.GET("/type", types.GetTypeHandler(types.GetType(mongodb)))

	e.Logger.Fatal(e.Start(":1323"))
}
