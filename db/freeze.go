package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *mongoHandler) Freeze(ctx context.Context, sec int) error {

	res := r.client.Database("admin").RunCommand(ctx, bson.D{{Key: "replSetFreeze", Value: sec}})

	return res.Err()

}
