package k8s

import (
	"context"
	"fmt"
	"strings"

	"github.com/mindmain/go-mongo-sidecar/types"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type MongoPod struct {
	Name string
	IP   string
}

func (m *MongoPod) String() string {
	if m == nil {
		return "(nil)"
	}

	return fmt.Sprintf("%s(%s)", m.Name, m.IP)
}

func (k *k8sHandler) GetPodsNamesWithMatchLabels(ctx context.Context, labels map[string]string) ([]*MongoPod, error) {

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

	var mongoPods []*MongoPod

	for _, pod := range pods.Items {

		if pod.Status.Phase != "Running" {
			continue
		}

		if pod.Namespace != types.KUBE_NAMESPACE {
			continue
		}

		mongoPods = append(mongoPods, &MongoPod{
			Name: pod.Name,
			IP:   pod.Status.PodIP,
		})
	}

	return mongoPods, nil
}
