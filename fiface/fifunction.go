package fiface

import (
	"context"
	"github.com/fdataflow/config"
)

type IFunction interface {
	Call(ctx context.Context, flow Flow) error
	SetConfig(s *config.FuncConfig)
	GetConfig() *config.FuncConfig
	SetFlow(f Flow)
	GetFlow() Flow
	CreateID()
	GetID() string
	GetPreID() string
	GetNextID() string
	Next() IFunction
	Prev() IFunction
	SetN(f IFunction)
	SetP(f IFunction)
}
