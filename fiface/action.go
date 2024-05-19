package fiface

type Action struct {
	AbortFlag          bool
	ForceEntryNextFlag bool
	DataReuseFlag      bool
	JumpFunc           string
}

type ActionFunc func(ops *Action)

func LoadActions(acts []ActionFunc) Action {
	var a Action
	if len(acts) == 0 {
		return a
	}
	for _, act := range acts {
		act(&a)
	}
	return a
}

func ActionAbort(action *Action) {
	action.AbortFlag = true
}

func ActionDataReuse(action *Action) {
	action.DataReuseFlag = true
}

func ActionForceEntryNext(action *Action) {
	action.ForceEntryNextFlag = true
}

func ActionJumpFunc(funcName string) ActionFunc {
	return func(act *Action) {
		act.JumpFunc = funcName
	}
}
