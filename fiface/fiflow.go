package fiface

import (
	"context"
	"github.com/fdataflow/config"
)

type IFlow interface {
	Run(ctx context.Context) error
	Link(fConf *config.FuncConfig, fParam config.FParam) error
	CommitRow(row interface{}) error
}
