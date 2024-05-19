package fflow

import (
	"context"
	"github.com/fdataflow/fiface"
)

func (flow *DataFlow) dealAction(ctx context.Context, fn fiface.IFunction) (fiface.IFunction, error) {
	var err error
	if flow.act.DataReuse {
		err = flow.commitReuseData(ctx)
	} else {
		err = flow.commitCurData(ctx)
	}
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

func (flow *DataFlow) commitReuseData(ctx context.Context) error {
	if len(flow.data[flow.PrevFunctionID]) == 0 {
		flow.abort = true
		return nil
	}
	flow.data[flow.ThisFunctionID] = flow.data[flow.PrevFunctionID]
	flow.buffer = flow.buffer[0:0]
	return nil
}
