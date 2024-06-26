package config

import (
	"fmt"
	"github.com/fdataflow/fcommon"
)

type FParam map[string]string

type DataFlowSource struct {
	Name       string   `yaml:"name" json:"name,omitempty"`
	MustFields []string `yaml:"must_fields" json:"must_fields,omitempty"` // source必传字段
}

type FuncOption struct {
	ConnName      string `yaml:"conn_name" json:"conn_name,omitempty"`
	RetryTimes    int    `yaml:"retry_times" json:"retry_times,omitempty"`
	RetryDuration int    `yaml:"retry_duration" json:"retry_duration,omitempty"`
	Params        FParam `yaml:"params" json:"params,omitempty"`
}

type FuncConfig struct {
	DataFlowType string         `yaml:"data_flow_type" json:"data_flow_type,omitempty"`
	FName        string         `yaml:"f_name" json:"f_name,omitempty"`
	FMode        string         `yaml:"f_mode" json:"f_mode,omitempty"`
	Source       DataFlowSource `yaml:"source" json:"source"`
	Option       FuncOption     `yaml:"option" json:"option"`
	connConf     *ConnConfig    `yaml:"conn_conf"`
}

func (fc *FuncConfig) GetConnConf() (*ConnConfig, error) {
	if fc.connConf == nil {
		return nil, fmt.Errorf(" GetConnConf error! connConf == nil ")
	}
	return fc.connConf, nil
}

func (fc *FuncConfig) SetConnConf(connConf *ConnConfig) error {
	if connConf == nil {
		panic(fmt.Errorf(" SetConnConf error! connConf == nil"))
	}
	fc.connConf = connConf
	return nil
}

func NewFuncConfig(funcName string, mode fcommon.DataFlowMode, source *DataFlowSource, option *FuncOption) *FuncConfig {
	var config FuncConfig
	config.FName = funcName
	if source == nil {
		return nil
	}
	config.Source = *source
	config.FMode = string(mode)
	if mode == fcommon.Save || mode == fcommon.Load {
		if option == nil {
			return nil
		} else if option.ConnName == "" {
			return nil
		}
	}
	if option != nil {
		config.Option = *option
	}
	return &config
}
