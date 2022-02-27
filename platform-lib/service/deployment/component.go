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

type ServiceComponent struct {
	Name string
	Type string
}

func (s *ServiceComponent) containerName() string {
	return fmt.Sprintf("service-component-%s-%s", s.Type, s.Name)
}

func (s *ServiceComponent) dockerComposeServiceSpec() (dockerComposeServiceV2, error) {
	var dcs dockerComposeServiceV2

	switch s.Type {
	default:
		return dcs, ErrUnsupportedServiceComponentType
	case "postgres":
		dcs.Image = "postgres:13"
	}

	dcs.ContainerName = s.containerName()

	return dcs, nil
}

var (
	ErrUnsupportedServiceComponentType = errors.New("unsupported service component type")
)
