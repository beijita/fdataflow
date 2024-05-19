package fflow

import (
	"context"
	"errors"
	"fmt"
	"github.com/fdataflow/fcommon"
	"github.com/fdataflow/fiface"
	"log"
	"reflect"
	"sync"
)

var _poolOnce sync.Once

type DataFlowPool struct {
	funcRouter funcRouter
	funcLock   sync.RWMutex
	flowRouter flowRouter
	flowLock   sync.RWMutex

	connInitRouter fiface.ConnInitRouter
	cirLock        sync.RWMutex
	cTree          connTree
	ctLock         sync.RWMutex
	connectorMap   map[string]fiface.IConnector
}

var _pool *DataFlowPool

func Pool() *DataFlowPool {
	_poolOnce.Do(func() {
		_pool = new(DataFlowPool)
		_pool.funcRouter = make(funcRouter)
		_pool.flowRouter = make(flowRouter)
		_pool.connInitRouter = make(fiface.ConnInitRouter)
		_pool.cTree = make(connTree)
		_pool.connectorMap = make(map[string]fiface.IConnector)
	})
	return _pool
}

func (pool *DataFlowPool) GetFlows() []fiface.IFlow {
	pool.flowLock.RLock()
	defer pool.flowLock.RUnlock()
	var flows []fiface.IFlow
	for _, flow := range pool.flowRouter {
		flows = append(flows, flow)
	}
	return flows
}

func (pool *DataFlowPool) CaaSInit(name string, conn fiface.ConnInit) {
	pool.cirLock.Lock()
	defer pool.cirLock.Unlock()
	if _, ok := pool.connInitRouter[name]; ok {
		pool.connInitRouter[name] = conn
	} else {
		panic(fmt.Errorf(" CaaSInit is repeat name=%v", name))
	}
	log.Println(fmt.Sprintf(" CaaSInit success! name = %v ", name))
}

func (pool *DataFlowPool) CallConnInit(conn fiface.IConnector) error {
	pool.cirLock.RLock()
	defer pool.cirLock.RUnlock()

	init, ok := pool.connInitRouter[conn.GetName()]
	if !ok {
		panic(fmt.Errorf(" not found! init connector cname = %s", conn.GetName()))
	}
	return init(conn)
}

func (pool *DataFlowPool) CallConnector(ctx context.Context, flow fiface.IFlow, conn fiface.IConnector, args interface{}) error {
	f := flow.GetThisFUnction()
	fConf := f.GetConfig()
	mode := fcommon.DataFlowMode(fConf.FMode)
	if callback, ok := pool.cTree[conn.GetName()][mode][fConf.FName]; ok {
		return callback(ctx, conn, flow, args)
	}
	return fmt.Errorf("not fond connector! connName=%v, funcName=%v, mode=%v ", conn.GetName(), fConf.FName, mode)
}

func (pool *DataFlowPool) CaaS(connName, funcName string, mode fcommon.DataFlowMode, c fiface.CaaS) {
	pool.ctLock.Lock()
	defer pool.ctLock.Unlock()

	if _, ok := pool.cTree[connName]; !ok {
		pool.cTree[connName] = make(connSL)
		pool.cTree[connName][fcommon.Save] = make(fiface.ConnFuncRouter)
		pool.cTree[connName][fcommon.Load] = make(fiface.ConnFuncRouter)
	}
	if _, ok := pool.cTree[connName][mode][funcName]; !ok {
		pool.cTree[connName][mode][funcName] = c
	} else {
		panic(fmt.Errorf("CaaS repeat! connName=%v, funcName=%v, mode=%v", connName, funcName, mode))
	}
}

func (pool *DataFlowPool) AddFlow(name string, flow fiface.IFlow) {
	pool.flowLock.Lock()
	defer pool.flowLock.Unlock()

	if _, ok := pool.flowRouter[name]; ok {
		log.Println("error pool.flowRouter exist flow name=", name)
	} else {
		pool.flowRouter[name] = flow
	}
}

func (pool *DataFlowPool) GetFlow(name string) fiface.IFlow {
	pool.flowLock.RLock()
	defer pool.flowLock.RUnlock()
	return pool.flowRouter[name]
}

func (pool *DataFlowPool) FaaS(name string, f FaaS) {
	pool.funcLock.Lock()
	defer pool.funcLock.Unlock()
	faaSDesc, err := NewFaaSDesc(name, f)
	if err != nil {
		return
	}
	if _, ok := pool.funcRouter[name]; ok {
		log.Println("error pool.flowRouter exist flow name=", name)
		panic(fmt.Sprintf("error pool.flowRouter exist flow name= %s ", name))
	} else {
		pool.funcRouter[name] = faaSDesc
	}
}

func (pool *DataFlowPool) CallFunction(ctx context.Context, name string, flow fiface.IFlow) error {
	pool.funcLock.Lock()
	defer pool.funcLock.Unlock()
	if f, ok := pool.funcRouter[name]; ok {
		faaSDesc := f.(FaaSDesc)
		params := make([]reflect.Value, 0, faaSDesc.ArgNum)
		for _, argType := range faaSDesc.ArgsType {
			if isFlowType(argType) {
				params = append(params, reflect.ValueOf(flow))
			} else if isContextType(argType) {
				params = append(params, reflect.ValueOf(ctx))
			} else if isSliceType(argType) {
				//params = append(params, argType.Elem())
				value, err := faaSDesc.Serialize.UnMarshal(flow.InputData(), argType)
				if err != nil {
				} else {
					params = append(params, value)
					continue
				}
			} else {
				params = append(params, reflect.Zero(argType))
			}
		}
		retValues := faaSDesc.FuncValue.Call(params)
		ret := retValues[0].Interface()
		if ret == nil {
			return nil
		}
		return ret.(error)
	}
	return errors.New("FuncName " + name + " Can not find in NsPool Not Added. ")
}
