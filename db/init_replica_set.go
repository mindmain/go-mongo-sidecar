package db

import (
	"context"

	"github.com/mindmain/go-mongo-sidecar/types"
	"go.mongodb.org/mongo-driver/bson"
)

func (h *mongoHandler) Init(ctx context.Context, hosts []string) error {

	res := h.client.Database("admin").RunCommand(ctx, bson.D{
		{
			Key: "replSetInitiate",
			Value: bson.M{
				"members": hostsToMembers(hosts),
				"_id":     types.MONGO_REPLICA_SET.Get(),
			},
		},
	})

	return res.Err()
}
