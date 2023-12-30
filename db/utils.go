package db

import (
	"strconv"
	"strings"
)

type configurationReplicaMember struct {
	ID   int    `bson:"_id"`
	Host string `bson:"host"`
}

func hostsToMembers(hosts []string) []*configurationReplicaMember {
	var members = make([]*configurationReplicaMember, 0, len(hosts))

	for _, host := range hosts {
		members = append(members, &configurationReplicaMember{
			ID:   getNumberStatefulSet(host),
			Host: host,
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
