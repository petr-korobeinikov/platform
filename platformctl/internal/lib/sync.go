package lib

import (
	"context"
	"os"
	"os/exec"
)

func Sync(ctx context.Context) error {
	cmd := exec.CommandContext(ctx, `go`, `mod`, `download`, `-x`)

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
