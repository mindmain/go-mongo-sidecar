package db

import (
	"context"
	"strings"
)

type HandlerMongoReplicaSet interface {
	ReplicaSetReconfig(ctx context.Context, hosts []string) error
	IsInitialized(ctx context.Context) (bool, error)
	IsPrimary(ctx context.Context) (bool, error)
	IsSecondary(ctx context.Context) (bool, error)
	GetReplicaSetConfig(ctx context.Context) (*replicaSetConfig, error)
	InitReplicaSet(ctx context.Context, hsots []string) error
}

type isMasterResult struct {
	IsMaster    bool `bson:"ismaster"`
	IsSecondary bool `bson:"secondary"`
}

type replicaSetMember struct {
	ID      int    `bson:"_id"`
	Host    string `bson:"host"`
	Healthy *int   `bson:"health,omitempty"`
}

type replicaSetConfig struct {
	ID      string              `bson:"_id"`
	Version int                 `bson:"version"`
	Members []*replicaSetMember `bson:"members"`
}

func (r *replicaSetConfig) GetMember(host string) *replicaSetMember {
	for _, member := range r.Members {
		if member.Host == host {
			return member
		}
	}

	return nil
}

func (r *replicaSetConfig) IsMember(host string) bool {
	for _, member := range r.Members {
		if member.Host == host || strings.Contains(member.Host, host) {
			return true
		}
	}

	return false
}

func (r *replicaSetConfig) LengthMemberLive() int {

	var count int
	for _, member := range r.Members {
		if member.Healthy != nil && *member.Healthy == 1 {
			count++
		}
	}
	return count
}

func (r *replicaSetConfig) GetMemberNotLive() []*replicaSetMember {

	var members []*replicaSetMember
	for _, member := range r.Members {
		if member.Healthy != nil && *member.Healthy != 1 {
			members = append(members, member)
		}
	}
	return members
}

type messageOk struct {
	Ok bool `bson:"ok"`
}
