package opendam

import (
	"context"
	"fmt"

	"gocloud.dev/blob"
	_ "gocloud.dev/blob/azureblob"
	_ "gocloud.dev/blob/fileblob"
	_ "gocloud.dev/blob/gcsblob"
	_ "gocloud.dev/blob/s3blob"
)

func BlobFactory(storageURL string) (*blob.Bucket, error) {
	bucket, err := blob.OpenBucket(context.Background(), storageURL)
	if err != nil {
		return nil, fmt.Errorf("could not open bucket: %v", err)
	}
	return bucket, nil
}
