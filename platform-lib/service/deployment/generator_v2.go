package deployment

var (
	_ generatorV2 = (*dockerComposeGeneratorV2)(nil)
)

type (
	generatorV2 interface {
		Generate(SpecGenerationRequest) (SpecGenerationResponse, error)
	}

	SpecGenerationRequest struct {
		ServiceName           string
		ServiceNamespace      string
		IP                    string
		Environment           map[string]string
		ServiceComponentList  []*ServiceComponent
		PlatformComponentList []*PlatformComponent
	}

	SpecGenerationResponse struct {
		FileList map[string]string
	}
)
