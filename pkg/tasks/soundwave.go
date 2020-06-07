package tasks

import "fmt"

// The Soundwave task will create an image soundwave file for an audio asset, upload it, and attach it as a format to the asset
func (t Tasker) Soundwave(loc string) error {
	fmt.Printf("soundwave %s\n", loc)
	return nil
}
