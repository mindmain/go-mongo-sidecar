package db

import "context"

func (h *mongoHandler) Status(ctx context.Context) (*ReplicaSetStatus, error) {

	var result ReplicaSetStatus

	err := h.client.Database("admin").RunCommand(ctx, map[string]string{"replSetGetStatus": "1"}).Decode(&result)

	if err != nil {
		return nil, err
	}

	return &result, nil

}
