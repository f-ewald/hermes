package templates

import (
	"embed"
	_ "embed"
)

// Templates is a file system that contains all embedded templates for the output.
//go:embed *.tpl
var Templates embed.FS
