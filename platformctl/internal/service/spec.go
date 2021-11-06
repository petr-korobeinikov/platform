package service

import (
	"context"
	"fmt"
)

func ReadSpec(ctx context.Context) (*Spec, error) {
	fmt.Println("detect service in current directory")

	return nil, nil
}

type (
	Spec struct {
		Name string
	}
)
