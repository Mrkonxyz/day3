package catagory

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCatagory(db *mongo.Database) func(context.Context) ([]Catagory, error) {
	return func(ctx context.Context) ([]Catagory, error) {

		collecttion := db.Collection("categories")

		cur, err := collecttion.Find(ctx, bson.D{})

		if err != nil {
			return nil, err
		}
		var result []Catagory
		if err := cur.All(ctx, &result); err != nil {
			return nil, err
		}
		return result, nil
	}
}
