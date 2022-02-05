package service

import (
	"context"
	"fmt"

	"github.com/pkorobeinikov/platform/platform-lib/service/spec"
)

func Env(ctx context.Context, serviceEnv string) error {
	s, err := spec.Read()
	if err != nil {
		return err
	}

	environment := s.EnvironmentFor(serviceEnv)
	for k, v := range environment {
		fmt.Printf("export %s=%s\n", k, v)
	}

	return nil
}
