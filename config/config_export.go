package config

import (
	"fmt"
	"github.com/fdataflow/fcommon"
	"github.com/fdataflow/fiface"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func ExportConfigYaml(flow fiface.IFlow, savePath string) error {
	data, err := yaml.Marshal(flow.GetFlowConfig())
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(savePath+fcommon.DataFlowIDTypeFlow+"-"+flow.GetName()+".yaml", data, 0644)
	if err != nil {
		return err
	}
	flowConfig := flow.GetFlowConfig()
	if flowConfig == nil {
		return fmt.Errorf("")
	}
	for _, fp := range flowConfig.Flows {
		fConf := flow.GetFuncConfigByFuncName(fp.FuncName)
		if fConf == nil {
			return fmt.Errorf("")
		}
		fcData, err := yaml.Marshal(fConf)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(savePath+fcommon.DataFlowIDTypeFunction+"-"+fp.FuncName+".yaml", fcData, 0644)
		if err != nil {
			return err
		}
		if fConf.Option.ConnName != "" {
			connConf, err := fConf.GetConnConf()
			if err != nil {
				return err
			}
			connData, err := yaml.Marshal(connConf)
			if err != nil {
				return err
			}
			err = ioutil.WriteFile(savePath+fcommon.DataFlowIDTypeConnector+"-"+connConf.ConnName+".yaml", connData, 0644)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
