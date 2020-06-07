package tasks

import "github.com/RichardKnop/machinery/v1/tasks"

func ProcessAsset(URL string) *tasks.Chain {
	upload, _ := tasks.NewSignature("upload", []tasks.Arg{{Type: "string", Value: URL}})
	extract, _ := tasks.NewSignature("extract", nil)
	chain, _ := tasks.NewChain(upload, extract)

	return chain
}

func ProcessImage() *tasks.Group {
	analysis, _ := tasks.NewSignature("imageanalysis", []tasks.Arg{{Type: "string", Value: "dddd"}})
	thumbnail, _ := tasks.NewSignature("imagecreation", []tasks.Arg{{Type: "string", Value: "dddd"}})
	group, _ := tasks.NewGroup(analysis, thumbnail)

	return group
}

func ProcessAudio() *tasks.Signature {
	soundwave, _ := tasks.NewSignature("soundwave", []tasks.Arg{{Type: "string", Value: "dddd"}})

	return soundwave
}
