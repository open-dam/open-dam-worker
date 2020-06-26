package tasks

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/RichardKnop/machinery/v1"
	"github.com/barasher/go-exiftool"
	opendamclient "github.com/open-dam/open-dam-go-client"
	"github.com/open-dam/open-dam-worker/pkg/opendam"
	"github.com/sirupsen/logrus"
	"gocloud.dev/blob"
)

type Tasker struct {
	server *machinery.Server
	logger *logrus.Entry
	bucket *blob.Bucket
	api    *opendamclient.DefaultApiService
	exif   *exiftool.Exiftool
	dir    string
}

func NewTasker(cfg opendam.Config, server *machinery.Server, log *logrus.Entry) (*Tasker, error) {
	bucket, err := blob.OpenBucket(context.Background(), cfg.BlobConnection)
	if err != nil {
		return nil, fmt.Errorf("could not open bucket: %v", err)
	}

	client := opendamclient.NewAPIClient(&opendamclient.Configuration{
		Host:   cfg.OpenDAMHost,
		Scheme: "http",
	})

	exif, err := exiftool.NewExiftool()
	if err != nil {
		return nil, fmt.Errorf("could not start exiftool: %v", err)
	}

	return &Tasker{
		server: server,
		logger: log,
		bucket: bucket,
		api:    client.DefaultApi,
		exif:   exif,
		dir:    cfg.WorkingDir,
	}, nil
}

func (t *Tasker) Download(url string) ([]byte, error) {
	t.logger.WithField("url", url).Debug("downloading file from bucket")
	return t.bucket.ReadAll(context.Background(), url)
}

func (t *Tasker) DownloadFile(url string) (string, error) {
	b, err := t.Download(url)
	if err != nil {
		return "", err
	}

	path := filepath.Join(t.dir, url)
	err = ioutil.WriteFile(path, b, os.ModePerm)
	if err != nil {
		return "", err
	}
	return path, nil
}

func (t *Tasker) Close() {
	_ = t.exif.Close()
}
