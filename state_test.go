package machine_test

import (
	"testing"

	"github.com/alinz/machine"
	"golang.org/x/net/context"
)

func state1(ctx context.Context) <-chan machine.StateFn {
	return machine.MakeStateFn(func(pipe chan machine.StateFn) {

	})
}

func state2(ctx context.Context) <-chan machine.StateFn {
	return machine.MakeStateFn(func(pipe chan machine.StateFn) {

	})
}

func TestSimpleState(t *testing.T) {

}
