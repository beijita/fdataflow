package log

import "context"

type FlowLogger interface {
	InfoFc(ctx context.Context, str string, v ...interface{})
	WarnFc(ctx context.Context, str string, v ...interface{})
	ErrorFc(ctx context.Context, str string, v ...interface{})
	DebugFc(ctx context.Context, str string, v ...interface{})

	InfoF(str string, v ...interface{})
	WarnF(str string, v ...interface{})
	ErrorF(str string, v ...interface{})
}

var dataFlowLog FlowLogger

func SetLogger(newLog FlowLogger) {
	dataFlowLog = newLog
}

func Logger() FlowLogger {
	return dataFlowLog
}
