package db

import (
	"strconv"
	"strings"

	"github.com/mindmain/go-mongo-sidecar/k8s"
	"github.com/mindmain/go-mongo-sidecar/types"
)

type configurationReplicaMember struct {
	ID   int    `bson:"_id"`
	Host string `bson:"host"`
}

func hostsToMembers(hosts []*k8s.MongoPod) []*configurationReplicaMember {
	var members = make([]*configurationReplicaMember, 0, len(hosts))

	for _, host := range hosts {
		members = append(members, &configurationReplicaMember{
			ID:   getNumberStatefulSet(host.Name.String()),
			Host: host.Name.WithService(types.HEADLESS_SERVICE),
		})
	}

	return members
}

func getNumberStatefulSet(host string) int {

	hostname := strings.Split(host, ".")[0]

	number, err := strconv.Atoi(strings.Split(hostname, "-")[len(strings.Split(hostname, "-"))-1])

	if err != nil {
		return 0
	}

	return number

}
