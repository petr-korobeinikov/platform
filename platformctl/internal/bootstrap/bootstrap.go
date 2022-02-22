package bootstrap

import (
	"context"
	"path"

	"github.com/carlmjohnson/requests"
	"github.com/mitchellh/go-homedir"
)

func Bootstrap(ctx context.Context, bootstrapURL string) error {
	home, err := homedir.Dir()
	if err != nil {
		return err
	}

	return requests.
		URL(bootstrapURL).
		ToFile(path.Join(home, ".platformctl", "platformctl.yaml")).
		Fetch(ctx)
}
