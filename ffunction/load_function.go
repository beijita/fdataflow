package ffunction

import (
	"context"
	"github.com/fdataflow/fiface"
	"log"
)

type LoadFunction struct {
	BaseFunction
}

func (f *LoadFunction) Call(ctx context.Context, flow fiface.Flow) {
	log.Println("LoadFunction Call flow=", flow)
	return
}
