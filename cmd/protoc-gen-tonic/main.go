package main

import (
	"os"

	"github.com/wasilibs/go-protoc-gen-tonic/internal/runner"
	"github.com/wasilibs/go-protoc-gen-tonic/internal/wasm"
)

func main() {
	os.Exit(runner.Run("protoc-gen-tonic", os.Args[1:], wasm.ProtocGenTonic, os.Stdin, os.Stdout, os.Stderr))
}
