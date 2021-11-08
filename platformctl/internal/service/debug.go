package service

import (
	"context"
	"fmt"
)

func Debug(ctx context.Context) error {
	fmt.Println("Debug service")

	return nil
}
