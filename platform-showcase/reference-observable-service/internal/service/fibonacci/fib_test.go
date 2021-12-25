package fibonacci_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	. "reference-observable-service/internal/service/fibonacci"
)

func TestCountingService_Count(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		ctx := context.Background()

		sut := NewCountingService(10)
		_, err := sut.Count(ctx, 3)

		assert.NoError(t, err)
	})

	t.Run("too distant number", func(t *testing.T) {
		ctx := context.Background()

		sut := NewCountingService(10)
		_, err := sut.Count(ctx, 999)

		assert.ErrorIs(t, err, ErrFibonacciNumberIsTooDistant)
	})

	t.Run("negative", func(t *testing.T) {
		ctx := context.Background()

		sut := NewCountingService(10)
		_, err := sut.Count(ctx, -1)

		assert.ErrorIs(t, err, ErrFibonacciNumberIsNegative)
	})
}
