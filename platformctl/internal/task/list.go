package task

import (
	"context"
	"errors"
	"fmt"

	"github.com/pkorobeinikov/platform/platform-lib/service/spec"
)

func List(ctx context.Context) error {
	s, err := spec.Read()
	if err != nil {
		return err
	}

	if len(s.Task) == 0 {
		return ErrNoTaskDefinitionFound
	}

	fmt.Println("List of available tasks:")
	for _, task := range s.Task {
		fmt.Println("\t", task.Name)
	}

	return nil
}

var (
	ErrNoTaskDefinitionFound = errors.New("no task definition found in service manifest")
)
