# Go mongo sidecar

Go mongo sidecar is an application written in golang, which using the sidecar patter provides for the realignment and configuration of a mongo replicaSet.

## How use it

use this images to use go-mongo-sidecar `ghcr.io/mindmain/go-mongo-sidecar:0.1.0`.
You must define follow environment variables:

* MONGO_REPLICA_SET
* HEADLESS_SERVICE
* SIDECAR_SELECTOR_POD

see [examples](./examples/minikube.yaml)

## How does it work

Each go sidecar has a wait time and a pause time, only the node that is recognized as primary will make reconfiguration changes to the mongo replicaSet.
They will be read through the kubernets client through a selector with labels.
Then they will be passed to mongo which receives the configuration.

Every cycle check status and the conditions change, the sidecar will report the change in the logs like the following example:

```text
2023/12/31 15:49:39 [INFO]
Detect change status: sidecar mongo-0 i'm primary
replica members: mongo-0 (PRIMARY), mongo-1 (SECONDARY), mongo-2 (SECONDARY)
matched pods: mongo-0, mongo-1, mongo-2
```

## environment variables

MONGO_REPLICA_SET is **REQUIRED** is the name of replicaset used on command startup mongo.

|Name|Description| Default value |
|---|---|---|
|MONGO_PORT| set the mongo will use to connect |27017|
|MONGO_HOST| host used to connect to service mongo managed of sidecar  |localhost|
|MONGO_USER| if connection required authentication |*empty*|
|MONGO_PASSWORD| if connection required authentication |*empty*|
|MONGO_REPLICA_SET| name of replica set used into commandline stratup mongod |*empity*|
|HEADLESS_SERVICE| complete dns of service clusterIp none  |mongo.default.svc.cluster.local|
|SIDECAR_TIME_TO_WAIT| time must wait sidecar before run|30|
|SIDECAR_TIME_SLEEP| How much time must pass between one status request and another |5|
|SIDECAR_SELECTOR_POD| selector to localize pods inside cluster (many times it coincides with the match label of the service) |app=mongo|
