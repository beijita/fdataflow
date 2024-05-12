package ffunction

import (
	"context"
	"github.com/fdataflow/fiface"
	"log"
)

type CalculateFunction struct {
	BaseFunction
}

func (f *CalculateFunction) Call(ctx context.Context, flow fiface.IFlow) error {
	log.Println("CalculateFunction Call flow=", flow)
	return nil
}
