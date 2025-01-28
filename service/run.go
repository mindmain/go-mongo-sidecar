package service

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/mindmain/go-mongo-sidecar/types"
)

func (s *sidecarService) Run(ctx context.Context) error {

	// check if all dependencies are ok and set env variables
	s.check()
	s.initDuration()

	log.Printf("[INFO] starting sidecar time to wait mongo startup %.2f seconds check status of replica set every %.2f seconds", s.waitDuration.Seconds()+s.sleepDuration.Seconds(), s.sleepDuration.Seconds())
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

		isPrimary, err := s.mongoHandler.IsPrimary(ctx)
		if err != nil {
			log.Println("[WARN] error to get primary status: ", err)
			continue
		}

		isSecondary, err := s.mongoHandler.IsSecondary(ctx)

		if err != nil {
			log.Println("[WARN] error to get secondary status: ", err)
			continue
		}

		if isPrimary {
			s.serviceRole = primary
		} else if isSecondary {
			s.serviceRole = secondary
		} else {
			s.serviceRole = unknown
		}

		if isPrimary {

			hosts := addServiceToPodsNames(pods, types.HEADLESS_SERVICE)
			mongoMembersLive := len(status.Members)
			morePodsOfMembers := len(hosts) > mongoMembersLive
			lessPodsOfMembers := len(hosts) < mongoMembersLive
			if morePodsOfMembers || lessPodsOfMembers {
				if morePodsOfMembers {
					log.Printf("[INFO] more pods of members, pods: %d members: %d ", len(hosts), mongoMembersLive)
					for _, h := range hosts {
						log.Println("[pods found]", h)
					}
				}
				if lessPodsOfMembers {
					log.Printf("[INFO] less pods of members, pods: %d members: %d ", len(hosts), mongoMembersLive)
					for _, h := range hosts {
						log.Println("[pods found]", h)
					}
				}
				if err := s.mongoHandler.Reconfig(ctx, pods); err != nil {
					log.Println("[WARN] error to reconfig replica set: ", err)
					continue
				} else {
					log.Println("[INFO] replica set reconfigured")
				}
			}

		}

		// if primary not exists, force reconfig
		if isSecondary {
			notPrimaryMember := 0
			for _, member := range status.Members {
				if member.StateStr != "PRIMARY" {
					notPrimaryMember++
				}
			}

			hostname, err := os.Hostname()

			if err != nil {
				log.Println("[WARN] error to get hostname: ", err)
				continue
			}

			isPod0 := strings.Contains(hostname, "0")

			// if not primary member is equal to members and hostname contains 0, force reconfig from this pod
			if notPrimaryMember == len(status.Members) {
				if isPod0 {
					log.Println("[INFO] primary not exists, force reconfig")
					if err := s.mongoHandler.Reconfig(ctx, pods); err != nil {
						log.Println("[WARN] error to reconfig replica set: ", err)
						continue
					} else {
						log.Printf("[INFO] replica set reconfigured")
					}
				} else {
					log.Println("[INFO] primary not exists, but this pod is not 0, freeze replica set to 5 seconds")
					if err := s.mongoHandler.Freeze(ctx, 5); err != nil {
						log.Println("[WARN] error to freeze replica set: ", err)
						continue
					} else {
						log.Printf("[INFO] replica set freezed")
					}
				}
			}
		}

		s.printStatus(status, pods)
	}

}
