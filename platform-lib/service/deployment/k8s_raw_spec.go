package deployment

import "github.com/pkorobeinikov/platform/platform-lib/service/spec"

func (g *K8SRawSpecGenerator) Generate(s *spec.Spec) ([]byte, error) {
	panic(ErrUnimplementedGenerator)
}

func NewK8SRawSpecGenerator() *K8SRawSpecGenerator {
	return &K8SRawSpecGenerator{}
}

type K8SRawSpecGenerator struct {
}
