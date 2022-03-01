package deployment_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/pkorobeinikov/platform/platform-lib/service/deployment"
	"github.com/pkorobeinikov/platform/platform-lib/service/env"
)

func TestDockerComposeGeneratorV2_Generate(t *testing.T) {
	t.Run(`empty`, func(t *testing.T) {
		defer env.Registry().Clear()

		expected := "services: {}\n"

		sut := NewDockerComposeGeneratorV2()

		actual, err := sut.Generate(SpecGenerationRequest{})

		assert.NoError(t, err)
		assert.Equal(t, expected, actual.FileList[DockerComposeFile])
		assert.Equal(t, "", actual.FileList[env.File])
	})

	t.Run(`multiple service component`, func(t *testing.T) {
		defer env.Registry().Clear()

		expected := `services:
  service-component-postgres-master:
    container_name: service-component-postgres-master
    image: postgres:13
    environment:
      POSTGRES_DB: ${SERVICE_COMPONENT_POSTGRES_MASTER_DATABASE}
      POSTGRES_PASSWORD: ${SERVICE_COMPONENT_POSTGRES_MASTER_SERVICE_PASSWORD_RW}
      POSTGRES_USER: ${SERVICE_COMPONENT_POSTGRES_MASTER_SERVICE_USER_RW}
  service-component-postgres-olap:
    container_name: service-component-postgres-olap
    image: postgres:13
    environment:
      POSTGRES_DB: ${SERVICE_COMPONENT_POSTGRES_OLAP_DATABASE}
      POSTGRES_PASSWORD: ${SERVICE_COMPONENT_POSTGRES_OLAP_SERVICE_PASSWORD_RW}
      POSTGRES_USER: ${SERVICE_COMPONENT_POSTGRES_OLAP_SERVICE_USER_RW}
`

		expectedEnv := `SERVICE_COMPONENT_POSTGRES_MASTER_DATABASE="service"
SERVICE_COMPONENT_POSTGRES_MASTER_SERVICE_PASSWORD_RW="postgres_secret"
SERVICE_COMPONENT_POSTGRES_MASTER_SERVICE_USER_RW="service_rw"
SERVICE_COMPONENT_POSTGRES_OLAP_DATABASE="service"
SERVICE_COMPONENT_POSTGRES_OLAP_SERVICE_PASSWORD_RW="postgres_secret"
SERVICE_COMPONENT_POSTGRES_OLAP_SERVICE_USER_RW="service_rw"`

		given := SpecGenerationRequest{
			ServiceName:      "wordcounter-svc",
			ServiceNamespace: "wordcounter-ns",
			ServiceComponentList: []*ServiceComponent{
				{
					Name: "master",
					Type: "postgres",
				},
				{
					Name: "olap",
					Type: "postgres",
				},
			},
			PlatformComponentList: nil,
		}

		sut := NewDockerComposeGeneratorV2()

		actual, err := sut.Generate(given)

		assert.NoError(t, err)
		assert.Equal(t, expected, actual.FileList[DockerComposeFile])
		assert.Equal(t, expectedEnv, actual.FileList[env.File])
	})

	t.Run(`service component + platform component`, func(t *testing.T) {
		defer env.Registry().Clear()

		expected := `services:
  platform-component-kafka-kafka-broker:
    container_name: platform-component-kafka-kafka-broker
    image: kafka-broker
  platform-component-kafka-kafka-kafdrop:
    container_name: platform-component-kafka-kafka-kafdrop
    image: kafdrop
  platform-component-kafka-kafka-zookeeper:
    container_name: platform-component-kafka-kafka-zookeeper
    image: kafka-zookeeper
  platform-component-opentracing-opentracing:
    container_name: platform-component-opentracing-opentracing
    image: opentracing
  service-component-postgres-master:
    container_name: service-component-postgres-master
    image: postgres:13
    environment:
      POSTGRES_DB: ${SERVICE_COMPONENT_POSTGRES_MASTER_DATABASE}
      POSTGRES_PASSWORD: ${SERVICE_COMPONENT_POSTGRES_MASTER_SERVICE_PASSWORD_RW}
      POSTGRES_USER: ${SERVICE_COMPONENT_POSTGRES_MASTER_SERVICE_USER_RW}
`

		expectedEnv := `SERVICE_COMPONENT_POSTGRES_MASTER_DATABASE="service"
SERVICE_COMPONENT_POSTGRES_MASTER_SERVICE_PASSWORD_RW="postgres_secret"
SERVICE_COMPONENT_POSTGRES_MASTER_SERVICE_USER_RW="service_rw"`

		given := SpecGenerationRequest{
			ServiceName:      "wordcounter-svc",
			ServiceNamespace: "wordcounter-ns",
			ServiceComponentList: []*ServiceComponent{
				{
					Name: "master",
					Type: "postgres",
				},
			},
			PlatformComponentList: []*PlatformComponent{
				{
					Name: "kafka",
					Type: "kafka",
				},
				{
					Name: "opentracing",
					Type: "opentracing",
				},
			},
		}

		sut := NewDockerComposeGeneratorV2()

		actual, err := sut.Generate(given)

		assert.NoError(t, err)
		assert.Equal(t, expected, actual.FileList[DockerComposeFile])
		assert.Equal(t, expectedEnv, actual.FileList[env.File])
	})

	t.Run(`full service component`, func(t *testing.T) {
		defer env.Registry().Clear()

		expected := `services:
  service-component-minio-minio:
    container_name: service-component-minio-minio
    image: quay.io/minio/minio:latest
    environment:
      MINIO_ROOT_PASSWORD: ${SERVICE_COMPONENT_MINIO_MINIO_MINIO_ROOT_PASSWORD}
      MINIO_ROOT_USER: ${SERVICE_COMPONENT_MINIO_MINIO_MINIO_ROOT_USER}
  service-component-postgres-master:
    container_name: service-component-postgres-master
    image: postgres:13
    environment:
      POSTGRES_DB: ${SERVICE_COMPONENT_POSTGRES_MASTER_DATABASE}
      POSTGRES_PASSWORD: ${SERVICE_COMPONENT_POSTGRES_MASTER_SERVICE_PASSWORD_RW}
      POSTGRES_USER: ${SERVICE_COMPONENT_POSTGRES_MASTER_SERVICE_USER_RW}
  service-component-vault-vault:
    container_name: service-component-vault-vault
    image: vault:1.9.2
    environment:
      VAULT_DEV_LISTEN_ADDRESS: ${SERVICE_COMPONENT_VAULT_VAULT_VAULT_DEV_LISTEN_ADDRESS}
      VAULT_DEV_ROOT_TOKEN_ID: ${SERVICE_COMPONENT_VAULT_VAULT_VAULT_DEV_ROOT_TOKEN_ID}
    cap_add:
    - IPC_LOCK
`

		expectedEnv := `SERVICE_COMPONENT_MINIO_MINIO_MINIO_ROOT_PASSWORD="minio_secret"
SERVICE_COMPONENT_MINIO_MINIO_MINIO_ROOT_USER="minio"
SERVICE_COMPONENT_POSTGRES_MASTER_DATABASE="service"
SERVICE_COMPONENT_POSTGRES_MASTER_SERVICE_PASSWORD_RW="postgres_secret"
SERVICE_COMPONENT_POSTGRES_MASTER_SERVICE_USER_RW="service_rw"
SERVICE_COMPONENT_VAULT_VAULT_VAULT_DEV_LISTEN_ADDRESS="0.0.0.0:8200"
SERVICE_COMPONENT_VAULT_VAULT_VAULT_DEV_ROOT_TOKEN_ID="vault_secret"`

		given := SpecGenerationRequest{
			ServiceName:      "wordcounter-svc",
			ServiceNamespace: "wordcounter-ns",
			ServiceComponentList: []*ServiceComponent{
				{
					Name: "master",
					Type: "postgres",
				},
				{
					Name: "minio",
					Type: "minio",
				},
				{
					Name: "vault",
					Type: "vault",
				},
			},
			PlatformComponentList: nil,
		}

		sut := NewDockerComposeGeneratorV2()

		actual, err := sut.Generate(given)

		assert.NoError(t, err)
		assert.Equal(t, expected, actual.FileList[DockerComposeFile])
		assert.Equal(t, expectedEnv, actual.FileList[env.File])
	})
}
