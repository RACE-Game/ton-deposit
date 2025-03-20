package assets

import "embed"

//go:embed *.jpeg *.gif
var Assets embed.FS
