package deployment_test

import (
	"bytes"
	"testing"

	. "github.com/pkorobeinikov/platform/platform-lib/service/deployment"
	. "github.com/pkorobeinikov/platform/platform-lib/service/spec"
)

func TestDockerComposeGenerator_Generate(t *testing.T) {
	given := &Spec{
		Name: "wordcounter",
		Component: []Component{
			{
				Type:    "postgres",
				Enabled: true,
			},
		},
	}

	sut := NewDockerComposeGenerator()

	actual, _ := sut.Generate(given)
	if !bytes.Equal(expected, actual) {
		t.Errorf("docker compose spec is generated incorrectly")
	}
}

var (
	expected = []byte(`version: "3"
services:
  app:
    container_name: app
    image: ${SERVICE_IMAGE_NAME}:${SERVICE_IMAGE_TAG}
    restart: always
    ports:
    - 9000:9000
  postgres:
    container_name: postgres
    image: postgres:14
    restart: always
    ports:
    - 5432:5432
    environment:
      POSTGRES_DB: ${POSTGRES_DB}
      POSTGRES_PASSWORD: ${POSTGRES_PASSWORD}
      POSTGRES_USER: ${POSTGRES_USER}
`)
)