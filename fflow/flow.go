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

func (flow *DataFlow) CommitRow(row interface{}) error {
	return nil
}

func (flow *DataFlow) Run(ctx context.Context) error {
	flowNode := flow.FlowHead
	if flow.Conf.Status == int(fcommon.Disable) {
		return nil
	}
	for flowNode != nil {
		err := flowNode.Call(ctx, flow)
		if err != nil {
			return err
		} else {
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
