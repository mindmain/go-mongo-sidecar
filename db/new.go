package db

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func NewMongoHandler(client *mongo.Client) HandlerMongoReplicaSet {
	return &mongoHandler{
		client: client,
	}
}

type mongoHandler struct {
	client *mongo.Client
}
