package opendam

// Job - The state of a single job
type Job struct {

	// The unique id of the job
	JobID string `json:"job_id"`

	// The state of the job
	State string `json:"state"`

	// The asset that this job is processing work for
	AssetID string `json:"asset_id"`
}
