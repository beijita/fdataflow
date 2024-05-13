package fflow

import (
	"context"
	"errors"
	"fmt"
	"github.com/fdataflow/fiface"
	"log"
	"sync"
)

var _poolOnce sync.Once

type DataFlowPool struct {
	funcRouter funcRouter
	funcLock   sync.RWMutex
	flowRouter flowRouter
	flowLock   sync.RWMutex
}

var _pool *DataFlowPool

func Pool() *DataFlowPool {
	_poolOnce.Do(func() {
		_pool = new(DataFlowPool)
		_pool.funcRouter = make(funcRouter)
		_pool.flowRouter = make(flowRouter)
	})
	return _pool
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

	if _, ok := pool.funcRouter[name]; ok {
		log.Println("error pool.flowRouter exist flow name=", name)
		panic(fmt.Sprintf("error pool.flowRouter exist flow name= %s ", name))
	} else {
		pool.funcRouter[name] = f
	}
}

func (pool *DataFlowPool) CallFunction(ctx context.Context, name string, flow fiface.IFlow) error {
	pool.funcLock.Lock()
	defer pool.funcLock.Unlock()
	if f, ok := pool.funcRouter[name]; ok {
		return f(ctx, flow)
	}
	return errors.New("FuncName " + name + " Can not find in NsPool Not Added. ")
}
