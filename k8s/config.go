package k8s

import (
	"os"

	"github.com/mindmain/go-mongo-sidecar/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func KubeClient() (*kubernetes.Clientset, error) {
	tokenFile, err := os.ReadFile(types.KUBE_CONFIG_TOKEN_PATH.Get())

	if err != nil {
		return nil, err
	}

	client, err := kubernetes.NewForConfig(&rest.Config{
		BearerToken: string(tokenFile),
		Host:        "https://kubernetes.default.svc",
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
		},
	})

	if err != nil {
		return nil, err
	}

	return client, nil
}
