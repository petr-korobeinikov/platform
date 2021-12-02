package spec

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

const File = "service.yaml"

func (s *Spec) EnabledComponent() []string {
	l := make([]string, 0)

	for _, c := range s.Component {
		if c.Enabled {
			l = append(l, c.Type)
		}
	}

	return l
}

type (
	Spec struct {
		Name      string      `yaml:"name"`
		Component []Component `yaml:"component"`
	}

	Component struct {
		Type    string `yaml:"type"`
		Enabled bool   `yaml:"enabled"`
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

var (
	ErrSpecFileDoesNotExists = errors.New(fmt.Sprintf("%s does not found in project directory", File))
	ErrSpecReading           = errors.New(fmt.Sprintf("can't read %s", File))
	ErrSpecInvalid           = errors.New(fmt.Sprintf("%s contains errors", File))
)
