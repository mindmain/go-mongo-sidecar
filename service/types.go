package service

import (
	"context"
	"time"

	"github.com/mindmain/go-mongo-sidecar/db"
	"github.com/mindmain/go-mongo-sidecar/k8s"
)

type role string

const (
	primary   role = "primary"
	secondary role = "secondary"
	unknown   role = "unknown"
)

type SidecarService interface {
	Run(ctx context.Context) error
}

func NewSidecarService(
	mongoHandler db.HandlerMongoReplicaSet,
	k8sHandler k8s.HandlerKubernetes,

) SidecarService {
	return &sidecarService{
		mongoHandler: mongoHandler,
		k8sHandler:   k8sHandler,
	}
}

type sidecarService struct {
	mongoHandler  db.HandlerMongoReplicaSet
	k8sHandler    k8s.HandlerKubernetes
	sleepDuration time.Duration
	waitDuration  time.Duration
	state         string
	serviceRole   role
}
