package machine_test

import (
	"fmt"
	"testing"

	"github.com/alinz/machine"
	"golang.org/x/net/context"
)

func state1(ctx context.Context) machine.StateFnPipeout {
	fn := machine.MakeStateFn(func(pipe machine.StateFnPipein) {
		fmt.Println("state1")
		pipe.Next(state2)
	})

	fmt.Println("Started")

	return fn
}

func state2(ctx context.Context) machine.StateFnPipeout {
	return machine.MakeStateFn(func(pipe machine.StateFnPipein) {
		fmt.Println("state2")
		pipe.Close()
	})
}

func TestSimpleState(t *testing.T) {
	<-machine.Run(nil, state1)

	fmt.Println("Done")
}
