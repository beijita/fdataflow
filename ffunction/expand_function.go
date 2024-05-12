package ffunction

import (
	"context"
	"fmt"
	"github.com/fdataflow/fiface"
	"log"
)

type ExpandFunction struct {
	BaseFunction
}

func (f *ExpandFunction) Call(ctx context.Context, flow fiface.IFlow) error {
	log.Println("ExpandFunction Call flow=", flow)
	for i, row := range flow.InputData() {
		log.Println(fmt.Sprintf(" ExpandFunction index=%d, row=%v", i, row))
	}
	return nil
}
