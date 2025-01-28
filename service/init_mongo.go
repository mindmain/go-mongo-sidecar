package service

import (
	"context"
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
		if err = s.mongoHandler.Init(ctx, pods); err != nil {
			return err
		} else {
			return nil
		}

	}

}
