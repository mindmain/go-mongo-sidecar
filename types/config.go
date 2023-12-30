package types

const MONGO_PORT EvirontmentVariable = "MONGO_PORT"
const MONGO_HOST EvirontmentVariable = "MONGO_HOST"
const MONGO_USER EvirontmentVariable = "MONGO_USER"
const MONGO_PASSWORD EvirontmentVariable = "MONGO_PASSWORD"
const MONGO_REPLICA_SET EvirontmentVariable = "MONGO_REPLICA_SET"

const HEADLESS_SERVICE EvirontmentVariable = "HEADLESS_SERVICE"

const SIDECAR_TIME_TO_WAIT EvirontmentVariable = "SIDECAR_TIME_TO_WAIT"
const SIDECAR_TIME_SLEEP EvirontmentVariable = "SIDECAR_TIME_SLEEP"
const SIDECAR_SELECTOR_POD EvirontmentVariable = "SIDECAR_SELECTOR_POD"

const KUBE_NAMESPACE EvirontmentVariable = "KUBE_NAMESPACE"
const KUBE_CONFIG_TOKEN_PATH EvirontmentVariable = "KUBE_CONFIG_TOKEN_PATH"

var defaultValues = map[EvirontmentVariable]string{
	MONGO_PORT:             "27017",
	MONGO_HOST:             "localhost",
	HEADLESS_SERVICE:       "mongo",
	SIDECAR_TIME_TO_WAIT:   "30",
	SIDECAR_TIME_SLEEP:     "5",
	KUBE_NAMESPACE:         "default",
	KUBE_CONFIG_TOKEN_PATH: "/var/run/secrets/kubernetes.io/serviceaccount/token",
	MONGO_PASSWORD:         "",
	MONGO_USER:             "",
	SIDECAR_SELECTOR_POD:   "app=mongo",
	MONGO_REPLICA_SET:      "",
}
