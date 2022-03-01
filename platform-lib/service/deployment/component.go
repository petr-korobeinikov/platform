package deployment

import (
	"errors"
	"fmt"
	"strings"

	"github.com/pkorobeinikov/platform/platform-lib/service/env"
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
	case "minio":
		dcsList = append(dcsList, dockerComposeServiceV2{
			ContainerName: c.containerName(),
			Image:         "minio",
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
			// Префикс "service_" обозначает именно "сервисного" пользователя,
			// под которым выполняется приложение.
			// Для запуска миграций должен быть добавлен отдельный пользователь
			// с правами на создание и изменение структуры таблиц.
			// Суффикс "_rw" означает read/write — именно те права, которыми наделён сервисный пользователь.
			Environment: map[string]string{
				"POSTGRES_USER":     c.dockerComposeServiceEnvVarName("service_user_rw"),
				"POSTGRES_PASSWORD": c.dockerComposeServiceEnvVarName("service_password_rw"),
				"POSTGRES_DB":       c.dockerComposeServiceEnvVarName("database"),
			},
		})

		// Потребуется продумать систему суффиксов ролей:
		// - rw (read/write)  = operational
		// - fa (full access) = maintenance
		// - ro (read only)   = reader
		env.Registry().
			Register(c.componentEnvVarName("service_user_rw"), "service_rw").
			Register(c.componentEnvVarName("service_password_rw"), "postgres_secret").
			Register(c.componentEnvVarName("database"), "service")
	case "minio":
		dcsList = append(dcsList, dockerComposeServiceV2{
			ContainerName: c.containerName(),
			Image:         "quay.io/minio/minio:latest",
			Environment: map[string]string{
				"MINIO_ROOT_USER":     c.dockerComposeServiceEnvVarName("MINIO_ROOT_USER"),
				"MINIO_ROOT_PASSWORD": c.dockerComposeServiceEnvVarName("MINIO_ROOT_PASSWORD"),
			},
			Ports: []string{
				"9500:9500",
				"9501:9501",
			},
			Command: `server /data --address ":9500" --console-address ":9501"`,
		})

		env.Registry().
			Register(c.componentEnvVarName("MINIO_ROOT_USER"), "minio").
			Register(c.componentEnvVarName("MINIO_ROOT_PASSWORD"), "minio_secret")
	case "vault":
		dcsList = append(dcsList, dockerComposeServiceV2{
			ContainerName: c.containerName(),
			Image:         "vault:1.9.2",
			Environment: map[string]string{
				"VAULT_DEV_LISTEN_ADDRESS": c.dockerComposeServiceEnvVarName("VAULT_DEV_LISTEN_ADDRESS"),
				"VAULT_DEV_ROOT_TOKEN_ID":  c.dockerComposeServiceEnvVarName("VAULT_DEV_ROOT_TOKEN_ID"),
			},
			CapAdd: []string{"IPC_LOCK"},
		})

		env.Registry().
			Register(c.componentEnvVarName("VAULT_DEV_LISTEN_ADDRESS"), "0.0.0.0:8200").
			Register(c.componentEnvVarName("VAULT_DEV_ROOT_TOKEN_ID"), "vault_secret")
	}

	return
}

func (c *ServiceComponent) componentEnvVarName(s string) string {
	return strings.ToUpper(
		strings.ReplaceAll(
			fmt.Sprintf(
				"%s_%s",
				c.containerName(),
				s,
			),
			"-",
			"_",
		),
	)
}

func (c *ServiceComponent) dockerComposeServiceEnvVarName(s string) string {
	return fmt.Sprintf(
		"${%s}",
		strings.ToUpper(
			strings.ReplaceAll(
				fmt.Sprintf(
					"%s_%s",
					c.containerName(),
					s,
				),
				"-",
				"_",
			),
		),
	)
}

var (
	ErrUnsupportedServiceComponentType  = errors.New("unsupported service component type")
	ErrUnsupportedPlatformComponentType = errors.New("unsupported platform component type")
)
