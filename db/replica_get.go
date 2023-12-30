package db

import "context"

type replSetGetConfigResult struct {
	Config ReplicaSetConfig `bson:"config"`
	Ok     int              `bson:"ok"`
}

func (m *mongoHandler) getConfig(ctx context.Context) (*ReplicaSetConfig, error) {

	var result replSetGetConfigResult

	res := m.client.Database("admin").RunCommand(ctx, map[string]string{"replSetGetConfig": "1"}).Decode(&result)

	if res != nil {
		return nil, res
	}

	return &result.Config, nil

}
