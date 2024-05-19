package fconn

import (
	"context"
	"github.com/fdataflow/config"
	"github.com/fdataflow/fcommon"
	"github.com/fdataflow/fflow"
	"github.com/fdataflow/fid"
	"github.com/fdataflow/fiface"
	"sync"
)

type Connector struct {
	CID      string
	ConnName string
	Conf     *config.ConnConfig
	onceInit sync.Once
}

func NewConnector(conf *config.ConnConfig) *Connector {
	return &Connector{
		CID:      fid.DataFlowID(fcommon.DataFlowIDTypeConnector),
		ConnName: conf.ConnName,
		Conf:     conf,
		onceInit: sync.Once{},
	}
}

func (c *Connector) Init() error {
	var err error
	c.onceInit.Do(func() {
		err = fflow.Pool().CallConnInit(c)
	})
	return err
}

func (c *Connector) Call(ctx context.Context, flow fiface.IFlow, args interface{}) error {
	return fflow.Pool().CallConnector(ctx, flow, c, args)
}

func (c *Connector) GetID() string {
	return c.CID
}

func (c *Connector) GetName() string {
	return c.ConnName
}

func (c *Connector) GetConfig() *config.ConnConfig {
	return c.Conf
}
