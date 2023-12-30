package k8s

import (
	"k8s.io/client-go/kubernetes"
)

func NewK8sHandler(client *kubernetes.Clientset) HandlerKubernetes {
	return &k8sHandler{
		client: client,
	}
}

type k8sHandler struct {
	client *kubernetes.Clientset
}
