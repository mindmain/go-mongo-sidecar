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

						if morePodsOfMembers {

							newHosts := []string{}

							for _, host := range hosts {
								if !replicaConfig.IsMember(host) {
									newHosts = append(newHosts, host)
								}
							}

							for _, host := range newHosts {
								if err := s.mongoHandler.AddMember(ctx, host); err != nil {
									log.Println("[WARN] error to add member: ", host, err)
									continue
								} else {
									log.Println("[INFO] new replica member added: ", host)
								}
							}

						}

						if lessPodsOfMembers {

							for _, host := range hosts {
								if !replicaConfig.IsMember(host) {
									if err := s.mongoHandler.RemoveMember(ctx, host); err != nil {
										log.Println("[WARN] error to remove member: ", host, err)
										continue
									} else {
										log.Println("[INFO] replica member removed: ", host)
									}
								}
							}

						}

					}
				}

			}

		}

		time.Sleep(sleep)

	}

}
