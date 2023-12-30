package db

import (
	"context"
	"time"

	"github.com/mindmain/go-mongo-sidecar/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnLocal() (*mongo.Client, error) {

	var port = types.MONGO_PORT.Get()
	var host = types.MONGO_HOST.Get()

	var psw = types.MONGO_PASSWORD.Get()
	var user = types.MONGO_USER.Get()

	var uri = ""
	if user == "" && psw == "" {
		uri = "mongodb://" + host + ":" + port
	} else {
		uri = "mongodb://" + user + ":" + psw + "@" + host + ":" + port + "/?authSource=admin"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetDirect(true))

	if err != nil {
		return nil, err
	}

	return client, nil
}
