package platform

import (
	"context"
	"errors"
)

func (b *unsupportedBridge) Start(context.Context) error {
	return ErrUnsupportedVirturalMachine
}

func (b *unsupportedBridge) Stop(context.Context) error {
	return ErrUnsupportedVirturalMachine
}

func (b *unsupportedBridge) IP(context.Context) (string, error) {
	return "", ErrUnsupportedVirturalMachine
}

func newUnsupportedBridge() *unsupportedBridge {
	return &unsupportedBridge{}
}

type unsupportedBridge struct {
}

var ErrUnsupportedVirturalMachine = errors.New("unsupported virtual machine")
