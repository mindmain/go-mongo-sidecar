package service

import (
	"context"

	"github.com/mindmain/go-mongo-sidecar/types"
)

func (s *sidecarService) initMongo(ctx context.Context) error {

	if s.isInitialized {
		return nil
	}

	isInitialized, err := s.mongoHandler.IsInitialized(ctx)

	if err != nil {
		return err
	}

	if isInitialized {
		s.isInitialized = true
		return nil
	} else {

		var err error

		pods, err := s.pods(ctx)

		if err != nil {
			return err
		}
		if err = s.mongoHandler.Init(ctx, addServiceToPodsNames(pods, types.HEADLESS_SERVICE.Get())); err != nil {
			return err
		} else {
			s.isInitialized = true
			return nil
		}

	}

}
