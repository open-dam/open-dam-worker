package tasks

import (
	"context"
	"fmt"

	"github.com/RichardKnop/machinery/v1"
	"github.com/sirupsen/logrus"
	"gocloud.dev/blob"
)

type Tasker struct {
	server *machinery.Server
	logger *logrus.Entry
	bucket *blob.Bucket
}

func NewTasker(server *machinery.Server, log *logrus.Entry) (*Tasker, error) {
	bucket, err := blob.OpenBucket(context.Background(), "file:///containerfiles")
	if err != nil {
		return nil, fmt.Errorf("could not open bucket: %v", err)
	}
	return &Tasker{
		server: server,
		logger: log,
		bucket: bucket,
	}, nil
}
