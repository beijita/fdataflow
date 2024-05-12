package config

import (
	"fmt"
	"github.com/fdataflow/fcommon"
)

type ConnConfig struct {
	DataFlowType string                `yaml:"data_flow_type" json:"data_flow_type,omitempty"`
	ConnName     string                `yaml:"conn_name" json:"conn_name,omitempty"`
	AddrString   string                `yaml:"addr_string,omitempty"`
	Type         fcommon.ConnectorType `yaml:"type,omitempty"`
	Key          string                `yaml:"key,omitempty"`
	Params       map[string]string     `yaml:"params,omitempty"`
	Load         []string              `yaml:"load,omitempty"`
	Save         []string              `yaml:"save,omitempty"`
}

func NewConnConfig(connName, addr, key string, tType fcommon.ConnectorType, param FParam) *ConnConfig {
	var result ConnConfig
	result.ConnName = connName
	result.AddrString = addr
	result.Key = key
	result.DataFlowType = string(tType)
	result.Params = param
	return &result
}

func (c *ConnConfig) WithFunc(conf *FuncConfig) error {
	switch fcommon.DataFlowMode(conf.FMode) {
	case fcommon.Save:
		c.Save = append(c.Save, conf.FName)
	case fcommon.Load:
		c.Load = append(c.Load, conf.FName)
	default:
		return fmt.Errorf("func binding error! mode=%s", conf.FMode)
	}
	return nil
}
