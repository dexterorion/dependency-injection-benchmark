package shared

import (
	"context"

	"cloud.google.com/go/datastore"
)

const (
	DIBenchmarkProjectId = "di-benchmark"
)

func GetDatastoreClient(ctx context.Context, projectId string) (*datastore.Client, error) {
	client, err := datastore.NewClient(ctx, projectId)
	if err != nil {
		return nil, err
	}

	return client, nil
}

type TestModel struct {
	Name string `datastore:"name"`
}
