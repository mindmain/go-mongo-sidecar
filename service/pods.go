package service

import (
	"context"

	"github.com/mindmain/go-mongo-sidecar/k8s"
	"github.com/mindmain/go-mongo-sidecar/types"
)

func (s *sidecarService) pods(ctx context.Context) ([]*k8s.MongoPod, error) {

	// make selector to get pods with label es. app=mongo
	selector, err := stringLabelToMap(types.SIDECAR_SELECTOR_POD)
	if err != nil {
		return nil, err
	}

	pods, err := s.k8sHandler.GetPodsNamesWithMatchLabels(ctx, selector)

	if err != nil {
		return nil, err
	}

	return pods, nil

}
