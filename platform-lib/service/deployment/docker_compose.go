package deployment

import (
	"errors"

	"gopkg.in/yaml.v2"

	"github.com/pkorobeinikov/platform/platform-lib/service/env"
	"github.com/pkorobeinikov/platform/platform-lib/service/spec"
)

const DockerComposeFile = ".platform/docker-compose/docker-compose.yaml"

func (g *DockerComposeGenerator) Generate(s *spec.Spec) ([]byte, error) {
	var dcs dockerComposeSpec

	dcs.Version = "3"
	dcs.Services = make(map[string]dockerComposeService)

	dcs.Services["service"] = dockerComposeService{
		ContainerName: "service",
		Image:         "${SERVICE_IMAGE_NAME}:${SERVICE_IMAGE_TAG}",
		Restart:       "always",
		Ports:         []string{"9000:9000"},
		Environment: map[string]string{
			"SERVICE": "${SERVICE}",
		},
	}

	env.Registry().
		Register("SERVICE", s.Name).
		Register("SERVICE_IMAGE_NAME", s.Name).
		Register("SERVICE_IMAGE_TAG", "latest")

	for _, c := range s.Component {
		containerName, image, ports, environment, err := componentContainerSpec(s.Name, c)
		if err != nil {
			return nil, err
		}

		if c.Enabled {
			dcs.Services[c.ID()] = dockerComposeService{
				ContainerName: containerName,
				Image:         image,
				Restart:       "always",
				Ports:         ports,
				Environment:   environment,
			}
		}
	}

	b, err := yaml.Marshal(dcs)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func NewDockerComposeGenerator() *DockerComposeGenerator {
	return &DockerComposeGenerator{}
}

func componentContainerSpec(serviceName string, c *spec.Component) (containerName, image string, ports []string, environment map[string]string, err error) {
	switch c.Type {
	case "postgres":
		containerName = c.ID()
		image = "postgres:14"
		ports = []string{"5432:5432"}
		environment = map[string]string{
			"POSTGRES_USER":     c.FormatEnvVarNameEscaped("user"),
			"POSTGRES_PASSWORD": c.FormatEnvVarNameEscaped("password"),
			"POSTGRES_DB":       c.FormatEnvVarNameEscaped("db"),
		}

		env.Registry().
			Register(c.FormatEnvVarName("user"), serviceName).
			Register(c.FormatEnvVarName("password"), "secret").
			Register(c.FormatEnvVarName("db"), serviceName)
	default:
		err = ErrUnsupportedComponent
	}

	return containerName, image, ports, environment, err
}

type (
	DockerComposeGenerator struct{}

	dockerComposeSpec struct {
		Version  string                          `yaml:"version"`
		Services map[string]dockerComposeService `yaml:"services"`
	}

	dockerComposeService struct {
		ContainerName string            `yaml:"container_name"`
		Image         string            `yaml:"image"`
		Restart       string            `yaml:"restart"`
		Ports         []string          `yaml:"ports,omitempty"`
		Environment   map[string]string `yaml:"environment,omitempty"`
	}
)

var (
	ErrUnsupportedComponent = errors.New("unsupported component")
)
