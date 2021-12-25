package service

import (
	"context"
	"errors"
	"fmt"
	"path"

	"github.com/pkorobeinikov/platform/platform-lib/filesystem"
	"github.com/pkorobeinikov/platform/platform-lib/service/spec"
	"github.com/pkorobeinikov/platform/platform-lib/service/validation"
)

func Create(ctx context.Context, serviceName string) (err error) {
	if err = validation.EnsureServiceNameValid(serviceName); err != nil {
		return
	}

	if filesystem.IsDirectoryExists(serviceName) {
		return errServiceDirectoryAlreadyExists
	}

	if err = filesystem.MkDir(serviceName); err != nil {
		return
	}

	// region Service Template
	if err = filesystem.MkDir(path.Join(serviceName, "cmd", "service")); err != nil {
		return
	}

	err = filesystem.Touch(
		path.Join(serviceName, "go.mod"),
		filesystem.WithContentsOfString(
			fmt.Sprintf(gomodFmt, serviceName),
		),
	)
	if err != nil {
		return
	}

	err = filesystem.Touch(
		path.Join(serviceName, spec.File),
		filesystem.WithContentsOfString(
			fmt.Sprintf(specFmt, serviceName),
		),
	)
	if err != nil {
		return
	}

	err = filesystem.Touch(
		path.Join(serviceName, "cmd", "service", "main.go"),
		filesystem.WithContentsOfString(
			fmt.Sprintf(entrypoint, serviceName),
		),
	)
	if err != nil {
		return
	}

	err = filesystem.Touch(
		path.Join(serviceName, ".gitignore"),
		filesystem.WithContentsOfString(gitignore),
	)
	if err != nil {
		return
	}
	// endregion

	return
}

const (
	gomodFmt = `module %s

go 1.17
`

	specFmt = `name: %s

environment:
  _: ~
  local: ~
  staging: ~
`

	entrypoint = `package main

func main() {
	println("Hello, I am %s!")
}
`

	gitignore = `.platform
`
)

var (
	errServiceDirectoryAlreadyExists = errors.New("service directory already exists")
)
