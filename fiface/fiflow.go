package fiface

import (
	"context"
	"github.com/fdataflow/config"
	"github.com/fdataflow/fcommon"
	"time"
)

type IFlow interface {
	Run(ctx context.Context) error
	Link(fConf *config.FuncConfig, fParam config.FParam) error
	CommitRow(row interface{}) error
	InputData() fcommon.DataFlowRowArr

	GetName() string
	GetThisFUnction() IFunction
	GetThisFuncConf() *config.FuncConfig

	GetConnector() IConnector
	GetConnConf() *config.ConnConfig
	GetFlowConfig() *config.DataFlowConfig
	GetFuncConfigByFuncName(funcName string) *config.FuncConfig

	Next(acts ...ActionFunc) error

	GetCacheData(key string) interface{}
	SetCacheData(key string, value interface{}, expireTime time.Duration)
}
