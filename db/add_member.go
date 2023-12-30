package db

import (
	"context"

	"github.com/mindmain/go-mongo-sidecar/types"
)

func (h *mongoHandler) AddMember(ctx context.Context, host string) error {

	var ok messageOk
	res := h.client.Database("admin").RunCommand(ctx, map[string]string{"replSetAdd": host})

	if err := res.Decode(&ok); err != nil {
		return err
	}

	if !ok.Ok {
		return types.ErrorNotAddMember
	}

	return nil
}
