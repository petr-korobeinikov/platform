package service

import (
	"context"
	"os"
	"os/exec"

	"github.com/pkorobeinikov/platform/platform-lib/service/deployment"
	"github.com/pkorobeinikov/platform/platform-lib/service/env"
	"github.com/pkorobeinikov/platform/platform-lib/service/spec"
	"platformctl/internal/cfg"
	"platformctl/internal/platform"
)

func Debug(ctx context.Context) error {
	s, err := spec.Read()
	if err != nil {
		return err
	}

	ip, err := platform.IP(ctx)
	if err != nil {
		return err
	}

	generator := deployment.NewDockerComposeGeneratorV2()

	var serviceComponentList []*deployment.ServiceComponent
	for _, serviceComponent := range s.Component {
		serviceComponentList = append(serviceComponentList, &deployment.ServiceComponent{
			Name: serviceComponent.Name,
			Type: serviceComponent.Type,
		})
	}

	// Need to be extracted from the config
	platformComponentList := []*deployment.PlatformComponent{
		{
			Name: "kafka",
			Type: "kafka",
		},
		{
			Name: "opentracing",
			Type: "opentracing",
		},
		{
			Name: "minio",
			Type: "minio",
		},
	}

	deploymentSpec, err := generator.Generate(deployment.SpecGenerationRequest{
		ServiceName:           s.Name,
		ServiceNamespace:      "default",
		IP:                    ip,
		DeploymentType:        deployment.TypePartial,
		Environment:           s.EnvironmentFor(cfg.ServiceEnv),
		ServiceComponentList:  serviceComponentList,
		PlatformComponentList: platformComponentList,
	})
	if err != nil {
		return err
	}

	// Needs rework with filesystem.WriteFile()
	err = deployment.WriteDockerComposeFile([]byte(deploymentSpec.FileList[deployment.DockerComposeFile]))
	if err != nil {
		return err
	}

	// Needs rework with filesystem.WriteFile()
	err = os.WriteFile(env.File, []byte(deploymentSpec.FileList[env.File]), 0644)
	if err != nil {
		return err
	}

	args := deployment.DockerComposeArgs(cfg.PlatformFlavorContainerRuntimeCtl, s.Name, `up`, `-d`)

	cmd := exec.CommandContext(ctx, args[0], args[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
