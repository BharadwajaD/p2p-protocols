package peers

type State struct{
    data map[string] any
}

func NewState() State {
    return State{
        data: map[string] any{},
    }
}

func (s *State) AddState(desc string, state any) {
	s.data[desc] = state
}

type func_t func(...any) any

func (s *State) Run(fn func_t, args ...any) any{
    return fn(args)
}
