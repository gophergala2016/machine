package machine_test

import (
	"fmt"
	"testing"

	"github.com/gophergala2016/machine"
	"golang.org/x/net/context"
)

func state1(ctx context.Context, transitoner machine.Transitioner) {
	fmt.Println("this is state 1")
	transitoner.Next(state2)
}

func state2(ctx context.Context, transitoner machine.Transitioner) {
	fmt.Println("this is state 2")
	transitoner.Next(state3)
}

func state3(ctx context.Context, transitoner machine.Transitioner) {
	fmt.Println("this is state 3")
	transitoner.Done()
}

func TestSimpleState(t *testing.T) {
	localMachine := machine.NewLocalMachine()

	ctx := context.Background()

	localMachine.Run(ctx, state1).Wait(1)
}
