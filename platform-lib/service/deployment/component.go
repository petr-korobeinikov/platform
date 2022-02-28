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

func (c *PlatformComponent) dockerComposeServiceSpecList() (dcsList []dockerComposeServiceV2, err error) {
	switch c.Type {
	default:
		return nil, ErrUnsupportedPlatformComponentType
	case "kafka":
		dcsList = append(
			dcsList,
			dockerComposeServiceV2{
				ContainerName: c.containerName() + "-broker",
				Image:         "kafka-broker",
			},
			dockerComposeServiceV2{
				ContainerName: c.containerName() + "-zookeeper",
				Image:         "kafka-zookeeper",
			},
			dockerComposeServiceV2{
				ContainerName: c.containerName() + "-kafdrop",
				Image:         "kafdrop",
			},
		)
	case "opentracing":
		dcsList = append(dcsList, dockerComposeServiceV2{
			ContainerName: c.containerName(),
			Image:         "opentracing",
		})
	}

	return
}

type ServiceComponent struct {
	Name string
	Type string
}

func (c *ServiceComponent) containerName() string {
	return fmt.Sprintf("service-component-%s-%s", c.Type, c.Name)
}

func (c *ServiceComponent) dockerComposeServiceSpecList() (dcsList []dockerComposeServiceV2, err error) {
	switch c.Type {
	default:
		return nil, ErrUnsupportedServiceComponentType
	case "postgres":
		dcsList = append(dcsList, dockerComposeServiceV2{
			ContainerName: c.containerName(),
			Image:         "postgres:13",
		})
	case "minio":
		dcsList = append(dcsList, dockerComposeServiceV2{
			ContainerName: c.containerName(),
			Image:         "quay.io/minio/minio:latest",
		})
	}

	return
}

var (
	ErrUnsupportedServiceComponentType  = errors.New("unsupported service component type")
	ErrUnsupportedPlatformComponentType = errors.New("unsupported platform component type")
)
