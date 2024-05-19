package fiface

import (
	"context"
	"github.com/fdataflow/config"
)

type IConnector interface {
	Init() error
	Call(ctx context.Context, flow IFlow, args interface{}) error
	GetID() string
	GetName() string
	GetConfig() *config.ConnConfig

	GetMetaData(key string) interface{}
	SetMetaData(key string, value interface{})
}

type ConnInit func(conn IConnector) error

type ConnInitRouter map[string]ConnInit

type CaaS func(context.Context, IConnector, IFlow, interface{}) error

type ConnFuncRouter map[string]CaaS
