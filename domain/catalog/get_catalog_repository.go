package catalog

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCatalog(db *mongo.Database) func(context.Context) ([]Catalog, error) {
	return func(ctx context.Context) ([]Catalog, error) {

		collecttion := db.Collection("catalogs")

		cur, err := collecttion.Find(ctx, bson.D{})

		if err != nil {
			return nil, err
		}
		var result []Catalog
		if err := cur.All(ctx, &result); err != nil {
			return nil, err
		}
		return result, nil
	}
}
