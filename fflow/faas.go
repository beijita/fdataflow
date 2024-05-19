package fflow

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

//type FaaS func(ctx context.Context, flow fiface.IFlow) error
type FaaS interface{}

type FaaSDesc struct {
	FnName    string
	f         interface{}
	fName     string
	ArgsType  []reflect.Type
	ArgNum    int
	FuncType  reflect.Type
	FuncValue reflect.Value
}

func NewFaaSDesc(fnName string, f FaaS) (*FaaSDesc, error) {
	funcValue := reflect.ValueOf(f)
	funcType := funcValue.Type()
	if !isFuncType(funcType) {
		return nil, fmt.Errorf("")
	}
	if funcType.NumOut() != 1 || funcType.Out(0) != reflect.TypeOf((*error)(nil)).Elem() {
		return nil, fmt.Errorf("")
	}
	argsType := make([]reflect.Type, funcType.NumIn())
	fullName := runtime.FuncForPC(funcValue.Pointer()).Name()
	var containsDataFlow, containsCtx bool
	for i := 0; i < funcType.NumIn(); i++ {
		paramType := funcType.In(i)
		if isFlowType(paramType) {
			containsDataFlow = true
		} else if isContextType(paramType) {
			containsCtx = true
		} else if isSliceType(paramType) {
			itemType := paramType.Elem()
			if itemType.Kind() == reflect.Ptr {
				itemType = itemType.Elem()
			}
		}
		argsType[i] = paramType
	}
	if !containsDataFlow {
		return nil, fmt.Errorf("")
	}
	if !containsCtx {
		return nil, fmt.Errorf("")
	}
	return &FaaSDesc{
		FnName:    fnName,
		f:         f,
		fName:     fullName,
		ArgsType:  argsType,
		ArgNum:    len(argsType),
		FuncType:  funcType,
		FuncValue: funcValue,
	}, nil
}

func isFuncType(paramType reflect.Type) bool {
	return paramType.Kind() == reflect.Func
}

func isFlowType(paramType reflect.Type) bool {
	flowType := reflect.TypeOf((*DataFlow)(nil))
	return paramType.Implements(flowType)
}

func isContextType(paramType reflect.Type) bool {
	typeName := paramType.Name()
	return strings.Contains(typeName, "Context")
}

func isSliceType(paramType reflect.Type) bool {
	return paramType.Kind() == reflect.Slice
}
