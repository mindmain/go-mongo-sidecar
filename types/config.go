package types

import (
	"cmp"
	"os"
)

var MONGO_PORT = cmp.Or(os.Getenv("MONGO_PORT"), "27017")
var MONGO_HOST = cmp.Or(os.Getenv("MONGO_HOST"), "localhost")
var MONGO_USER = cmp.Or(os.Getenv("MONGO_USER"), "")
var MONGO_REPLICA_SET = cmp.Or(os.Getenv("MONGO_REPLICA_SET"), "")
var MONGO_PASSWORD = cmp.Or(os.Getenv("MONGO_PASSWORD"), "")

var HEADLESS_SERVICE = cmp.Or(os.Getenv("HEADLESS_SERVICE"), "mongo")
var SIDECAR_TIME_TO_WAIT = cmp.Or(os.Getenv("SIDECAR_TIME_TO_WAIT"), "30")
var SIDECAR_TIME_SLEEP = cmp.Or(os.Getenv("SIDECAR_TIME_SLEEP"), "5")
var SIDECAR_SELECTOR_POD = cmp.Or(os.Getenv("SIDECAR_SELECTOR_POD"), "app=mongo")

var KUBE_NAMESPACE = cmp.Or(os.Getenv("KUBE_NAMESPACE"), "default")
var KUBE_CONFIG_TOKEN_PATH = cmp.Or(os.Getenv("KUBE_CONFIG_TOKEN_PATH"), "/var/run/secrets/kubernetes.io/serviceaccount/token")
