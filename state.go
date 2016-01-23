package machine

import "golang.org/x/net/context"

//StateFnResult is a type wraps <-chan StateFn. the main reason of this type is
//to hides and simplified the coding.
type StateFnResult <-chan StateFn

//StateFn every state in the application must have this signiture.
//every state will accept context.Context so they can maintain the state on the
//running file.
type StateFn func(context.Context) StateFnResult

//NextStateFn is a type which hides the chan StateFn.
type NextStateFn chan StateFn

//Next is a method to push the next state to channel. in your code, each state
//might endup going into different state.
func (n NextStateFn) Next(next StateFn) {
	n <- next
}

//Close will close this channel and let everyone that it won't sending any more
//statefn.
func (n NextStateFn) Close() {
	close(n)
}

//Done is a simple type which returns from Rum function.
//you can call wait on this object and it wait until the state machine compeletes
type Done <-chan struct{}

//Wait a very thin wrapper around returning type which blocks a go routine until
//it gets closed.
func (d *Done) Wait() {
	<-(*d)
}

//Run runs the state by providing the start state.
//every state machine must be started by a start state.
func Run(ctx context.Context, start StateFn) Done {
	done := make(chan struct{})

	if ctx == nil {
		ctx = context.Background()
	}

	go func() {
		state, ok := <-start(ctx)

		for ok {
			select {
			case state, ok = <-state(ctx):
			case _, ok = <-ctx.Done():
			}
		}

		close(done)
	}()

	return done
}

//MakeStateFn a utility function to remove a bolierplate of creating
//result channel.
func MakeStateFn(fn func(NextStateFn)) StateFnResult {
	result := make(NextStateFn)
	fn(result)
	return StateFnResult(result)
}
