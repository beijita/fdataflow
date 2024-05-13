package ffunction

import (
	"context"
	"fmt"
	"github.com/fdataflow/fflow"
	"github.com/fdataflow/fiface"
	"log"
)

type LoadFunction struct {
	BaseFunction
}

func (f *LoadFunction) Call(ctx context.Context, flow fiface.IFlow) error {
	log.Println("LoadFunction Call flow=", flow)
	err := fflow.Pool().CallFunction(ctx, f.GetConfig().FName, flow)
	if err != nil {
		log.Println(fmt.Sprintf(" LoadFunction FName=%s, flow=%v", f.GetConfig().FName, flow))
		return err
	}
	return nil
}
