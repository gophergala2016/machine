package machine

import "golang.org/x/net/context"

//State this is the state fucntion signiture.
type State func(context.Context, Transitioner)

//Joiner I needed an interface so I can work with async operation.
type Joiner interface {
	//Wait sometimes you need to block and wait unitl the Joiner operation
	//completes before you move on. So this method should be blocking operation.
	Wait(timeout int64)
}

//Transitioner is an interface which can abstarcted common oeprations which
//state machine needed to operate.
type Transitioner interface {
	//Next gets a next state. Once you call it you should not do anything on
	//caller. becuase state machine has moved on to next state.
	Next(State)
	//Fork is a special method which I found it useful because it runs each pass
	//states as a initial state of a brand new state machine. this makes it useful
	//when a state contains sub states which needs to be called in parallel.
	//it returns a Joiner which helps called to wait until all those state machine
	//complete their operations.
	//context.Context is being used to have a share context between thoses state
	//machine.
	Fork(context.Context, ...State) Joiner
	//Done is method that tells state machien that we are Done and no more states
	//are going to be processed.
	Done()
}

//Machine is base blocking start point.
type Machine interface {
	//Run is the start point of state machine. context.Context is a share context
	//between all those states inside state machine. the second argument is an
	//initial state of our state machine.
	Run(context.Context, State) Joiner
}
