package platform

import (
	"context"
	"errors"
)

func (b *unstartableBridge) Start(context.Context) error {
	return ErrUnstartableVirturalMachine
}

func (b *unstartableBridge) Stop(context.Context) error {
	return ErrUnstartableVirturalMachine
}

func (b *unstartableBridge) IP(context.Context) (string, error) {
	return "127.0.0.1", nil
}

func newUnstartableBridge() *unstartableBridge {
	return &unstartableBridge{}
}

type unstartableBridge struct {
}

var ErrUnstartableVirturalMachine = errors.New("this vm does not require start and stop")
