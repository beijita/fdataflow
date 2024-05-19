package fflow

import (
	"github.com/fdataflow/fcommon"
	"github.com/fdataflow/fiface"
)

type funcRouter map[string]FaaS

type flowRouter map[string]fiface.IFlow

type connSL map[fcommon.DataFlowMode]fiface.ConnFuncRouter

type connTree map[string]connSL
