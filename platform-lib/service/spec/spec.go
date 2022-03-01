package spec

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

const File = "platform.yaml"

type (
	Spec struct {
		Name        string                       `yaml:"name"`
		Component   []*Component                 `yaml:"component"`
		Environment map[string]map[string]string `yaml:"environment"`
		Task        []*Task                      `yaml:"task"`
	}

	Task struct {
		Name     string       `yaml:"name"`
		Image    string       `yaml:"image"`
		Argument TaskArgument `yaml:"argument"`
	}

	TaskArgument struct {
		Workdir string `yaml:"workdir"`
		Command string `yaml:"command"`
	}
)

func (s *Spec) EnvironmentFor(environmentName string) map[string]string {
	out := make(map[string]string)

	if global, ok := s.Environment["_"]; ok {
		for k, v := range global {
			out[k] = v
		}
	}

	if wanted, ok := s.Environment[environmentName]; ok {
		for k, v := range wanted {
			out[k] = v
		}
	}

	return out
}

func (s *Spec) ShellEnvironmentFor(environmentName string) []string {
	var out []string

	envmap := s.EnvironmentFor(environmentName)
	for k, v := range envmap {
		out = append(out, fmt.Sprintf("%s=%s", k, v))
	}

	return out
}

func (s *Spec) SetEnvironmentFor(environmentName string) error {
	for k, v := range s.EnvironmentFor(environmentName) {
		if err := os.Setenv(k, v); err != nil {
			return err
		}
	}
	return nil
}

func (s *Spec) TaskByName(name string) (*Task, error) {
	for _, t := range s.Task {
		if t.Name == name {
			return t, nil
		}
	}

	return nil, ErrTaskNotDefined
}

func Read() (*Spec, error) {
	if _, err := os.Stat(File); os.IsNotExist(err) {
		return nil, ErrSpecFileDoesNotExists
	}

	y, err := os.ReadFile(File)
	if err != nil {
		return nil, ErrSpecReading
	}

	var spec Spec
	if err := yaml.Unmarshal(y, &spec); err != nil {
		return nil, ErrSpecInvalid
	}

	return &spec, nil
}

func ExistInCurrentDirectory() bool {
	_, err := os.Stat(File)

	return err == nil
}
