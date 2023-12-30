package db

import "context"

func (h *mongoHandler) InitReplicaSet(ctx context.Context, hosts []string) error {

	var members []*replicaSetMember

	for i, host := range hosts {
		members = append(members, &replicaSetMember{
			ID:   i,
			Host: host,
		})
	}

	res := h.client.Database("admin").RunCommand(ctx, map[string]interface{}{
		"replSetInitiate": map[string]interface{}{
			"members": members,
		},
	})

	return res.Err()
}
