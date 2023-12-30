package db

func hostsToMembers(hosts []string) []*replicaSetMember {
	var members []*replicaSetMember

	for i, host := range hosts {
		members = append(members, &replicaSetMember{
			ID:   i,
			Host: host,
		})
	}

	return members
}
