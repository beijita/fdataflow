package ffunction

import (
	"context"
	"github.com/fdataflow/fiface"
	"log"
)

type SaveFunction struct {
	BaseFunction
}

func (f *SaveFunction) Call(ctx context.Context, flow fiface.IFlow) error {
	log.Println("SaveFunction Call flow=", flow)
	return nil
}
