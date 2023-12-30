package db

import (
	"context"

	"github.com/mindmain/go-mongo-sidecar/types"
	"go.mongodb.org/mongo-driver/bson"
)

func (h *mongoHandler) ReplicaSetReconfig(ctx context.Context, hosts []string) error {

	var ok messageOk

	res := h.client.Database("admin").RunCommand(ctx, bson.D{
		{
			Key: "replSetReconfig",
			Value: bson.M{
				"_id":     types.MONGO_REPLICA_SET.Get(),
				"members": hostsToMembers(hosts),
			},
		},
		{Key: "force", Value: true},
	})

	if err := res.Decode(&ok); err != nil {
		return err
	}

	if !ok.Ok {
		return types.ErrorNotRemoveMember
	}

	return nil

}
