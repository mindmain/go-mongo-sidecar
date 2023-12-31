package service

import (
	"context"

	"github.com/mindmain/go-mongo-sidecar/types"
)

func (s *sidecarService) pods(ctx context.Context) ([]string, error) {

	// make selector to get pods with label es. app=mongo
	selector, err := stringLabelToMap(types.SIDECAR_SELECTOR_POD.Get())
	if err != nil {
		return nil, err
	}

	pods, err := s.k8sHandler.GetPodsNamesWithMatchLabels(ctx, selector)

	if err != nil {
		return nil, err
	}

	return pods, nil

}
