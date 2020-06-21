package tasks

import "context"

// The Upload task will upload the file at the given URL to the configured storage provider
// and return the location, if successful.
func (t Tasker) Upload(URL string) (string, error) {
	err := t.bucket.WriteAll(context.Background(), "", nil, nil)

	return "ggg", err
}
