package opendam

// Asset - An asset is a single managed digital asset
type Asset struct {

	// The unique id of the asset
	AssetID string `json:"asset_id"`

	// The kind of media
	Kind string `json:"kind"`

	Version Version `json:"version"`

	File File `json:"file"`

	// additional assets/files associated with the asset
	Formats []Asset `json:"formats,omitempty"`

	// A list of metadata tags associated with the asset
	Tags []string `json:"tags,omitempty"`

	// Any user supplied metadata for the asset
	Metadata map[string]map[string]interface{} `json:"metadata,omitempty"`
}

type Assets struct {
	Assets []Asset `json:"assets,omitempty"`
}

// File - The file associated with an asset or format
type File struct {

	// the name of the file
	Name string `json:"name"`

	// A URL address to the file content
	Source string `json:"source"`

	// the http Content-Type used with the file
	ContentType string `json:"content_type"`

	// The size of the file in bytes
	Size float32 `json:"size"`

	// The width of the file
	Width float32 `json:"width,omitempty"`

	// The height of the file
	Height float32 `json:"height,omitempty"`

	// The duration of the file in seconds
	Duration float32 `json:"duration,omitempty"`
}

// Version - version data of an asset
type Version struct {

	// The version number
	Number float32 `json:"number"`

	// A point in time represented as milliseconds from the Epoch (UTC)
	Timestamp float32 `json:"timestamp"`

	JobID string `json:"job_id"`
}
