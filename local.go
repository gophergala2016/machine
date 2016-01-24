package state

import "golang.org/x/net/context"

type localDone <-chan struct{}

func (l localDone) Wait(timeout int64) {
	<-l
}

type localStateTransition chan State

func (ls localStateTransition) Next(state State) {
	ls <- state
}

func (ls localStateTransition) Fork(states ...State) Joiner {
	return nil
}

type localMachine struct {
	done         localDone
	transitioner localStateTransition
}

func (lm LocalMachine) Run(ctx context.Context, initialState State) Joiner {
	go func() {
		ok := true

		for ok {
			select {
			case state, ok := <-lm.transitioner:
				if ok {
					state(ctx, lm.transitioner)
				}
			case _, ok = <-ctx.Done():
			}
		}

		defer close(localDone)
	}()

	initialState(ctx, lm.transitioner)

	return lm.done
}

func NewLocalMachine() Machine {
	return &localMachine{
		transitioner: make(chan State, 1),
	}
}
