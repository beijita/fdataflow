package fiface

import (
	"context"
	"github.com/fdataflow/config"
)

type Flow interface {
	Run(ctx context.Context) error
	Link(fConf *config.FuncConfig, fParam config.FParam) error
}
