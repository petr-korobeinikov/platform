package docker

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"os"
	"os/exec"
)

func IsConformsToMinimalRequirements(ctx context.Context) (bool, error) {
	var (
		b bytes.Buffer
		j info
	)

	args := []string{
		"docker",
		"info",
		`--format={{ json . }}`,
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

	if j.NCPU < expectedMinimalInfo.NCPU || j.MemTotal < expectedMinimalInfo.MemTotal {
		return false, ErrNotConformsToMinimalRequirements
	}

	return true, nil
}

var (
	ErrNotConformsToMinimalRequirements = errors.New(`docker setup is not conforms to minimal requirements`)

	expectedMinimalInfo = info{
		NCPU:     4,
		MemTotal: 4173522944,
	}
)

type (
	info struct {
		NCPU     uint `json:"NCPU"`
		MemTotal uint `json:"MemTotal"`
	}
)
