package config

import (
	"fmt"
	"github.com/fdataflow/common"
	"testing"
)

func TestNewFuncConfig(t *testing.T) {
	source := DataFlowSource{
		Name:       "直播间",
		MustFields: []string{"room_id", "user_id"},
	}
	funcOption := FuncOption{
		ConnName:      "name1",
		RetryTimes:    3,
		RetryDuration: 1000,
		Params:        FParam{"param1": "value1", "param2": "value2"},
	}
	funcConfig := NewFuncConfig("funcNamefzs", common.Calculate, &source, &funcOption)
	fmt.Println("funcConfig=", funcConfig)
}
