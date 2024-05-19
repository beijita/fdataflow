package fiface

import (
	"github.com/fdataflow/fcommon"
	"reflect"
)

type ISerialize interface {
	UnMarshal(arr fcommon.DataFlowRowArr, p reflect.Type) (reflect.Value, error)
	Marshal(interface{}) (fcommon.DataFlowRowArr, error)
}
