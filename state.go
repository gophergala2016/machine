package machine

import "golang.org/x/net/context"

type StateFnPipein chan<- StateFn

func (p StateFnPipein) Next(fn StateFn) {
	p <- fn
}

func (p StateFnPipein) Close() {
	close(p)
}

//StateFn every state in the application must have this signiture.
//every state will accept context.Context so they can maintain the state on the
//running file.
//rememeber, error should not leakout. since error in go treated as type
//it can be used to go into different state. So technically, you do not need
//to return an error.
type StateFn func(context.Context, StateFnPipein)

//Run runs the state by providing the start state.
//every state machine must be started by a start state.
func Run(ctx context.Context, state StateFn) <-chan struct{} {
	done := make(chan struct{})
	pipe := make(chan StateFn, 1)

	//if context passes to run is not initialized, then we assing the background one.
	if ctx == nil {
		ctx = context.Background()
	}

	//runs the state machine in a gorotine.
	go func() {
		ok := true
		//if it's ok, then it goes into infinite of calling the next state.
		//context is being use here to first share values between each state and
		//also more importantly if context is timeout or cancel, it stop the
		//forloop process.
		for ok {
			select {
			//we are waiting for state to return a new stat via state channel.
			case state, ok = <-pipe:
				if ok {
					state(ctx, pipe)
				}
			//or context gets canceled or timeout.
			case _, ok = <-ctx.Done():

			}
		}

		//this is a signal value which will be use to notify that the entore state
		//machine is completed.
		defer close(done)
	}()

	//it start off by calling start state to kick off the state machine
	state(ctx, pipe)

	return done
}
