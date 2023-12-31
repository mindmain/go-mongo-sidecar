# Go mongo sidecar

Go mongo sidecar is an application written in golang, which using the sidecar patter provides for the realignment and configuration of a mongo replicaSet.

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
