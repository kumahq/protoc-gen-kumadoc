package main

import (
	pgs "github.com/lyft/protoc-gen-star"
)

func main() {
	initOptions := []pgs.InitOption{
		pgs.DebugEnv("DEBUG"),
	}

	module := New()

	pgs.Init(initOptions...).
		RegisterModule(module).
		Render()
}
