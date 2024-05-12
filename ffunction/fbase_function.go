package ffunction

import (
	"context"
	"github.com/fdataflow/config"
	"github.com/fdataflow/fcommon"
	"github.com/fdataflow/fid"
	"github.com/fdataflow/fiface"
)

type BaseFunction struct {
	ID       string
	Config   *config.FuncConfig
	Flow     fiface.Flow
	NextFunc fiface.IFunction
	PrevFunc fiface.IFunction
}

func NewDataFlowFunction(flow fiface.Flow, config *config.FuncConfig) fiface.IFunction {
	var f fiface.IFunction
	switch fcommon.DataFlowMode(config.FMode) {
	case fcommon.Verify:
		f = new(VerifyFunction)
	case fcommon.Save:
		f = new(SaveFunction)
	case fcommon.Load:
		f = new(LoadFunction)
	case fcommon.Calculate:
		f = new(CalculateFunction)
	case fcommon.Expand:
		f = new(ExpandFunction)
	default:
		return nil
	}
	f.CreateID()
	f.SetConfig(config)
	f.SetFlow(flow)
	return f
}

func (b *BaseFunction) Call(ctx context.Context, flow fiface.Flow) {
}

func (b *BaseFunction) SetConfig(fc *config.FuncConfig) {
	if fc == nil {
		return
	}
	b.Config = fc
}

func (b *BaseFunction) GetConfig() *config.FuncConfig {
	return b.Config
}

func (b *BaseFunction) SetFlow(f fiface.Flow) {
	b.Flow = f
}

func (b *BaseFunction) GetFlow() fiface.Flow {
	return b.Flow
}

func (b *BaseFunction) CreateID() {
	b.ID = fid.DataFlowID(fcommon.DataFlowIDTypeFunction)
}

func (b *BaseFunction) GetID() string {
	return b.ID
}

func (b *BaseFunction) GetPreID() string {
	if b.PrevFunc == nil {
		return fcommon.FunctionIDFirstVirtual
	}
	return b.PrevFunc.GetID()
}

func (b *BaseFunction) GetNextID() string {
	if b.NextFunc == nil {
		return fcommon.FunctionIDLastVirtual
	}
	return b.NextFunc.GetID()
}

func (b *BaseFunction) Next() fiface.IFunction {
	return b.NextFunc
}

func (b *BaseFunction) Prev() fiface.IFunction {
	return b.PrevFunc

}

func (b *BaseFunction) SetN(f fiface.IFunction) {
	b.NextFunc = f
}

func (b *BaseFunction) SetP(f fiface.IFunction) {
	b.PrevFunc = f
}
