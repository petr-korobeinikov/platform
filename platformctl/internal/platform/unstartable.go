package platform

import (
	"context"
	"errors"
)

func (b *unstartableBridge) Start(ctx context.Context) error {
	return ErrUnstartableVirturalMachine
}

func (b *unstartableBridge) Stop(ctx context.Context) error {
	return ErrUnstartableVirturalMachine
}

func newUnstartableBridge() *unstartableBridge {
	return &unstartableBridge{}
}

type unstartableBridge struct {
}

var ErrUnstartableVirturalMachine = errors.New("this vm does not require start and stop")
