package db

import "context"

type getConfigResult struct {
	Config replicaSetConfig `bson:"config"`
}

func (h *mongoHandler) GetReplicaSetConfig(ctx context.Context) (*replicaSetConfig, error) {

	var result getConfigResult

	res := h.client.Database("admin").RunCommand(ctx, map[string]string{"replSetGetConfig": "1"}).Decode(&result)

	if res != nil {
		return nil, res
	}

	return &result.Config, nil

}
