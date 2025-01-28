package db

import (
	"context"

	"github.com/mindmain/go-mongo-sidecar/k8s"
	"github.com/mindmain/go-mongo-sidecar/types"
	"go.mongodb.org/mongo-driver/bson"
)

func (h *mongoHandler) Init(ctx context.Context, hosts []*k8s.MongoPod) error {

	res := h.client.Database("admin").RunCommand(ctx, bson.D{
		{
			Key: "replSetInitiate",
			Value: bson.M{
				"members": hostsToMembers(hosts),
				"_id":     types.MONGO_REPLICA_SET,
			},
		},
	})

	return res.Err()
}
