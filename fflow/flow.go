package fflow

import (
	"context"
	"errors"
	"github.com/fdataflow/config"
	"github.com/fdataflow/fcommon"
	"github.com/fdataflow/ffunction"
	"github.com/fdataflow/fid"
	"github.com/fdataflow/fiface"
	"sync"
)

type DataFlow struct {
	ID             string
	Name           string
	Conf           *config.DataFlowConfig
	FuncMap        map[string]fiface.IFunction
	FlowHead       fiface.IFunction
	FlowTail       fiface.IFunction
	flowLock       sync.RWMutex
	ThisFunction   fiface.IFunction
	ThisFunctionID string
	PrevFunctionID string
	funcParams     map[string]config.FParam
	fParamsLock    sync.RWMutex

	buffer    fcommon.DataFlowRowArr
	data      fcommon.DataFlowDataMap
	inputData fcommon.DataFlowRowArr
}

func (flow *DataFlow) GetName() string {
	return flow.Name
}

func (flow *DataFlow) GetThisFUnction() fiface.IFunction {
	return flow.ThisFunction
}

func (flow *DataFlow) GetThisFuncConf() *config.FuncConfig {
	return flow.ThisFunction.GetConfig()
}

func (flow *DataFlow) CommitRow(row interface{}) error {
	flow.buffer = append(flow.buffer, row)
	return nil
}

func (flow *DataFlow) Run(ctx context.Context) error {
	flowNode := flow.FlowHead
	if flow.Conf.Status == int(fcommon.Disable) {
		return nil
	}

	flow.PrevFunctionID = fcommon.FunctionIDFirstVirtual
	err := flow.commitSrcData(ctx)
	if err != nil {
		return err
	}
	for flowNode != nil {
		flowID := flowNode.GetID()
		flow.ThisFunction = flowNode
		flow.ThisFunctionID = flowID
		inputData, err := flow.getCurData(ctx)
		if err != nil {
			return err
		} else {
			flow.inputData = inputData
		}

		err = flowNode.Call(ctx, flow)
		if err != nil {
			return err
		} else {
			err = flow.commitCurData(ctx)
			if err != nil {
				return err
			}
			flow.PrevFunctionID = flow.ThisFunctionID
			flowNode = flowNode.Next()
		}
	}
	return nil
}

func (flow *DataFlow) Link(fConf *config.FuncConfig, fParam config.FParam) error {
	f := ffunction.NewDataFlowFunction(flow, fConf)
	flow.appendFunc(f, fParam)
	return nil
}

func (flow *DataFlow) appendFunc(f fiface.IFunction, param config.FParam) error {
	if f == nil {
		return errors.New("")
	}
	flow.fParamsLock.Lock()
	defer flow.fParamsLock.Unlock()
	if flow.FlowHead == nil {
		flow.FlowHead = f
		flow.FlowTail = f

		f.SetN(nil)
		f.SetP(nil)
	} else {
		f.SetP(flow.FlowTail)
		f.SetN(nil)
		flow.FlowTail.SetN(f)
		flow.FlowTail = f
	}
	flow.FuncMap[f.GetID()] = f
	params := make(config.FParam)
	for k, v := range f.GetConfig().Option.Params {
		params[k] = v
	}
	for k, v := range param {
		params[k] = v
	}
	flow.funcParams[f.GetID()] = params
	return nil
}

func NewDataFlow(conf *config.DataFlowConfig) fiface.IFlow {
	flow := new(DataFlow)
	flow.ID = fid.DataFlowID(fcommon.DataFlowIDTypeFlow)
	flow.Name = conf.FlowName
	flow.Conf = conf
	flow.FuncMap = make(map[string]fiface.IFunction)
	flow.funcParams = make(map[string]config.FParam)
	flow.data = make(fcommon.DataFlowDataMap)
	return flow
}

func (flow *DataFlow) commitSrcData(ctx context.Context) error {
	dataLen := len(flow.buffer)
	rowArr := make(fcommon.DataFlowRowArr, 0, dataLen)
	for _, row := range flow.buffer {
		rowArr = append(rowArr, row)
	}
	flow.clearData(flow.data)
	flow.data[fcommon.FunctionIDFirstVirtual] = rowArr
	flow.buffer = flow.buffer[0:0]
	return nil
}

func (flow *DataFlow) clearData(dataMap fcommon.DataFlowDataMap) {
	for k, _ := range dataMap {
		delete(dataMap, k)
	}
}

func (flow *DataFlow) commitCurData(ctx context.Context) error {
	dataLen := len(flow.buffer)
	if dataLen == 0 {
		return nil
	}
	rowArr := make(fcommon.DataFlowRowArr, 0, dataLen)
	for _, row := range flow.buffer {
		rowArr = append(rowArr, row)
	}
	flow.data[flow.ThisFunctionID] = rowArr
	flow.buffer = flow.buffer[0:0]
	return nil
}

func (flow *DataFlow) getCurData(ctx context.Context) (fcommon.DataFlowRowArr, error) {
	if "" == flow.PrevFunctionID {
		return nil, errors.New("error flow.PrevFunctionID is blank string ")
	}
	if _, ok := flow.data[flow.PrevFunctionID]; !ok {
		return nil, errors.New("error flow.data is not contain flow.PrevFunctionID ")
	}
	return flow.data[flow.PrevFunctionID], nil
}

func (flow *DataFlow) InputData() fcommon.DataFlowRowArr {
	return flow.inputData
}
