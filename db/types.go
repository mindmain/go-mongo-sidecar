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
	Status(ctx context.Context) (*ReplicaSetConfig, error)
	//getConfig(ctx context.Context) (*ReplicaSetConfig, error)
	Init(ctx context.Context, hsots []string) error
}

type isMasterResult struct {
	IsMaster    bool `bson:"ismaster"`
	IsSecondary bool `bson:"secondary"`
}

type ReplicaMember struct {
	ID     int    `bson:"_id"`
	Host   string `bson:"host"`
	Health int    `bson:"health,omitempty"`
}

type ReplicaSetConfig struct {
	ID      string           `bson:"_id"`
	Version int              `bson:"version"`
	Members []*ReplicaMember `bson:"members"`
}

func (r *ReplicaSetConfig) GetMember(host string) *ReplicaMember {
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

func (r *ReplicaSetConfig) GetMemberNotLive() []*ReplicaMember {

	var members []*ReplicaMember
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

type messageOk struct {
	Ok bool `bson:"ok"`
}
