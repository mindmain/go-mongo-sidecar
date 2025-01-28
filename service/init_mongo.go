package service

import (
	"context"

	"github.com/mindmain/go-mongo-sidecar/types"
)

func (s *sidecarService) initMongo(ctx context.Context) error {

	isInitialized, err := s.mongoHandler.IsInitialized(ctx)

	if err != nil {
		return err
	}

	if isInitialized {
		return nil
	} else {

		var err error

		pods, err := s.pods(ctx)

		if err != nil {
			return err
		}
		if err = s.mongoHandler.Init(ctx, addServiceToPodsNames(pods, types.HEADLESS_SERVICE)); err != nil {
			return err
		} else {
			return nil
		}

	}

}
