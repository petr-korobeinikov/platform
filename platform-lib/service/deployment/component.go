package deployment

import (
	"errors"
	"fmt"
)

type PlatformComponent struct {
	Name string
	Type string
}

func (c *PlatformComponent) containerName() string {
	return fmt.Sprintf("platform-component-%s-%s", c.Type, c.Name)
}

func (s *PlatformComponent) dockerComposeServiceSpecList() (dcsList []dockerComposeServiceV2, err error) {
	switch s.Type {
	default:
		return nil, ErrUnsupportedPlatformComponentType
	case "kafka":
		dcsList = append(
			dcsList,
			dockerComposeServiceV2{
				ContainerName: s.containerName() + "-broker",
				Image:         "kafka-broker",
			},
			dockerComposeServiceV2{
				ContainerName: s.containerName() + "-zookeeper",
				Image:         "kafka-zookeeper",
			},
			dockerComposeServiceV2{
				ContainerName: s.containerName() + "-kafdrop",
				Image:         "kafdrop",
			},
		)
	case "opentracing":
		dcsList = append(dcsList, dockerComposeServiceV2{
			ContainerName: s.containerName(),
			Image:         "opentracing",
		})
	}

	return
}

type ServiceComponent struct {
	Name string
	Type string
}

func (s *ServiceComponent) containerName() string {
	return fmt.Sprintf("service-component-%s-%s", s.Type, s.Name)
}

func (s *ServiceComponent) dockerComposeServiceSpecList() (dcsList []dockerComposeServiceV2, err error) {
	switch s.Type {
	default:
		return nil, ErrUnsupportedServiceComponentType
	case "postgres":
		dcsList = append(dcsList, dockerComposeServiceV2{
			ContainerName: s.containerName(),
			Image:         "postgres:13",
		})
	}

	return
}

var (
	ErrUnsupportedServiceComponentType  = errors.New("unsupported service component type")
	ErrUnsupportedPlatformComponentType = errors.New("unsupported platform component type")
)
