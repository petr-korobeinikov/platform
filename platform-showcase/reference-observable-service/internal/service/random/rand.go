package random

import (
	"context"
	"math/rand"
	"time"

	"github.com/opentracing/opentracing-go"
)

func (g *Generator) Generate(ctx context.Context) int {
	span, ctx := opentracing.StartSpanFromContext(ctx, "random.Generator_Generate")
	defer span.Finish()

	// Emulate long operation
	time.Sleep(30 * time.Millisecond)

	n := rand.Intn(g.max-g.min+1) + g.min
	span.SetTag("n", n)

	return n
}

func NewGenerator(min, max int) *Generator {
	return &Generator{
		min: min,
		max: max,
	}
}

type Generator struct {
	min int
	max int
}
