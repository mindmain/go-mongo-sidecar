package db

import (
	"context"
	"strings"
)

type HandlerMongoReplicaSet interface {
	Reconfig(ctx context.Context, hosts []string) error
	IsInitialized(ctx context.Context) (bool, error)
	IsPrimary(ctx context.Context) (bool, error)
	IsSecondary(ctx context.Context) (bool, error)
	Status(ctx context.Context) (*ReplicaSetStatus, error)
	Init(ctx context.Context, hsots []string) error
}

type isMasterResult struct {
	IsMaster    bool `bson:"ismaster"`
	IsSecondary bool `bson:"secondary"`
}

type ReplicaMemberConfig struct {
	ID     int    `bson:"_id"`
	Host   string `bson:"host"`
	Health int    `bson:"health,omitempty"`
}

type ReplicaSetConfig struct {
	ID      string                 `bson:"_id"`
	Version int                    `bson:"version"`
	Members []*ReplicaMemberConfig `bson:"members"`
}

func (r *ReplicaSetConfig) GetMember(host string) *ReplicaMemberConfig {
	for _, member := range r.Members {
		if member.Host == host {
			return member
		}
	}

	return nil
}

func (r *ReplicaSetConfig) IsMember(host string) bool {
	for _, member := range r.Members {
		if member.Host == host || strings.Contains(member.Host, host) {
			return true
		}
	}

	return false
}

func (r *ReplicaSetConfig) LengthMemberLive() int {

	var count int
	for _, member := range r.Members {
		if member.Health == 1 {
			count++
		}
	}
	return count
}

func (r *ReplicaSetConfig) GetMemberNotLive() []*ReplicaMemberConfig {

	var members []*ReplicaMemberConfig
	for _, member := range r.Members {
		if member.Health != 1 {
			members = append(members, member)
		}
	}
	return members
}

func (r *ReplicaSetConfig) MembersNames() []string {

	var members []string

	for _, member := range r.Members {
		members = append(members, strings.Split(member.Host, ".")[0])
	}

	return members

}

type ReplocaMemberStatus struct {
	ID       int    `bson:"_id"`
	Name     string `bson:"name"`
	Health   int    `bson:"health"`
	State    int    `bson:"state"`
	StateStr string `bson:"stateStr"`
}
type ReplicaSetStatus struct {
	SetName string                 `bson:"set"`
	Members []*ReplocaMemberStatus `bson:"members"`
	Ok      int                    `bson:"ok"`
}

func (r *ReplicaSetStatus) GetMember(host string) *ReplocaMemberStatus {
	for _, member := range r.Members {
		if member.Name == host {
			return member
		}
	}

	return nil
}

func (r *ReplicaSetStatus) IsMember(host string) bool {
	for _, member := range r.Members {
		if member.Name == host || strings.Contains(member.Name, host) {
			return true
		}
	}

	return false
}

func (r *ReplicaSetStatus) MembersNames() []string {

	var members []string

	for _, member := range r.Members {
		members = append(members, strings.Split(member.Name, ".")[0])
	}

	return members
}

func (r *ReplicaSetStatus) IsPrimary(host string) bool {
	for _, member := range r.Members {
		if member.Name == host {
			return member.State == 1
		}
	}

	return false
}

func (r *ReplicaSetStatus) IsSecondary(host string) bool {
	for _, member := range r.Members {
		if member.Name == host {
			return member.State == 2
		}
	}

	return false
}

func (r *ReplicaSetStatus) MembersPrintStatus() string {

	var status []string

	for _, member := range r.Members {
		status = append(status, strings.Join([]string{member.Name, member.StateStr}, ":"))
	}

	return strings.Join(status, " - ")

}

func (r *ReplicaSetStatus) LengthMemberLive() int {

	var count int
	for _, member := range r.Members {
		if member.Health == 1 {
			count++
		}
	}
	return count
}

type messageOk struct {
	Ok bool `bson:"ok"`
}
