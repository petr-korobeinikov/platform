package greeting

import (
	"context"
	"fmt"
)

func (g *greeter) Greet(ctx context.Context, name string) (string, error) {
	return fmt.Sprintf("Hello, %s!", name), nil
}

func NewGreeter() *greeter {
	return &greeter{}
}

type greeter struct {
}
