package fid

import (
	"github.com/fdataflow/fcommon"
	"github.com/google/uuid"
	"strings"
)

func DataFlowID(prefix ...string) (dataFlowID string) {
	idStr := strings.Replace(uuid.New().String(), "-", "", -1)
	return formatDataflowID(idStr, prefix...)
}

func formatDataflowID(IDStr string, prefix ...string) string {
	var dataflowID string
	for _, p := range prefix {
		dataflowID += p
		dataflowID += fcommon.DataFlowIDJoinChar
	}
	return dataflowID + IDStr
}
