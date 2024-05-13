package fiface

import (
	"context"
	"github.com/fdataflow/config"
	"github.com/fdataflow/fcommon"
)

type IFlow interface {
	Run(ctx context.Context) error
	Link(fConf *config.FuncConfig, fParam config.FParam) error
	CommitRow(row interface{}) error
	InputData() fcommon.DataFlowRowArr

	GetName() string
	GetThisFUnction() IFunction
	GetThisFuncConf() *config.FuncConfig
}
