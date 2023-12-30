package db

import (
	"strconv"
	"strings"
)

func hostsToMembers(hosts []string) []*replicaSetMember {
	var members []*replicaSetMember

	for _, host := range hosts {
		members = append(members, &replicaSetMember{
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
