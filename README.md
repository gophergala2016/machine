# machine

## introduction

machine is a go library to write state machine. The idea came to me when I watched [Rob Pike's Lexical Scanning in Go](https://www.youtube.com/watch?v=HxaD_trXwRE). So the basic idea is that you create a State Machine and start with initial state. Every state either generates a new state or stops the system. There are no errors coming from state machine. Since an error can lead to another state.

I tried to provide a simple and powerful interfaces so one can easily extend the state machine into next level.


## details

there are only 3 interfaces which help me to abstract the complexity of state machine.

```go
type Joiner interface {
	Wait(timeout int64)
}
```

```go
type Transitioner interface {
	Next(State)
	Fork(context.Context, ...State) Joiner
	Done()
}
```

```go
type Machine interface {
	Run(context.Context, State) Joiner
}
```

And also the main type of my state machine `State` which is a simple function that accepts `context.Context` and `Transitioner`.

```go
type State func(context.Context, Transitioner)
```

by just having these 4 things, we can easily build any state machines.


### Joiner
`Joiner` has a single method. it should be block until either timeout passes or machine run finishes and processes all states.

### Transitioner
`Transitioner` is a blue print of state pipeline. It provides state to go to another state.

`Fork` is special case which creates a brand new state machine for each states that being pass into it and joins them at the end by providing a single `Joiner` object.

`Done` is being used to tell the state machine that state has hit the end of the process.

### Machine
`Machine` is the interface for the state machine engine. At the moment I have provided a simple local state machine. but consider this, you can extend it and make it as a distributed state machine.

`Run` starts the engine by providing `context.Context` and first initial `State`. `context.Context` is being used as a way to share context with all states inside that particular state machine and ability to cancel and stop state machine in the middle of process.

## Next

- implement a distributed version of state machine using this interface
- implement more examples

## Continue

The project will be continued at github.com/alinz/machine
