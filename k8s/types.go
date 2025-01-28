package k8s

import (
	"context"
)

type HandlerKubernetes interface {
	GetPodsNamesWithMatchLabels(ctx context.Context, labels map[string]string) ([]*MongoPod, error)
}
