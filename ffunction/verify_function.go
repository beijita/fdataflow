package ffunction

import (
	"context"
	"github.com/fdataflow/fiface"
	"log"
)

type VerifyFunction struct {
	BaseFunction
}

func (v *VerifyFunction) Call(ctx context.Context, flow fiface.Flow) error {
	log.Println("VerifyFunction Call flow=", flow)
	return nil
}
