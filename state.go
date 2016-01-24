package machine

import "golang.org/x/net/context"

type State func(context.Context, Transitioner)

type Joiner interface {
	Wait(timeout int64)
}

type Transitioner interface {
	Next(State)
	Fork(...State) Joiner
	Done()
}

type Machine interface {
	Run(context.Context, State) Joiner
}
