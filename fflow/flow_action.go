package fflow

import (
	"context"
	"github.com/fdataflow/fiface"
)

func (flow *DataFlow) dealAction(ctx context.Context, fn fiface.IFunction) (fiface.IFunction, error) {
	err := flow.commitCurData(ctx)
	if err != nil {
		return nil, err
	}
	flow.PrevFunctionID = flow.ThisFunctionID
	fn = fn.Next()
	if flow.act.Abort {
		flow.abort = true
	}
	flow.act = fiface.Action{}
	return fn, nil
}
