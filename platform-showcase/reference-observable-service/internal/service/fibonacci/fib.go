package fibonacci

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/opentracing/opentracing-go"
)

func (s *CountingService) Count(ctx context.Context, n int) (int, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "fibonacci.CountingService_Count")
	defer span.Finish()

	// Tag adds to "Tags" section
	span.SetTag("max_fib_number", s.maxFibNumber)

	// BaggageItem adds to "Log" section. Use "baggage" wisely.
	span.SetBaggageItem("max_fib_number", strconv.Itoa(s.maxFibNumber))

	if n > s.maxFibNumber {
		span.SetTag("error", true)
		span.SetTag("error.message", ErrFibonacciNumberIsTooDistant)

		return 0, ErrFibonacciNumberIsTooDistant
	}

	if n < 0 {
		span.SetTag("error", true)
		span.SetTag("error.message", ErrFibonacciNumberIsNegative)
		return 0, ErrFibonacciNumberIsNegative
	}

	return fib(ctx, n), nil
}

func NewCountingService(maxFibNumber int) *CountingService {
	return &CountingService{
		maxFibNumber: maxFibNumber,
	}
}

type CountingService struct {
	maxFibNumber int
}

func fib(ctx context.Context, n int) int {
	span, ctx := opentracing.StartSpanFromContext(ctx, "fibonacci.fib")
	defer span.Finish()

	// Emulate long operation
	time.Sleep(3 * time.Millisecond)

	if n == 0 {
		span.SetTag("result", 0)
		return 0
	}

	if n < 3 {
		span.SetTag("result", 1)
		return 1
	}

	span.SetTag("n-1", n-1)
	span.SetTag("n-2", n-2)

	return fib(ctx, n-1) + fib(ctx, n-2)
}

var (
	ErrFibonacciNumberIsTooDistant = errors.New("fibonacci number is too distant")
	ErrFibonacciNumberIsNegative   = errors.New("fibonacci number is negative")
)
