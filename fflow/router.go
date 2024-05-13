package fflow

import (
	"context"
	"github.com/fdataflow/fiface"
)

type FaaS func(ctx context.Context, flow fiface.IFlow) error

type funcRouter map[string]FaaS

type flowRouter map[string]fiface.IFlow
