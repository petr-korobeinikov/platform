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
		var (
			brokerName    = c.containerName() + "-broker"
			zookeeperName = c.containerName() + "-zookeeper"
			kafdropName   = c.containerName() + "-kafdrop"
		)

		dcsList = append(
			dcsList,
			dockerComposeServiceV2{
				ContainerName: brokerName,
				Image:         "confluentinc/cp-kafka:5.5.1",
				Restart:       "always",
				DependsOn: []string{
					zookeeperName,
				},
				Environment: map[string]string{
					"KAFKA_REST_HOST_NAME":                           brokerName,
					"KAFKA_BROKER_ID":                                "1",
					"KAFKA_ZOOKEEPER_CONNECT":                        zookeeperName + ":2181",
					"KAFKA_ADVERTISED_LISTENERS":                     "PLAINTEXT://" + brokerName + ":29092,PLAINTEXT_HOST://localhost:9092",
					"KAFKA_LISTENER_SECURITY_PROTOCOL_MAP":           "PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT",
					"KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR":         "1",
					"KAFKA_TRANSACTION_STATE_LOG_MIN_ISR":            "1",
					"KAFKA_TRANSACTION_STATE_LOG_REPLICATION_FACTOR": "1",
					"KAFKA_GROUP_INITIAL_REBALANCE_DELAY_MS":         "0",
					"KAFKA_JMX_PORT":                                 "9101",
					"KAFKA_AUTO_CREATE_TOPICS_ENABLE":                "true",
					"KAFKA_LOG4J_LOGGERS":                            "org.apache.zookeeper=ERROR,org.apache.kafka=ERROR,kafka=ERROR,kafka.cluster=ERROR,kafka.controller=ERROR,kafka.coordinator=ERROR,kafka.log=ERROR,kafka.server=ERROR,kafka.zookeeper=ERROR,state.change.logger=ERROR",
				},
			},
			dockerComposeServiceV2{
				ContainerName: zookeeperName,
				Image:         "confluentinc/cp-zookeeper:5.5.1",
				Restart:       "always",
				Environment: map[string]string{
					"ZOOKEEPER_CLIENT_PORT": "2181",
					"ZOOKEEPER_TICK_TIME":   "2000",
					"ALLOW_ANONYMOUS_LOGIN": "yes",
				},
			},
			dockerComposeServiceV2{
				ContainerName: kafdropName,
				Image:         "obsidiandynamics/kafdrop",
				Restart:       "always",
				Ports: []string{
					"9100:9100",
				},
				DependsOn: []string{
					zookeeperName,
				},
				Environment: map[string]string{
					"KAFKA_BROKERCONNECT": brokerName + ":29092",
					"SERVER_PORT":         "9100",
				},
			},
		)
	case "opentracing":
		dcsList = append(dcsList, dockerComposeServiceV2{
			ContainerName: c.containerName(),
			Image:         "jaegertracing/opentelemetry-all-in-one",
			Restart:       "always",
		})
	case "minio":
		dcsList = append(dcsList, dockerComposeServiceV2{
			ContainerName: c.containerName(),
			Image:         "quay.io/minio/minio:latest",
			Restart:       "always",
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
			Restart:       "always",
			Ports: []string{
				"5432:5432",
			},
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
			// !!! Добавить ip/хост виртуальной машины
			//Register(c.componentEnvVarName("IP"), os.Getenv("DOCKER_HOST")).
			//Register(c.componentEnvVarName("HOST"), os.Getenv("DOCKER_HOST")).
			Register(c.componentEnvVarName("service_user_rw"), "service_rw").
			Register(c.componentEnvVarName("service_password_rw"), "postgres_secret").
			Register(c.componentEnvVarName("database"), "service")
	case "minio":
		dcsList = append(dcsList, dockerComposeServiceV2{
			ContainerName: c.containerName(),
			Image:         "quay.io/minio/minio:latest",
			Restart:       "always",
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
			Restart:       "always",
			Environment: map[string]string{
				"VAULT_DEV_LISTEN_ADDRESS": c.dockerComposeServiceEnvVarName("VAULT_DEV_LISTEN_ADDRESS"),
				"VAULT_DEV_ROOT_TOKEN_ID":  c.dockerComposeServiceEnvVarName("VAULT_DEV_ROOT_TOKEN_ID"),
			},
			Ports: []string{
				"8200:8200",
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
