package db

import "context"

func (h *mongoHandler) Status(ctx context.Context) (*ReplicaSetStatus, error) {

	var result ReplicaSetStatus

	res := h.client.Database("admin").RunCommand(ctx, map[string]string{"replSetGetStatus": "1"}).Decode(&result)

	if res != nil {
		return nil, res
	}

	return &result, nil

}
