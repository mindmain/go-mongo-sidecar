package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mindmain/go-mongo-sidecar/db"
	"github.com/mindmain/go-mongo-sidecar/k8s"
	"github.com/mindmain/go-mongo-sidecar/types"
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
	mongoHandler db.HandlerMongoReplicaSet
	k8sHandler   k8s.HandlerKubernetes
}

func (s *sidecarService) Run(ctx context.Context) error {

	if s.k8sHandler == nil {
		return fmt.Errorf("k8s handler is nil")
	}

	if s.mongoHandler == nil {
		return fmt.Errorf("mongo handler is nil")
	}

	if types.MONGO_REPLICA_SET.Get() == "" {
		return fmt.Errorf("mongo replica set is empty please set name of replica set in the env variable: %s", string(types.MONGO_REPLICA_SET))
	}

	selector, err := stringLabelToMap(types.SIDECAR_SELECTOR_POD.Get())

	if err != nil {
		return err
	}

	pods, err := s.k8sHandler.GetPodsNamesWithMatchLabels(ctx, selector)

	if err != nil {
		return err
	}

	if len(pods) == 0 {
		return fmt.Errorf("not found pods with label app=mongo")
	}

	log.Println("[INFO] found pods: ", pods)

	sleep := time.Second * time.Duration(types.SIDECAR_TIME_SLEEP.Int64())
	wait := time.Second * time.Duration(types.SIDECAR_TIME_TO_WAIT.Int64())

	if sleep <= 0 {
		sleep = time.Second * 5
	}

	if wait <= 0 {
		wait = time.Second * 10
	}

	time.Sleep(wait)

	for {
		time.Sleep(sleep)

		pods, err = s.k8sHandler.GetPodsNamesWithMatchLabels(ctx, selector)

		if err != nil {
			log.Println("[WARN] error to get pods names with match labels: ", err)
			continue
		}

		hosts := addServiceToPodsNames(pods, types.HEADLESS_SERVICE.Get())

		if len(hosts) == 0 {
			log.Println("[WARN] not found hosts")
			continue
		}

		if isInitialized, err := s.mongoHandler.IsInitialized(ctx); err != nil {
			log.Println("[WARN] error to check mongo replica set is initialized: ", err)
			continue
		} else {
			if !isInitialized {
				log.Println("[INFO] mongo replica set is not initialized")
				if err := s.mongoHandler.InitReplicaSet(ctx, hosts); err != nil {
					log.Println("[WARN] error to init replica set: ", err)
					continue
				}
			} else {

				if isPrimary, err := s.mongoHandler.IsPrimary(ctx); err != nil {

					log.Println("[WARN] error to check mongo replica set is primary: ", err)
					continue
				} else {
					if isPrimary {

						replicaConfig, err := s.mongoHandler.GetReplicaSetConfig(ctx)

						if err != nil {
							log.Println("[WARN] error to get replica set config: ", err)
							continue
						}

						mongoMembersLive := replicaConfig.LengthMemberLive()
						morePodsOfMembers := len(hosts) > mongoMembersLive
						lessPodsOfMembers := len(hosts) < mongoMembersLive

						if morePodsOfMembers || lessPodsOfMembers {

							if morePodsOfMembers {
								log.Printf("[INFO] more pods of members, pods: %d members: %d ", len(hosts), mongoMembersLive)
							}

							if lessPodsOfMembers {
								log.Printf("[INFO] less pods of members, pods: %d members: %d ", len(hosts), mongoMembersLive)
							}

							if err := s.mongoHandler.ReplicaSetReconfig(ctx, hosts); err != nil {
								log.Println("[WARN] error to reconfig replica set: ", err)
								continue
							}

						}

					}
				}

			}

		}

	}

}
