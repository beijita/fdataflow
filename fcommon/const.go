package fcommon

import "time"

const (
	DataFlowIDTypeFlow      = "flow"
	DataFlowIDTypeConnector = "conn"
	DataFlowIDTypeFunction  = "func"
	DataFlowIDTypeGlobal    = "global"
	DataFlowIDJoinChar      = "_"
)

const (
	FunctionIDFirstVirtual = "FunctionIDFirstVirtual"
	FunctionIDLastVirtual  = "FunctionIDLastVirtual"
)

type DataFlowMode string

const (
	Verify    DataFlowMode = "Verify"
	Save      DataFlowMode = "Save"
	Load      DataFlowMode = "Load"
	Calculate DataFlowMode = "Calculate"
	Expand    DataFlowMode = "Expand"
)

type DataFLowOnOff int

const (
	Enable  DataFLowOnOff = 1
	Disable DataFLowOnOff = 0
)

type ConnectorType string

const (
	REDIS ConnectorType = "redis"
	MYSQL ConnectorType = "MySQL"
	KAFKA ConnectorType = "kafka"
	TIDB  ConnectorType = "tidb"
	ES    ConnectorType = "es"
)

// cache
const (
	DefaultFlowCacheCleanUp               = 5 * time.Minute
	DefaultExpiration       time.Duration = 0
)

// metrics
const METRICS_ROUTE = "/metrics"
