package db

import "context"

func (h *mongoHandler) IsSecondary(ctx context.Context) (bool, error) {

	var result isMasterResult
	res := h.client.Database("admin").RunCommand(ctx, map[string]string{"isMaster": "1"}).Decode(&result)

	if res != nil {
		return false, res
	}

	if result.IsSecondary {
		return true, nil
	}

	return false, nil

}
