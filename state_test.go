package machine_test

import (
	"fmt"
	"testing"

	"github.com/alinz/machine"
	"golang.org/x/net/context"
)

func state1(ctx context.Context, pipe machine.StateFnPipein) {
	fmt.Println("state1")
	pipe.Next(state2)
}

func state2(ctx context.Context, pipe machine.StateFnPipein) {
	fmt.Println("state2")
	pipe.Close()
}

func TestSimpleState(t *testing.T) {
	<-machine.Run(nil, state1)

	fmt.Println("Done")
}
