package fiface

type Action struct {
	Abort     bool
	DataReuse bool
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
	action.Abort = true
}

func ActionDataReuse(action *Action) {
	action.DataReuse = true
}
