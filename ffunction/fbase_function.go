package ffunction

import (
	"context"
	"github.com/fdataflow/config"
	"github.com/fdataflow/fcommon"
	"github.com/fdataflow/fid"
	"github.com/fdataflow/fiface"
	"sync"
)

type BaseFunction struct {
	ID       string
	Config   *config.FuncConfig
	Flow     fiface.IFlow
	NextFunc fiface.IFunction
	PrevFunc fiface.IFunction

	connector fiface.IConnector
	metaData  map[string]interface{}
	metaLock  sync.RWMutex
}

func (b *BaseFunction) GetMetaData(key string) interface{} {
	b.metaLock.RLock()
	defer b.metaLock.RUnlock()
	return b.metaData[key]
}

func (b *BaseFunction) SetMetaData(key string, value interface{}) {
	b.metaLock.Lock()
	defer b.metaLock.Unlock()
	b.metaData[key] = value
}

func (b *BaseFunction) GetConnector() fiface.IConnector {
	return b.connector
}

func (b *BaseFunction) SetConnector(connector fiface.IConnector) {
	b.connector = connector
}

func NewDataFlowFunction(flow fiface.IFlow, config *config.FuncConfig) fiface.IFunction {
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

func (b *BaseFunction) Call(ctx context.Context, flow fiface.IFlow) error {
	return nil
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

func (b *BaseFunction) SetFlow(f fiface.IFlow) {
	b.Flow = f
}

func (b *BaseFunction) GetFlow() fiface.IFlow {
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
