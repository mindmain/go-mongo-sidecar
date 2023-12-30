package db

import (
	"context"

	"github.com/mindmain/go-mongo-sidecar/types"
	"go.mongodb.org/mongo-driver/bson"
)

func (h *mongoHandler) Reconfig(ctx context.Context, hosts []string) error {

	var ok messageOk

	oldConfig, err := h.getConfig(ctx)

	if err != nil {
		return err
	}

	res := h.client.Database("admin").RunCommand(ctx, bson.D{
		{
			Key: "replSetReconfig",
			Value: bson.M{
				"_id":     types.MONGO_REPLICA_SET.Get(),
				"members": hostsToMembers(hosts),
				"version": oldConfig.Version + 1,
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
