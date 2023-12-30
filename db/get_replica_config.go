package db

import "context"

func (h *mongoHandler) GetReplicaSetConfig(ctx context.Context) (*replicaSetConfig, error) {

	var result replicaSetConfig

	res := h.client.Database("admin").RunCommand(ctx, map[string]string{"replSetGetConfig": "1"}).Decode(&result)

	if res != nil {
		return nil, res
	}

	return &result, nil

}
