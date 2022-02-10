package main

import (
	"context"
	"fmt"

	"service-boilerplate-go/internal/greeting"
)

func main() {
	ctx := context.Background()

	greeter := greeting.NewGreeter()

	hello, err := greeter.Greet(ctx, "service-boilerplate-go")
	if err != nil {
		return
	}

	fmt.Println(hello)
}
