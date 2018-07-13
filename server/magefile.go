//+build mage

package main

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/magefile/mage/sh"
)

// Runs "go build" for Who.  This generates the version info the binary.
func Build() error {
	ldf, err := flags()
	if err != nil {
		return err
	}

	name := "who"
	if runtime.GOOS == "windows" {
		name += ".exe"
	}

	return sh.RunV("go", "build", "-ldflags="+ldf, "github.com/natefinch/go-quickstart/server/cmd/my-server")
}

// Generates a new tag.  Expects the TAG environment variable to be set,
// which will create a new tag with that name.
func Tag() (err error) {
	if os.Getenv("TAG") == "" {
		return errors.New("TAG environment variable is required")
	}
	if err := sh.RunV("git", "tag", "-a", "$TAG"); err != nil {
		return err
	}
	return sh.RunV("git", "push", "origin", "$TAG")
}

func flags() (string, error) {
	timestamp := time.Now().Format(time.RFC3339)
	hash := hash()
	tag := tag()
	if tag == "" {
		tag = "dev"
	}
	return fmt.Sprintf(`-X "github.com/natefinch/go-quickstart/server/cmd/my-server/run.timestamp=%s" -X "github.com/natefinch/go-quickstart/server/cmd/my-server/run.commitHash=%s" -X "github.com/natefinch/go-quickstart/server/cmd/my-server/run.gitTag=%s"`, timestamp, hash, tag), nil
}

// tag returns the git tag for the current branch or "" if none.
func tag() string {
	s, _ := sh.Output("git", "describe", "--tags")
	return s
}

// hash returns the git hash for the current repo or "" if none.
func hash() string {
	hash, _ := sh.Output("git", "rev-parse", "--short", "HEAD")
	return hash
}
