package machine

import (
	"sync"
	"time"

	"golang.org/x/net/context"
)

type localDone chan struct{}

func (l localDone) Wait(timeout int64) {
	if timeout > 0 {
		select {
		case <-l:
		case <-time.After(time.Duration(timeout)):
		}
	} else {
		<-l
	}
}

type localStateTransition chan State

func (ls localStateTransition) Next(state State) {
	ls <- state
}

func (ls localStateTransition) Fork(ctx context.Context, states ...State) Joiner {
	var wg sync.WaitGroup
	wg.Add(len(states))

	var done localDone
	done = make(chan struct{})

	for _, state := range states {
		go func(initialState State) {
			defer wg.Done()

			NewLocalMachine().
				Run(ctx, initialState).
				Wait(0)
		}(state)
	}

	go func() {
		defer close(done)
		wg.Wait()
	}()

	return done
}

func (ls localStateTransition) Done() {
	close(ls)
}

type localMachine struct {
	done         localDone
	transitioner localStateTransition
}

func (lm *localMachine) Run(ctx context.Context, state State) Joiner {
	go func() {
		ok := true

		for ok {
			select {
			case state, ok = <-lm.transitioner:
				if ok {
					state(ctx, lm.transitioner)
				}
			case _, ok = <-ctx.Done():
			}
		}

		defer close(lm.done)
	}()

	state(ctx, lm.transitioner)

	return lm.done
}

func NewLocalMachine() Machine {
	return &localMachine{
		done:         make(chan struct{}),
		transitioner: make(chan State, 1),
	}
}
