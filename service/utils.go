package service

import (
	"fmt"
	"strings"
)

func stringLabelToMap(label string) (map[string]string, error) {

	couples := strings.Split(label, ",")

	result := make(map[string]string)

	for _, couple := range couples {

		keyValue := strings.Split(couple, "=")

		if len(keyValue) != 2 {
			return nil, fmt.Errorf("invalid label format")
		}

		result[keyValue[0]] = keyValue[1]
	}

	return result, nil

}

func addServiceToPodsNames(podsNames []string, serviceName string) []string {

	result := make([]string, len(podsNames))

	for i, podName := range podsNames {
		result[i] = fmt.Sprintf("%s.%s", podName, serviceName)
	}

	return result
}
