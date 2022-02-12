package spec

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

const File = "platform.yaml"

func (s *Spec) EnabledComponent() []string {
	l := make([]string, 0)

	for _, c := range s.Component {
		if c.Enabled {
			l = append(l, c.ID())
		}
	}

	l = append(
		l,
		"platform-observability-opentelemetry",
		"platform-kafka-zookeeper",
		"platform-kafka-broker",
		"platform-kafka-kafdrop",
		fmt.Sprintf("platform-sentinel-%s", s.Name),
	)

	return l
}

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

func (c *Component) ID() string {
	return fmt.Sprintf("component-%s-%s", c.Type, c.Name)
}

func (c *Component) FormatEnvVarName(v string) string {
	up := strings.ToUpper(c.ID())
	uv := strings.ToUpper(v)

	return strings.ReplaceAll(
		fmt.Sprintf("%s_%s", up, uv),
		"-",
		"_",
	)
}

func (c *Component) FormatEnvVarNameEscaped(v string) string {
	return fmt.Sprintf("${%s}", c.FormatEnvVarName(v))
}

type (
	Spec struct {
		Name        string                       `yaml:"name"`
		Component   []*Component                 `yaml:"component"`
		Environment map[string]map[string]string `yaml:"environment"`
		Task        []*Task                      `yaml:"task"`
	}

	Component struct {
		Type    string `yaml:"type"`
		Name    string `yaml:"name"`
		Enabled bool   `yaml:"enabled"`
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

var (
	ErrSpecFileDoesNotExists = errors.New(fmt.Sprintf("%s does not found in project directory", File))
	ErrSpecReading           = errors.New(fmt.Sprintf("can't read %s", File))
	ErrSpecInvalid           = errors.New(fmt.Sprintf("%s contains errors", File))
	ErrTaskNotDefined        = errors.New(fmt.Sprintf("task not defined in file %s", File))
)
