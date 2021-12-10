package deployment

import (
	"gopkg.in/yaml.v2"

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

	for _, c := range s.Component {
		if c.Enabled {
			dcs.Services[c.ID()] = dockerComposeService{
				ContainerName: c.ID(),
				Image:         "postgres:14",
				Restart:       "always",
				Ports:         []string{"5432:5432"},
				Environment: map[string]string{
					"POSTGRES_USER":     c.FormatEnvVarName("user"),
					"POSTGRES_PASSWORD": c.FormatEnvVarName("password"),
					"POSTGRES_DB":       c.FormatEnvVarName("db"),
				},
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
