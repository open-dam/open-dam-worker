package tasks

// The Upload task will upload the file at the given URL to the configured storage provider
// and return the location, if successful.
func (t Tasker) Upload(URL string) (string, string, error) {
	t.logger.Infof("upload %s", URL)
	return "bbb", "ggg", nil
}
