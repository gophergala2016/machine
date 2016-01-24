package machine_test

import (
	"fmt"
	"testing"

	"github.com/gophergala2016/machine"
	"golang.org/x/net/context"
)

func state1(ctx context.Context, transitioner machine.Transitioner) {
	fmt.Println("this is state 1")
	transitioner.Next(state2)
}

func state2(ctx context.Context, transitioner machine.Transitioner) {
	fmt.Println("this is state 2")
	transitioner.Next(state3)
}

func state3(ctx context.Context, transitioner machine.Transitioner) {
	fmt.Println("this is state 3")
	transitioner.Done()
}

func TestLocalStateRun(t *testing.T) {
	localMachine := machine.NewLocalMachine()

	ctx := context.Background()

	localMachine.Run(ctx, state1).Wait(0)
}

func state4(ctx context.Context, transitioner machine.Transitioner) {
	fmt.Println("state 4")

	transitioner.Fork(ctx, state1, state2, state3).Wait(0)

	transitioner.Done()
}

func TestLocalStateFork(t *testing.T) {
	localMachine := machine.NewLocalMachine()

	ctx := context.Background()

	localMachine.Run(ctx, state4)
}
