package db

import (
	"context"

	"github.com/mindmain/go-mongo-sidecar/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (h *mongoHandler) IsInitialized(ctx context.Context) (bool, error) {
	var out ReplicaSetConfig
	res := h.client.Database("admin").RunCommand(ctx, bson.M{
		"replSetGetStatus": 1,
	}, options.RunCmd()).Decode(&out)

	if res != nil {
		if types.ErrorNoReplicaSetConfig.Match(res) {
			return false, nil
		}
		return false, res
	}

	return out.ID != types.MONGO_REPLICA_SET.Get(), nil
}
