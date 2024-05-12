package fflow

import (
	"context"
	"github.com/fdataflow/config"
	"github.com/fdataflow/fcommon"
	"log"
	"testing"
)

func TestDataFlow_Run(t *testing.T) {
	roomSource := config.DataFlowSource{
		Name:       "直播间",
		MustFields: []string{"room_id", "room_name"},
	}
	anchorSource := config.DataFlowSource{
		Name:       "主播",
		MustFields: []string{"anchor_id", "anchor_name"},
	}

	roomFuncConfig := config.NewFuncConfig("roomFunc", fcommon.Calculate, &roomSource, nil)
	if roomFuncConfig == nil {
		log.Println(" roomFuncConfig == nil ")
		return
	}
	anchorFuncConfig := config.NewFuncConfig("anchorFunc", fcommon.Calculate, &anchorSource, nil)
	if anchorFuncConfig == nil {
		log.Println(" anchorFuncConfig == nil ")
		return
	}

	flowConfig := config.NewFlowConfig("firstFlow", fcommon.Enable)
	flow := NewDataFlow(flowConfig)
	err := flow.Link(anchorFuncConfig, nil)
	if err != nil {
		log.Println(" flow.Link(anchorFuncConfig err= ", err)
		return
	}
	err = flow.Link(roomFuncConfig, nil)
	if err != nil {
		log.Println(" flow.Link(roomFuncConfig err= ", err)
		return
	}
	err = flow.Run(context.Background())
	if err != nil {
		log.Println(" flow.Run err= ", err)
		return
	}
	log.Println("success !!!")
}
