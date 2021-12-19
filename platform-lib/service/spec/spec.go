package spec

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

const File = "service.yaml"

func (s *Spec) EnabledComponent() []string {
	l := make([]string, 0)

	for _, c := range s.Component {
		if c.Enabled {
			l = append(l, c.ID())
		}
	}

	l = append(
		l,
		"platform_observability_opentelemetry",
		"platform_kafka_zookeeper",
		"platform_kafka_broker",
		"platform_kafka_kafdrop",
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

func (c *Component) ID() string {
	return fmt.Sprintf("component_%s_%s", c.Type, c.Name)
}

func (c *Component) FormatEnvVarName(v string) string {
	up := strings.ToUpper(c.ID())
	uv := strings.ToUpper(v)

	return fmt.Sprintf("%s_%s", up, uv)
}

func (c *Component) FormatEnvVarNameEscaped(v string) string {
	return fmt.Sprintf("${%s}", c.FormatEnvVarName(v))
}

type (
	Spec struct {
		Name        string                       `yaml:"name"`
		Component   []*Component                 `yaml:"component"`
		Environment map[string]map[string]string `yaml:"environment"`
	}

	Component struct {
		Type    string `yaml:"type"`
		Name    string `yaml:"name"`
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

func ExistInCurrentDirectory() bool {
	_, err := os.Stat(File)

	return err == nil
}

var (
	ErrSpecFileDoesNotExists = errors.New(fmt.Sprintf("%s does not found in project directory", File))
	ErrSpecReading           = errors.New(fmt.Sprintf("can't read %s", File))
	ErrSpecInvalid           = errors.New(fmt.Sprintf("%s contains errors", File))
)
