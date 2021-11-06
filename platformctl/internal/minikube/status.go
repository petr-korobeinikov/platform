package minikube

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"reflect"

	"platformctl/internal/cfg"
)

func IsRunning(ctx context.Context) (bool, error) {
	var (
		b bytes.Buffer
		j status
	)

	args := []string{
		"minikube",
		fmt.Sprintf("--profile=%s", cfg.MinikubeProfile()),
		"--output=json",
		"status",
	}

	cmd := exec.CommandContext(ctx, args[0], args[1:]...)

	cmd.Stdout = &b
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return false, err
	}

	err := json.Unmarshal(b.Bytes(), &j)
	if err != nil {
		return false, err
	}

	if !reflect.DeepEqual(expectedRunningStatus, j) {
		return false, ErrMinikubeIsNotRunning
	}

	return true, nil
}

var (
	ErrMinikubeIsNotRunning = errors.New(`minikube is not running`)

	expectedRunningStatus = status{
		Name:       cfg.MinikubeProfile(),
		Host:       "Running",
		Kubelet:    "Running",
		APIServer:  "Running",
		Kubeconfig: "Configured",
	}
)

type (
	status struct {
		Name       string `json:"Name"`
		Host       string `json:"Host"`
		Kubelet    string `json:"Kubelet"`
		APIServer  string `json:"APIServer"`
		Kubeconfig string `json:"Kubeconfig"`
	}
)
