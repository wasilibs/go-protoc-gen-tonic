package main

import (
	"github.com/goyek/x/boot"
	"github.com/wasilibs/tools/tasks"
)

func main() {
	tasks.Define(tasks.Params{
		LibraryName: "protoc-gen-tonic",
		LibraryRepo: "hyperium/tonic",
		GoReleaser:  true,
	})
	boot.Main()
}
