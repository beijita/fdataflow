package ffunction

import (
	"context"
	"fmt"
	"github.com/fdataflow/fiface"
	"log"
)

type CalculateFunction struct {
	BaseFunction
}

func (f *CalculateFunction) Call(ctx context.Context, flow fiface.IFlow) error {
	log.Println("CalculateFunction Call flow=", flow)
	for i, row := range flow.InputData() {
		log.Println(fmt.Sprintf(" CalculateFunction index=%d, row=%v", i, row))
		flow.CommitRow(fmt.Sprintf(" CalculateFunction index=%d, row=%v", i, row))
	}
	return nil
}
