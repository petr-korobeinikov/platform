package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path"
	"time"

	"github.com/pkg/browser"

	"github.com/pkorobeinikov/platform/platform-lib/filesystem"
	"platformctl/internal/cfg"
	"platformctl/internal/platform"
)

func Doc(ctx context.Context) error {
	if filesystem.IsDirectoryExists("doc") {
		const (
			port = 8000
		)

		ip, err := platform.IP(ctx)
		if err != nil {
			return err
		}
		host := "http://" + ip

		url := fmt.Sprintf("%s:%d", host, port)

		pwd, err := os.Getwd()
		if err != nil {
			return err
		}

		go func() {

			for {
				r, err := http.Head(url)
				if err == nil && r.StatusCode == http.StatusOK {
					break
				}

				time.Sleep(500 * time.Millisecond)
			}

			_ = browser.OpenURL(url)
		}()

		// Only mkdocs-material supported at the moment
		args := []string{
			cfg.PlatformFlavorContainerRuntimeCtl,
			"run",
			"--pull", "always",
			"--rm",
			"-p", fmt.Sprintf("%[1]d:%[1]d", port),
			"-v", fmt.Sprintf("%s:/docs", path.Join(pwd, "doc")),
			"squidfunk/mkdocs-material",
		}
		cmd := exec.CommandContext(ctx, args[0], args[1:]...)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		// Run doc in blocking mode
		// Need a graceful shutdown?
		return cmd.Run()
	}

	return ErrDocDirDoesNotExists
}

var (
	ErrDocDirDoesNotExists = errors.New("doc directory does not exists")
)
