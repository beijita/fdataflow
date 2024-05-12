package config

import "github.com/fdataflow/fcommon"

type DataFlowFunctionParam struct {
	FuncName string `yaml:"func_name"`
	Params   FParam `yaml:"params"`
}

type DataFlowConfig struct {
	DataType string                  `yaml:"data_type"`
	Status   int                     `yaml:"status"`
	FlowName string                  `yaml:"flow_name"`
	Flows    []DataFlowFunctionParam `yaml:"flows"`
}

func NewFlowConfig(flowName string, enable fcommon.DataFLowOnOff) *DataFlowConfig {
	var config DataFlowConfig
	config.FlowName = flowName
	config.Flows = make([]DataFlowFunctionParam, 0)
	config.Status = int(enable)
	return &config
}

func (fc *DataFlowConfig) AppendFunctionConfig(params DataFlowFunctionParam) {
	fc.Flows = append(fc.Flows, params)
}
