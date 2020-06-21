package tasks

import (
	"context"

	opendamclient "github.com/open-dam/open-dam-go-client"
	"github.com/sirupsen/logrus"
)

// The Extract task will take the media from its uploaded location, and the content type from its original source,
// and extract metadata from the file. Determining the assets kind and any additional meta data from the file.
func (t Tasker) Extract(url, assetID string) error {
	p, err := t.DownloadFile(url)
	if err != nil {
		return err
	}

	m := t.exif.ExtractMetadata(p)
	if len(m) == 0 {
		t.logger.WithFields(logrus.Fields{
			"url":      url,
			"asset_id": assetID,
		}).Warn("no metadata extracted from file")
	}

	job, resp, err := t.api.PutAsset(context.Background(), assetID, opendamclient.AssetUpdate{
		Metadata: map[string]map[string]interface{}{
			"exiftool": m[0].Fields,
		},
	})
	if err != nil {
		return err
	}
	t.logger.WithFields(logrus.Fields{
		"job":  job,
		"resp": resp,
	}).Debug("asset put response")
	return nil
}
