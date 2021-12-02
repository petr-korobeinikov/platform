package deployment

import "github.com/pkorobeinikov/platform/platform-lib/service/spec"

func (g *HelmChartGenerator) Generate(s *spec.Spec) ([]byte, error) {
	panic(ErrUnimplementedGenerator)
}

func NewHelmChartGenerator() *HelmChartGenerator {
	return &HelmChartGenerator{}
}

type HelmChartGenerator struct {
}
