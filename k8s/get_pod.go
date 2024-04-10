package k8s

import (
	"context"
	"strings"

	"github.com/mindmain/go-mongo-sidecar/types"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *k8sHandler) GetPodsNamesWithMatchLabels(ctx context.Context, labels map[string]string) ([]string, error) {

	var s []string

	for k, v := range labels {
		s = append(s, k+"="+v)
	}

	pods, err := k.client.CoreV1().Pods("").List(ctx, v1.ListOptions{
		LabelSelector: strings.Join(s, ","),
	})

	if err != nil {
		return nil, err
	}

	var names []string

	for _, pod := range pods.Items {

		if pod.Status.Phase != "Running" {
			continue
		}

		if pod.Namespace != types.KUBE_NAMESPACE.Get() {
			continue
		}

		names = append(names, pod.Name)
	}

	return names, nil
}
