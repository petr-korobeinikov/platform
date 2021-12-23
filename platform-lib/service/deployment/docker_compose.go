package deployment

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/pkorobeinikov/platform/platform-lib/service/env"
	"github.com/pkorobeinikov/platform/platform-lib/service/spec"
)

const DockerComposeFile = ".platform/docker-compose/docker-compose.yaml"

func (g *DockerComposeGenerator) Generate(s *spec.Spec) ([]byte, error) {
	var dcs dockerComposeSpec

	dcs.Version = "3"
	dcs.Services = make(map[string]dockerComposeService)

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

	dcs.Services["platform_kafka_zookeeper"] = dockerComposeService{
		ContainerName: "kafka-zookeeper",
		Image:         "confluentinc/cp-zookeeper:5.5.1",
		Restart:       "always",
		Environment: map[string]string{
			"ZOOKEEPER_CLIENT_PORT": "2181",
			"ZOOKEEPER_TICK_TIME":   "2000",
			"ALLOW_ANONYMOUS_LOGIN": "yes",
		},
	}

	dcs.Services["platform_kafka_broker"] = dockerComposeService{
		ContainerName: "kafka-broker",
		Image:         "confluentinc/cp-kafka:5.5.1",
		DependsOn: []string{
			"platform_kafka_zookeeper",
		},
		Restart: "on-failure",
		Ports: []string{
			"9092:9092",
		},
		Environment: map[string]string{
			"KAFKA_REST_HOST_NAME":                           "broker",
			"KAFKA_BROKER_ID":                                "1",
			"KAFKA_ZOOKEEPER_CONNECT":                        "kafka-zookeeper:2181",
			"KAFKA_ADVERTISED_LISTENERS":                     "PLAINTEXT://kafka-broker:29092,PLAINTEXT_HOST://localhost:9092",
			"KAFKA_LISTENER_SECURITY_PROTOCOL_MAP":           "PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT",
			"KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR":         "1",
			"KAFKA_TRANSACTION_STATE_LOG_MIN_ISR":            "1",
			"KAFKA_TRANSACTION_STATE_LOG_REPLICATION_FACTOR": "1",
			"KAFKA_GROUP_INITIAL_REBALANCE_DELAY_MS":         "0",
			"KAFKA_JMX_PORT":                                 "9101",
			"KAFKA_AUTO_CREATE_TOPICS_ENABLE":                "true",
			"KAFKA_LOG4J_LOGGERS":                            "org.apache.zookeeper=ERROR,org.apache.kafka=ERROR,kafka=ERROR,kafka.cluster=ERROR,kafka.controller=ERROR,kafka.coordinator=ERROR,kafka.log=ERROR,kafka.server=ERROR,kafka.zookeeper=ERROR,state.change.logger=ERROR",
		},
	}

	env.Registry().
		Register("KAFKA_PORT", "9092")

	dcs.Services["platform_kafka_kafdrop"] = dockerComposeService{
		ContainerName: "kafka-kafdrop",
		Image:         "obsidiandynamics/kafdrop",
		DependsOn: []string{
			"platform_kafka_broker",
		},
		Restart: "always",
		Ports: []string{
			"9100:9100",
		},
		Environment: map[string]string{
			"KAFKA_BROKERCONNECT": "kafka-broker:29092",
			"SERVER_PORT":         "9100",
		},
	}

	dcs.Services["platform_observability_opentelemetry"] = dockerComposeService{
		ContainerName: "opentelemetry",
		Image:         "jaegertracing/opentelemetry-all-in-one",
		Restart:       "always",
		Ports: []string{
			"6831:6831",
			"16686:16686",
			"14268:14268",
		},
	}

	env.Registry().
		Register("OBSERVABILITY_JAEGER_COLLECTOR_HTTP_ENDPOINT", "http://localhost:14268/api/traces")

	env.Registry().
		Register("SERVICE", s.Name).
		Register("SERVICE_IMAGE_NAME", s.Name).
		Register("SERVICE_IMAGE_TAG", "latest")

	serviceEnvironment := make(map[string]string)
	for k := range env.Registry().All() {
		serviceEnvironment[k] = fmt.Sprintf("${%s}", k)
	}

	dcs.Services["service"] = dockerComposeService{
		ContainerName: "service",
		Image:         "${SERVICE_IMAGE_NAME}:${SERVICE_IMAGE_TAG}",
		Restart:       "always",
		Ports:         []string{"9000:9000"},
		Environment:   serviceEnvironment,
	}

	// Hide under feature flag?
	dcs.Services["platform_service_desktop"] = dockerComposeService{
		ContainerName: "platform_service_desktop",
		Image:         "platform-service-desktop",
		Restart:       "always",
		Ports:         []string{"80:80"},
		Environment: map[string]string{
			"LISTEN_ON":                  ":80",
			"COMPONENT_JAEGERUI_ENABLED": "true",
			"COMPONENT_KAFDROP_ENABLED":  "true",
			"COMPONENT_JAEGERUI_HOST":    "http://localhost:16686",
			"COMPONENT_KAFDROP_HOST":     "http://localhost:9100",
			"JAEGER_SERVICE_NAME":        "platform-service-desktop",
			"JAEGER_ENDPOINT":            "http://localhost:14268/api/traces",
			"JAEGER_SAMPLER_TYPE":        "const",
			"JAEGER_SAMPLER_PARAM":       "0",
		},
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

func WriteDockerComposeFile(deploymentSpec []byte) error {
	return os.WriteFile(DockerComposeFile, deploymentSpec, 0644)
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
		DependsOn     []string          `yaml:"depends_on,omitempty"`
		Ports         []string          `yaml:"ports,omitempty"`
		Environment   map[string]string `yaml:"environment,omitempty"`
	}
)

var (
	ErrUnsupportedComponent = errors.New("unsupported component")
)
