package db

import (
	"context"

	"github.com/mindmain/go-mongo-sidecar/types"
)

func (h *mongoHandler) IsInitialized(ctx context.Context) (bool, error) {

	status, err := h.Status(ctx)

	if err != nil {

		if types.ErrorNoReplicaSetConfig.Match(err) {
			return false, nil
		}

		return false, err
	}

	return status.SetName == types.MONGO_REPLICA_SET.Get(), nil

}
