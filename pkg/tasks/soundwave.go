package tasks

import (
	"bytes"
	"context"
	"image/jpeg"

	"github.com/google/uuid"
	"github.com/mdlayher/waveform"
	opendamclient "github.com/open-dam/open-dam-go-client"
	"github.com/sirupsen/logrus"
)

// The Soundwave task will create an image soundwave file for an audio asset, upload it, and attach it as a format to the asset
func (t Tasker) Soundwave(url, assetID string) error {
	b, err := t.Download(url)
	if err != nil {
		return err
	}

	img, err := waveform.Generate(bytes.NewReader(b))
	if err != nil {
		return err
	}

	wr, err := t.bucket.NewWriter(context.Background(), assetID+"/soundwave.jpg", nil)
	if err != nil {
		return err
	}
	defer wr.Close()

	if err = jpeg.Encode(wr, img, &jpeg.Options{Quality: 85}); err != nil {
		return err
	}

	thumbnail := opendamclient.Asset{
		AssetId: uuid.New().String(),
		Kind:    "image",
		Version: opendamclient.Version{},
		File: opendamclient.AssetFile{
			Name:        "soundwave.jpg",
			Source:      assetID + "/soundwave.jpg",
			ContentType: "image/jpeg",
		},
	}
	job, resp, err := t.api.PutAsset(context.Background(), assetID, opendamclient.AssetUpdate{
		Formats: []opendamclient.Asset{thumbnail},
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
