package fflow

import (
	"context"
	"github.com/fdataflow/config"
	"github.com/fdataflow/fcommon"
	"github.com/fdataflow/fid"
	"github.com/fdataflow/fiface"
)

type DataFlow struct {
	ID   string
	Name string
}

func (flow *DataFlow) Run(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (flow *DataFlow) Link(fConf *config.FuncConfig) {
	//TODO implement me
	panic("implement me")
}

func NewDataFlow(conf *config.DataFlowConfig) fiface.Flow {
	flow := new(DataFlow)
	flow.ID = fid.DataFlowID(fcommon.DataFlowIDTypeFlow)
	flow.Name = conf.FlowName
	return flow
}
