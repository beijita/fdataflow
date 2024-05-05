package log

import "context"

type DefaultLog struct {
}

func (log *DefaultLog) InfoFc(ctx context.Context, str string, v ...interface{}) {
}

func (log *DefaultLog) WarnFc(ctx context.Context, str string, v ...interface{}) {
}

func (log *DefaultLog) ErrorFc(ctx context.Context, str string, v ...interface{}) {
}

func (d DefaultLog) DebugFc(ctx context.Context, str string, v ...interface{}) {
}

func (log *DefaultLog) InfoF(str string, v ...interface{}) {
}

func (log *DefaultLog) WarnF(str string, v ...interface{}) {
}

func (log *DefaultLog) ErrorF(str string, v ...interface{}) {
}
