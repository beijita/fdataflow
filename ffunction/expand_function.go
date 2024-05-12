package ffunction

import (
	"context"
	"github.com/fdataflow/fiface"
	"log"
)

type ExpandFunction struct {
	BaseFunction
}

func (f *ExpandFunction) Call(ctx context.Context, flow fiface.Flow) error {
	log.Println("ExpandFunction Call flow=", flow)
	return nil
}
