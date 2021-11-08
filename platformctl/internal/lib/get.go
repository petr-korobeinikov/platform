package lib

import (
	"context"
	"os"
	"os/exec"
)

func Get(ctx context.Context, lib string) error {
	cmd := exec.CommandContext(ctx, `go`, `get`, `-x`, lib)

	// todo extract env into config
	//cmd.Env = append(os.Environ(), []string{
	//	`GONOPROXY=none`,
	//	`GOPROXY=https://nexus.dev.cloud.mts.ru/repository/golang-internal/`,
	//	`GONOSUMDB=dev.cloud.mts.ru/*`,
	//	`GOPRIVATE=*.dev.cloud.mts.ru`,
	//}...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
