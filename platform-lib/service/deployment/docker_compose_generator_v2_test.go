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
    restart: always
    environment:
      POSTGRES_DB: ${SERVICE_COMPONENT_POSTGRES_MASTER_DATABASE}
      POSTGRES_PASSWORD: ${SERVICE_COMPONENT_POSTGRES_MASTER_SERVICE_PASSWORD_RW}
      POSTGRES_USER: ${SERVICE_COMPONENT_POSTGRES_MASTER_SERVICE_USER_RW}
  service-component-postgres-olap:
    container_name: service-component-postgres-olap
    image: postgres:13
    restart: always
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
    image: confluentinc/cp-kafka:5.5.1
    restart: always
    depends_on:
    - platform-component-kafka-kafka-zookeeper
    environment:
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://platform-component-kafka-kafka-broker:29092,PLAINTEXT_HOST://localhost:9092
      KAFKA_AUTO_CREATE_TOPICS_ENABLE: "true"
      KAFKA_BROKER_ID: "1"
      KAFKA_GROUP_INITIAL_REBALANCE_DELAY_MS: "0"
      KAFKA_JMX_PORT: "9101"
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT
      KAFKA_LOG4J_LOGGERS: org.apache.zookeeper=ERROR,org.apache.kafka=ERROR,kafka=ERROR,kafka.cluster=ERROR,kafka.controller=ERROR,kafka.coordinator=ERROR,kafka.log=ERROR,kafka.server=ERROR,kafka.zookeeper=ERROR,state.change.logger=ERROR
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: "1"
      KAFKA_REST_HOST_NAME: platform-component-kafka-kafka-broker
      KAFKA_TRANSACTION_STATE_LOG_MIN_ISR: "1"
      KAFKA_TRANSACTION_STATE_LOG_REPLICATION_FACTOR: "1"
      KAFKA_ZOOKEEPER_CONNECT: platform-component-kafka-kafka-zookeeper:2181
  platform-component-kafka-kafka-kafdrop:
    container_name: platform-component-kafka-kafka-kafdrop
    image: obsidiandynamics/kafdrop
    restart: always
    depends_on:
    - platform-component-kafka-kafka-zookeeper
    environment:
      KAFKA_BROKERCONNECT: platform-component-kafka-kafka-broker:29092
      SERVER_PORT: "9100"
  platform-component-kafka-kafka-zookeeper:
    container_name: platform-component-kafka-kafka-zookeeper
    image: confluentinc/cp-zookeeper:5.5.1
    restart: always
    environment:
      ALLOW_ANONYMOUS_LOGIN: "yes"
      ZOOKEEPER_CLIENT_PORT: "2181"
      ZOOKEEPER_TICK_TIME: "2000"
  platform-component-opentracing-opentracing:
    container_name: platform-component-opentracing-opentracing
    image: jaegertracing/opentelemetry-all-in-one
    restart: always
  service-component-postgres-master:
    container_name: service-component-postgres-master
    image: postgres:13
    restart: always
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
    restart: always
    ports:
    - 9500:9500
    - 9501:9501
    environment:
      MINIO_ROOT_PASSWORD: ${SERVICE_COMPONENT_MINIO_MINIO_MINIO_ROOT_PASSWORD}
      MINIO_ROOT_USER: ${SERVICE_COMPONENT_MINIO_MINIO_MINIO_ROOT_USER}
    command: server /data --address ":9500" --console-address ":9501"
  service-component-postgres-master:
    container_name: service-component-postgres-master
    image: postgres:13
    restart: always
    environment:
      POSTGRES_DB: ${SERVICE_COMPONENT_POSTGRES_MASTER_DATABASE}
      POSTGRES_PASSWORD: ${SERVICE_COMPONENT_POSTGRES_MASTER_SERVICE_PASSWORD_RW}
      POSTGRES_USER: ${SERVICE_COMPONENT_POSTGRES_MASTER_SERVICE_USER_RW}
  service-component-vault-vault:
    container_name: service-component-vault-vault
    image: vault:1.9.2
    restart: always
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

	t.Run(`custom environment`, func(t *testing.T) {
		defer env.Registry().Clear()

		expectedEnv := `BAR="bar"
FOO="foo"`

		given := SpecGenerationRequest{
			ServiceName:      "wordcounter-svc",
			ServiceNamespace: "wordcounter-ns",
			Environment: map[string]string{
				"FOO": "foo",
				"BAR": "bar",
			},
		}

		sut := NewDockerComposeGeneratorV2()

		actual, err := sut.Generate(given)

		assert.NoError(t, err)
		assert.Equal(t, expectedEnv, actual.FileList[env.File])
	})

	t.Run(`platform component`, func(t *testing.T) {
		defer env.Registry().Clear()

		expected := `services:
  platform-component-kafka-kafka-broker:
    container_name: platform-component-kafka-kafka-broker
    image: confluentinc/cp-kafka:5.5.1
    restart: always
    depends_on:
    - platform-component-kafka-kafka-zookeeper
    environment:
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://platform-component-kafka-kafka-broker:29092,PLAINTEXT_HOST://localhost:9092
      KAFKA_AUTO_CREATE_TOPICS_ENABLE: "true"
      KAFKA_BROKER_ID: "1"
      KAFKA_GROUP_INITIAL_REBALANCE_DELAY_MS: "0"
      KAFKA_JMX_PORT: "9101"
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT
      KAFKA_LOG4J_LOGGERS: org.apache.zookeeper=ERROR,org.apache.kafka=ERROR,kafka=ERROR,kafka.cluster=ERROR,kafka.controller=ERROR,kafka.coordinator=ERROR,kafka.log=ERROR,kafka.server=ERROR,kafka.zookeeper=ERROR,state.change.logger=ERROR
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: "1"
      KAFKA_REST_HOST_NAME: platform-component-kafka-kafka-broker
      KAFKA_TRANSACTION_STATE_LOG_MIN_ISR: "1"
      KAFKA_TRANSACTION_STATE_LOG_REPLICATION_FACTOR: "1"
      KAFKA_ZOOKEEPER_CONNECT: platform-component-kafka-kafka-zookeeper:2181
  platform-component-kafka-kafka-kafdrop:
    container_name: platform-component-kafka-kafka-kafdrop
    image: obsidiandynamics/kafdrop
    restart: always
    depends_on:
    - platform-component-kafka-kafka-zookeeper
    environment:
      KAFKA_BROKERCONNECT: platform-component-kafka-kafka-broker:29092
      SERVER_PORT: "9100"
  platform-component-kafka-kafka-zookeeper:
    container_name: platform-component-kafka-kafka-zookeeper
    image: confluentinc/cp-zookeeper:5.5.1
    restart: always
    environment:
      ALLOW_ANONYMOUS_LOGIN: "yes"
      ZOOKEEPER_CLIENT_PORT: "2181"
      ZOOKEEPER_TICK_TIME: "2000"
  platform-component-minio-minio:
    container_name: platform-component-minio-minio
    image: quay.io/minio/minio:latest
    restart: always
  platform-component-opentracing-opentracing:
    container_name: platform-component-opentracing-opentracing
    image: jaegertracing/opentelemetry-all-in-one
    restart: always
`

		expectedEnv := ""

		given := SpecGenerationRequest{
			ServiceName:      "wordcounter-svc",
			ServiceNamespace: "wordcounter-ns",
			PlatformComponentList: []*PlatformComponent{
				{
					Name: "kafka",
					Type: "kafka",
				},
				{
					Name: "opentracing",
					Type: "opentracing",
				},
				{
					Name: "minio",
					Type: "minio",
				},
			},
		}

		sut := NewDockerComposeGeneratorV2()

		actual, err := sut.Generate(given)

		assert.NoError(t, err)
		assert.Equal(t, expected, actual.FileList[DockerComposeFile])
		assert.Equal(t, expectedEnv, actual.FileList[env.File])
	})
}
