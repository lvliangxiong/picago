package main

import (
	"embed"

	"github.com/lvliangxiong/picago/router"
)

var (
	//go:embed static/*
	staticResources embed.FS
	//go:embed template
	templates embed.FS
)

func main() {
	router.Init(staticResources, templates)
}
