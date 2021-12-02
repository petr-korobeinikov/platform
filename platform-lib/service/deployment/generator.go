package deployment

import (
	"errors"

	"github.com/pkorobeinikov/platform/platform-lib/service/spec"
)

var (
	_ generator = (*DockerComposeGenerator)(nil)
	_ generator = (*HelmChartGenerator)(nil)
	_ generator = (*K8SRawSpecGenerator)(nil)

	ErrUnimplementedGenerator = errors.New("this type of generator is currently unimplemented")
)

type generator interface {
	Generate(spec2 *spec.Spec) ([]byte, error)
}
