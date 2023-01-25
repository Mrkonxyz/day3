package types

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetType(db *mongo.Database) func(context.Context) ([]Type, error) {
	return func(ctx context.Context) ([]Type, error) {

		collecttion := db.Collection("privilege_type")

		cur, err := collecttion.Find(ctx, bson.D{})

		if err != nil {
			return nil, err
		}
		var result []Type
		if err := cur.All(ctx, &result); err != nil {
			return nil, err
		}
		return result, nil
	}
}
