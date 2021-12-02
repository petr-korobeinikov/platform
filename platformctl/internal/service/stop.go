package service

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/pkorobeinikov/platform/platform-lib/service/spec"
)

func Stop(ctx context.Context) error {
	fmt.Println("Debug service")

	spec, err := spec.Read()
	if err != nil {
		return err
	}

	fmt.Println("Service name:", spec.Name)

	cmd := exec.CommandContext(ctx, `docker-compose`, `down`, `--remove-orphans`)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
