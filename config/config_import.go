package config

import (
	"fmt"
	"github.com/fdataflow/fcommon"
	"github.com/fdataflow/fflow"
	"github.com/fdataflow/fiface"
	"gopkg.in/yaml.v3"
	"io/fs"
	"io/ioutil"
	"path"
	"path/filepath"
)

type AllConfig struct {
	FlowMap map[string]*DataFlowConfig
	FuncMap map[string]*FuncConfig
	ConnMap map[string]*ConnConfig
}

func InitAllConfig() *AllConfig {
	return &AllConfig{
		FlowMap: make(map[string]*DataFlowConfig),
		FuncMap: make(map[string]*FuncConfig),
		ConnMap: make(map[string]*ConnConfig),
	}
}

func ImportConfigYaml(loadPath string) error {
	walkYaml, err := parseConfigWalkYaml(loadPath)
	if err != nil {
		return err
	}
	for flowName, flowConfig := range walkYaml.FlowMap {
		newFlow := fflow.NewDataFlow(flowConfig)
		for _, fp := range flowConfig.Flows {
			buildFlow(walkYaml, fp, newFlow, flowName)
		}
		fflow.Pool().AddFlow(flowName, newFlow)
	}
	return nil
}

func buildFlow(ac *AllConfig, fp DataFlowFunctionParam, flow fiface.IFlow, name string) error {
	funcConfig, ok := ac.FuncMap[fp.FuncName]
	if !ok {
		return fmt.Errorf("")
	}
	if funcConfig.Option.ConnName != "" {
		connConf, ok := ac.ConnMap[funcConfig.Option.ConnName]
		if !ok {
			return fmt.Errorf("")
		}
		funcConfig.SetConnConf(connConf)
	}
	return flow.Link(funcConfig, fp.Params)
}

func parseConfigWalkYaml(loadPath string) (*AllConfig, error) {
	ac := InitAllConfig()
	err := filepath.Walk(loadPath, func(filePath string, info fs.FileInfo, err error) error {
		if suffix := path.Ext(filePath); suffix != ".yml" && suffix != ".yaml" {
			return nil
		}
		readFile, err := ioutil.ReadFile(filePath)
		if err != nil {
			return err
		}
		confMap := make(map[string]interface{})
		err = yaml.Unmarshal(readFile, confMap)
		if err != nil {
			return err
		}
		if dataFlowType, ok := confMap["dataflow_type"]; !ok {
			return fmt.Errorf("")
		} else {
			switch dataFlowType {
			case fcommon.DataFlowIDTypeFunction:
				return parseFuncConfig(ac, readFile, filePath, dataFlowType)
			case fcommon.DataFlowIDTypeConnector:
				return parseConnConfig(ac, readFile, filePath, dataFlowType)
			case fcommon.DataFlowIDTypeFlow:
				return parseFlowConfig(ac, readFile, filePath, dataFlowType)
			default:
				return fmt.Errorf("")
			}
		}

	})
	if err != nil {
		return nil, err
	}
	return ac, nil
}

func parseFlowConfig(ac *AllConfig, confData []byte, fileName string, dataFlowType interface{}) error {
	var flowConfig DataFlowConfig
	err := yaml.Unmarshal(confData, &flowConfig)
	if err != nil {
		return err
	}
	if fcommon.DataFLowOnOff(flowConfig.Status) == fcommon.Disable {
		return nil
	}
	_, ok := ac.FlowMap[flowConfig.FlowName]
	if ok {
		return fmt.Errorf("flow config is repeat! flowName=%v", flowConfig.FlowName)
	}
	ac.FlowMap[flowConfig.FlowName] = &flowConfig
	return nil
}

func parseFuncConfig(ac *AllConfig, confData []byte, fileName string, dataFlowType interface{}) error {
	var funcConfig FuncConfig
	err := yaml.Unmarshal(confData, &funcConfig)
	if err != nil {
		return err
	}
	_, ok := ac.FuncMap[funcConfig.FName]
	if ok {
		return fmt.Errorf("flow config is repeat! funcConfig.FName=%v", funcConfig.FName)
	}
	ac.FuncMap[funcConfig.FName] = &funcConfig
	return nil
}

func parseConnConfig(ac *AllConfig, confData []byte, fileName string, dataFlowType interface{}) error {
	var connConfig ConnConfig
	err := yaml.Unmarshal(confData, &connConfig)
	if err != nil {
		return err
	}
	_, ok := ac.ConnMap[connConfig.ConnName]
	if ok {
		return fmt.Errorf("flow config is repeat! connConfig.ConnName=%v", connConfig.ConnName)
	}
	ac.ConnMap[connConfig.ConnName] = &connConfig
	return nil
}
