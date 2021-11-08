package service

import (
	"context"
	"fmt"
)

func Create(ctx context.Context, name string) error {
	fmt.Println("Create service " + name)

	return nil
}
