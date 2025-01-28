package service

import (
	"fmt"

	"github.com/mindmain/go-mongo-sidecar/types"
)

func (s *sidecarService) check() error {
	if s.k8sHandler == nil {
		return fmt.Errorf("k8s handler is nil")
	}

	if s.mongoHandler == nil {
		return fmt.Errorf("mongo handler is nil")
	}

	if types.MONGO_REPLICA_SET == "" {
		return fmt.Errorf("mongo replica set is empty please set name of replica set in the env variable: %s", string(types.MONGO_REPLICA_SET))
	}

	return nil
}
