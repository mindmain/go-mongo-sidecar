package service

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mindmain/go-mongo-sidecar/db"
)

func getpodName(host string) string {
	return strings.Split(host, ".")[0]
}

func (s *sidecarService) printStatus(status *db.ReplicaSetStatus, pods []string) {

	hostname, err := os.Hostname()

	if err != nil {
		log.Println("[ERROR] error on get hostname: ", err)
		return
	}

	var states []string

	for _, member := range status.Members {
		states = append(states, fmt.Sprintf("%s (%s)", getpodName(member.Name), member.StateStr))
	}

	ns := fmt.Sprintf("[INFO]\nDetect change status: sidecar %s i'm %s\nreplica members: %s\nmatched pods: %s", hostname, s.serviceRole, strings.Join(states, ", "), strings.Join(pods, ", "))

	if s.state != ns {
		log.Println(ns)
		s.state = ns
	}

}
