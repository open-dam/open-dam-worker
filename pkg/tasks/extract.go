package tasks

// The Extract task will take the media from its uploaded location, and the content type from its original source,
// and extract metadata from the file. Determining the assets kind and any additional meta data from the file.
func (t Tasker) Extract(loc, contentType string) (string, map[string]interface{}, error) {
	t.logger.Infof("extract %s", loc)

	_, _ = t.server.SendGroup(ProcessImage(), 0)
	return "ccc", nil, nil
}
