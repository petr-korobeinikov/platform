package greeting_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	. "service-boilerplate-go/internal/greeting"
)

func TestGreeter_Greet(t *testing.T) {
	t.Run(`positive`, func(t *testing.T) {
		sut := NewGreeter()

		actual, err := sut.Greet(context.Background(), "World")

		assert.NoError(t, err)
		assert.Equal(t, "Hello, World!", actual)
	})
}
