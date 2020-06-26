package tasks

import (
	"bytes"
	"context"
	"image"
	"image/jpeg"
	_ "image/png"

	"github.com/google/uuid"
	"github.com/muesli/smartcrop"
	"github.com/muesli/smartcrop/nfnt"
	opendamclient "github.com/open-dam/open-dam-go-client"
	"github.com/sirupsen/logrus"
)

// The ImageCreation task will create a new image from the asset, upload it, and attach it as a format on the asset
func (t Tasker) ImageCreation(url, assetID string, w, h int) error {
	b, err := t.Download(url)
	if err != nil {
		return err
	}

	img, _, _ := image.Decode(bytes.NewReader(b))

	analyzer := smartcrop.NewAnalyzer(nfnt.NewDefaultResizer())
	topCrop, _ := analyzer.FindBestCrop(img, w, h)

	// The crop will have the requested aspect ratio, but you need to copy/scale it yourself
	t.logger.Infof("Top crop: %+v\n", topCrop)

	type SubImager interface {
		SubImage(r image.Rectangle) image.Image
	}
	croppedimg := img.(SubImager).SubImage(topCrop)

	wr, err := t.bucket.NewWriter(context.Background(), assetID+"/thumbnail.jpg", nil)
	if err != nil {
		return err
	}
	defer wr.Close()

	if err = jpeg.Encode(wr, croppedimg, &jpeg.Options{Quality: 85}); err != nil {
		return err
	}

	thumbnail := opendamclient.Asset{
		AssetId: uuid.New().String(),
		Kind:    "image",
		Version: opendamclient.Version{},
		File: opendamclient.AssetFile{
			Name:        "thumbnail.jpg",
			Source:      assetID + "/thumbnail.jpg",
			ContentType: "image/jpeg",
			Width:       float32(w),
			Height:      float32(h),
			Size:        0,
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
