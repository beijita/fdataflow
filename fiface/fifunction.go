package fiface

import (
	"context"
	"github.com/fdataflow/config"
)

type IFunction interface {
	Call(ctx context.Context, flow IFlow) error
	SetConfig(s *config.FuncConfig)
	GetConfig() *config.FuncConfig
	SetFlow(f IFlow)
	GetFlow() IFlow
	CreateID()
	GetID() string
	GetPreID() string
	GetNextID() string
	Next() IFunction
	Prev() IFunction
	SetN(f IFunction)
	SetP(f IFunction)
	GetConnector() IConnector
	SetConnector(connector IConnector)
	GetMetaData(key string) interface{}
	SetMetaData(key string, value interface{})
}
