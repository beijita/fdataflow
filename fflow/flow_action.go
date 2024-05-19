package fflow

import (
	"context"
	"fmt"
	"github.com/fdataflow/fiface"
)

func (flow *DataFlow) dealAction(ctx context.Context, fn fiface.IFunction) (fiface.IFunction, error) {
	var err error
	if flow.act.DataReuseFlag {
		err = flow.commitReuseData(ctx)
	} else {
		err = flow.commitCurData(ctx)
	}
	if err != nil {
		return nil, err
	}
	if flow.act.JumpFunc != "" {
		if _, ok := flow.FuncMap[flow.act.JumpFunc]; !ok {
			return nil, fmt.Errorf("")
		}
		jumpFunc := flow.FuncMap[flow.act.JumpFunc]
		flow.ThisFunctionID = jumpFunc.GetPreID()
		fn = jumpFunc
		flow.abort = false
	} else {
		flow.PrevFunctionID = flow.ThisFunctionID
		fn = fn.Next()
	}
	if flow.act.AbortFlag {
		flow.abort = true
	}
	flow.act = fiface.Action{}
	return fn, nil
}

func (flow *DataFlow) commitReuseData(ctx context.Context) error {
	if !flow.act.ForceEntryNextFlag && len(flow.data[flow.PrevFunctionID]) == 0 {
		flow.abort = true
		return nil
	}
	flow.data[flow.ThisFunctionID] = flow.data[flow.PrevFunctionID]
	flow.buffer = flow.buffer[0:0]
	return nil
}
