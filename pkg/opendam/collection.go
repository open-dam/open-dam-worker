package opendam

import (
	"context"
	"fmt"

	"gocloud.dev/docstore"
	_ "gocloud.dev/docstore/awsdynamodb"
	_ "gocloud.dev/docstore/gcpfirestore"
	_ "gocloud.dev/docstore/memdocstore"
	_ "gocloud.dev/docstore/mongodocstore"
)

func DocStoreFactory(collectionURL string) (*docstore.Collection, error) {
	coll, err := docstore.OpenCollection(context.Background(), collectionURL)
	if err != nil {
		return nil, fmt.Errorf("could not open collection: %v", err)
	}
	return coll, nil
}
