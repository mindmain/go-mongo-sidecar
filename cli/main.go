package main

import (
	"context"
	"log"

	"github.com/mindmain/go-mongo-sidecar/db"
	"github.com/mindmain/go-mongo-sidecar/k8s"
	"github.com/mindmain/go-mongo-sidecar/service"
)

func main() {

	client, err := k8s.KubeClient()

	if err != nil {
		log.Fatal(err)
	}
	kubeHandler := k8s.NewK8sHandler(client)
	mongoClient, err := db.MongoConnLocal()

	if err != nil {
		log.Fatal(err)
	}

	defer mongoClient.Disconnect(context.Background())

	handlerMongo := db.NewMongoHandler(mongoClient)

	sidecarService := service.NewSidecarService(handlerMongo, kubeHandler)

	if err := sidecarService.Run(context.Background()); err != nil {
		log.Fatal(err)
	}

}
