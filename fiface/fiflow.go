package fiface

import (
	"context"
	"github.com/fdataflow/config"
)

type Flow interface {
	Run(ctx context.Context)
	Link(fConf *config.FuncConfig)
}
