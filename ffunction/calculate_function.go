package ffunction

import (
	"context"
	"fmt"
	"github.com/fdataflow/fflow"
	"github.com/fdataflow/fiface"
	"log"
)

type CalculateFunction struct {
	BaseFunction
}

func (f *CalculateFunction) Call(ctx context.Context, flow fiface.IFlow) error {
	log.Println("CalculateFunction Call flow=", flow)
	err := fflow.Pool().CallFunction(ctx, f.GetConfig().FName, flow)
	if err != nil {
		log.Println(fmt.Sprintf(" CalculateFunction FName=%s, flow=%v", f.GetConfig().FName, flow))
		return err
	}
	return nil
}
