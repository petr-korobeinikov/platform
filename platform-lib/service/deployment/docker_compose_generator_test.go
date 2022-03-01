package deployment_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/pkorobeinikov/platform/platform-lib/service/deployment"
	. "github.com/pkorobeinikov/platform/platform-lib/service/spec"
)

func TestDockerComposeGenerator_Generate(t *testing.T) {
	given := &Spec{
		Name: "wordcounter",
		Component: []*Component{
			{
				Type: "postgres",
				Name: "postgres",
			},
			{
				Type: "vault",
				Name: "vault",
			},
		},
	}

	sut := NewDockerComposeGenerator()

	actual, _ := sut.Generate(given)

	assert.Equal(t, string(expected), string(actual))
}

var (
	expected = []byte(`version: "3"
services:
  component-postgres-postgres:
    container_name: component-postgres-postgres
    image: postgres:14
    restart: always
    ports:
    - 5432:5432
    environment:
      POSTGRES_DB: ${COMPONENT_POSTGRES_POSTGRES_DB}
      POSTGRES_PASSWORD: ${COMPONENT_POSTGRES_POSTGRES_PASSWORD}
      POSTGRES_USER: ${COMPONENT_POSTGRES_POSTGRES_USER}
  component-vault-vault:
    container_name: component-vault-vault
    image: vault:1.9.2
    restart: always
    ports:
    - 8200:8200
    environment:
      VAULT_DEV_LISTEN_ADDRESS: ${COMPONENT_VAULT_VAULT_VAULT_DEV_LISTEN_ADDRESS}
      VAULT_DEV_ROOT_TOKEN_ID: ${COMPONENT_VAULT_VAULT_VAULT_DEV_ROOT_TOKEN_ID}
    cap_add:
    - IPC_LOCK
  platform-kafka-broker:
    container_name: kafka-broker
    image: confluentinc/cp-kafka:5.5.1
    restart: on-failure
    depends_on:
    - platform-kafka-zookeeper
    ports:
    - 9092:9092
    environment:
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://kafka-broker:29092,PLAINTEXT_HOST://localhost:9092
      KAFKA_AUTO_CREATE_TOPICS_ENABLE: "true"
      KAFKA_BROKER_ID: "1"
      KAFKA_GROUP_INITIAL_REBALANCE_DELAY_MS: "0"
      KAFKA_JMX_PORT: "9101"
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT
      KAFKA_LOG4J_LOGGERS: org.apache.zookeeper=ERROR,org.apache.kafka=ERROR,kafka=ERROR,kafka.cluster=ERROR,kafka.controller=ERROR,kafka.coordinator=ERROR,kafka.log=ERROR,kafka.server=ERROR,kafka.zookeeper=ERROR,state.change.logger=ERROR
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: "1"
      KAFKA_REST_HOST_NAME: broker
      KAFKA_TRANSACTION_STATE_LOG_MIN_ISR: "1"
      KAFKA_TRANSACTION_STATE_LOG_REPLICATION_FACTOR: "1"
      KAFKA_ZOOKEEPER_CONNECT: kafka-zookeeper:2181
  platform-kafka-kafdrop:
    container_name: kafka-kafdrop
    image: obsidiandynamics/kafdrop
    restart: always
    depends_on:
    - platform-kafka-broker
    ports:
    - 9100:9100
    environment:
      KAFKA_BROKERCONNECT: kafka-broker:29092
      SERVER_PORT: "9100"
  platform-kafka-zookeeper:
    container_name: kafka-zookeeper
    image: confluentinc/cp-zookeeper:5.5.1
    restart: always
    environment:
      ALLOW_ANONYMOUS_LOGIN: "yes"
      ZOOKEEPER_CLIENT_PORT: "2181"
      ZOOKEEPER_TICK_TIME: "2000"
  platform-observability-opentelemetry:
    container_name: opentelemetry
    image: jaegertracing/opentelemetry-all-in-one
    restart: always
    ports:
    - 6831:6831
    - 16686:16686
    - 14268:14268
  platform-sentinel-wordcounter:
    container_name: platform-sentinel
    image: kubernetes/pause
    restart: always
  service:
    container_name: service
    image: ${SERVICE_IMAGE_NAME}:${SERVICE_IMAGE_TAG}
    restart: always
    ports:
    - 9000:9000
    environment:
      COMPONENT_POSTGRES_POSTGRES_DB: ${COMPONENT_POSTGRES_POSTGRES_DB}
      COMPONENT_POSTGRES_POSTGRES_PASSWORD: ${COMPONENT_POSTGRES_POSTGRES_PASSWORD}
      COMPONENT_POSTGRES_POSTGRES_USER: ${COMPONENT_POSTGRES_POSTGRES_USER}
      COMPONENT_VAULT_VAULT_VAULT_DEV_LISTEN_ADDRESS: ${COMPONENT_VAULT_VAULT_VAULT_DEV_LISTEN_ADDRESS}
      COMPONENT_VAULT_VAULT_VAULT_DEV_ROOT_TOKEN_ID: ${COMPONENT_VAULT_VAULT_VAULT_DEV_ROOT_TOKEN_ID}
      KAFKA_PORT: ${KAFKA_PORT}
      OBSERVABILITY_JAEGER_COLLECTOR_HTTP_ENDPOINT: ${OBSERVABILITY_JAEGER_COLLECTOR_HTTP_ENDPOINT}
      SERVICE: ${SERVICE}
      SERVICE_IMAGE_NAME: ${SERVICE_IMAGE_NAME}
      SERVICE_IMAGE_TAG: ${SERVICE_IMAGE_TAG}
`)
)
