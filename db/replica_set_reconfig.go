package db

import (
	"context"

	"github.com/mindmain/go-mongo-sidecar/k8s"
	"github.com/mindmain/go-mongo-sidecar/types"
	"go.mongodb.org/mongo-driver/bson"
)

func (h *mongoHandler) Reconfig(ctx context.Context, hosts []*k8s.MongoPod) error {

	var ok messageOk

	oldConfig, err := h.getConfig(ctx)

	if err != nil {
		return err
	}

	var config = bson.D{
		{
			Key: "replSetReconfig",
			Value: bson.M{
				"_id":     types.MONGO_REPLICA_SET,
				"members": hostsToMembers(hosts),
				"version": oldConfig.Version + 1,
			},
		},
		{Key: "force", Value: true},
	}

	res := h.client.Database("admin").RunCommand(ctx, config)

	if err := res.Decode(&ok); err != nil {
		return err
	}

	if !ok.Ok {
		return types.ErrorNotRemoveMember
	}

	return nil

}
