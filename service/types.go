package service

import (
	"context"
	"log"
	"time"

	"github.com/mindmain/go-mongo-sidecar/db"
	"github.com/mindmain/go-mongo-sidecar/k8s"
	"github.com/mindmain/go-mongo-sidecar/types"
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

func NewSidercarService(
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
	isInitialized bool
	state         string
	serviceRole   role
}

func (s *sidecarService) Run(ctx context.Context) error {

	// check if all dependencies are ok and setted env variables
	s.check()
	s.initDuration()

	log.Println("[INFO] starting sidecar ")
	s.wait()

	for {
		s.sleep()

		if err := s.initMongo(ctx); err != nil {
			log.Println("[ERROR] error on init mongo: ", err)
		}

		status, err := s.mongoHandler.Status(ctx)
		if err != nil {
			log.Println("[ERROR] error on get status: ", err)
			continue
		}
		pods, err := s.pods(ctx)

		if err != nil {
			log.Println("[ERROR] error on get pods: ", err)
			continue
		}

		s.printStatus(status, pods)

		if isPrimary, err := s.mongoHandler.IsPrimary(ctx); err != nil {
			log.Println("[WARN] error to get primary status: ", err)
			continue
		} else {

			if isPrimary {

				hosts := addServiceToPodsNames(pods, types.HEADLESS_SERVICE.Get())
				mongoMembersLive := status.LengthMemberLive()
				morePodsOfMembers := len(hosts) > mongoMembersLive
				lessPodsOfMembers := len(hosts) < mongoMembersLive
				if morePodsOfMembers || lessPodsOfMembers {
					if morePodsOfMembers {
						log.Printf("[INFO] more pods of members, pods: %d members: %d ", len(hosts), mongoMembersLive)
					}
					if lessPodsOfMembers {
						log.Printf("[INFO] less pods of members, pods: %d members: %d ", len(hosts), mongoMembersLive)
					}
					if err := s.mongoHandler.Reconfig(ctx, hosts); err != nil {
						log.Println("[WARN] error to reconfig replica set: ", err)
						continue
					} else {
						log.Println("[INFO] replica set reconfigured")
					}

				}

			}
		}

	}

}
