package service

import (
	"context"
	"os"

	"gopkg.in/yaml.v2"
)

func ReadSpec(ctx context.Context) (*Spec, error) {
	y, err := os.ReadFile(SpecFile)
	if err != nil {
		return nil, err
	}

	var spec Spec
	err = yaml.Unmarshal(y, &spec)
	if err != nil {
		return nil, err
	}

	return &spec, nil
}

type (
	Spec struct {
		Name string `yaml:"name"`
	}
)

const SpecFile = `service.yaml`
