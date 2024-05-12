package fcommon

type DataFlowRow interface{}

type DataFlowRowArr []DataFlowRow

type DataFlowDataMap map[string]DataFlowRowArr
