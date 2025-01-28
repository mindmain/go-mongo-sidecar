package service

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mindmain/go-mongo-sidecar/db"
	"github.com/mindmain/go-mongo-sidecar/k8s"
)

func getPodName(host string) string {
	return strings.Split(host, ".")[0]
}

func (s *sidecarService) printStatus(status *db.ReplicaSetStatus, pods []*k8s.MongoPod) {

	hostname, err := os.Hostname()

	if err != nil {
		log.Println("[ERROR] error on get hostname: ", err)
		return
	}

	var states []string

	for _, member := range status.Members {
		states = append(states, fmt.Sprintf("%s (%s)", getPodName(member.Name), member.StateStr))
	}

	printResult := make([]string, len(pods))

	for i, pod := range pods {
		printResult[i] = fmt.Sprintf("%s (%s)", pod.Name, pod.IP)
	}

	ns := fmt.Sprintf("[INFO]\nDetect change status: sidecar %s i'm %s\nreplica members: %s\nmatched pods: %s", hostname, s.serviceRole, strings.Join(states, ", "), strings.Join(printResult, ", "))

	if s.state != ns {
		log.Println(ns)
		s.state = ns
	}

}
