package platform

import (
	"context"
	"errors"
)

func (b *unsupportedBridge) Start(ctx context.Context) error {
	return ErrUnsupportedVirturalMachine
}

func (b *unsupportedBridge) Stop(ctx context.Context) error {
	return ErrUnsupportedVirturalMachine
}

func newUnsupportedBridge() *unsupportedBridge {
	return &unsupportedBridge{}
}

type unsupportedBridge struct {
}

var ErrUnsupportedVirturalMachine = errors.New("unsupported virtual machine")
