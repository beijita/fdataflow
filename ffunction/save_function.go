package ffunction

import (
	"context"
	"fmt"
	"github.com/fdataflow/fflow"
	"github.com/fdataflow/fiface"
	"log"
)

type SaveFunction struct {
	BaseFunction
}

func (f *SaveFunction) Call(ctx context.Context, flow fiface.IFlow) error {
	log.Println("SaveFunction Call flow=", flow)
	err := fflow.Pool().CallFunction(ctx, f.GetConfig().FName, flow)
	if err != nil {
		log.Println(fmt.Sprintf(" SaveFunction FName=%s, flow=%v", f.GetConfig().FName, flow))
		return err
	}
	return nil
}
