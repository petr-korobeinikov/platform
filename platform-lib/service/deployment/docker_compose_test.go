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
				Type:    "postgres",
				Name:    "postgres",
				Enabled: true,
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
  component_postgres_postgres:
    container_name: component_postgres_postgres
    image: postgres:14
    restart: always
    ports:
    - 5432:5432
    environment:
      POSTGRES_DB: ${COMPONENT_POSTGRES_POSTGRES_DB}
      POSTGRES_PASSWORD: ${COMPONENT_POSTGRES_POSTGRES_PASSWORD}
      POSTGRES_USER: ${COMPONENT_POSTGRES_POSTGRES_USER}
  platform_observability_opentelemetry:
    container_name: opentelemetry
    image: jaegertracing/opentelemetry-all-in-one
    restart: always
    ports:
    - 6831:6831
    - 16686:16686
    - 14268:14268
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
      OBSERVABILITY_JAEGER_COLLECTOR_HTTP_ENDPOINT: ${OBSERVABILITY_JAEGER_COLLECTOR_HTTP_ENDPOINT}
      SERVICE: ${SERVICE}
      SERVICE_IMAGE_NAME: ${SERVICE_IMAGE_NAME}
      SERVICE_IMAGE_TAG: ${SERVICE_IMAGE_TAG}
`)
)
