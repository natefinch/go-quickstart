package main

import (
	"os"

	"github.com/natefinch/go-quickstart/server/cmd/my-server/run"
)

// The meat of how to run the server goes in the run package.
func main() {
	os.Exit(run.Run())
}
