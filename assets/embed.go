package assets

import "embed"

//go:embed *.jpeg *.jpg *.gif *.mp4
var Assets embed.FS
