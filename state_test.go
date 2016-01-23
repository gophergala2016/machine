package state_test

import (
	"testing"

	"bitbucket.org/alinz/pigeon/state"
	"golang.org/x/net/context"
)

func state1(ctx context.Context) <-chan state.StateFn {
	return state.MakeStateFn(func(state state.NextStateFn) {
		state.Next(state2)
	})
}

func state2(ctx context.Context) <-chan state.StateFn {
	return state.MakeStateFn(func(state state.NextStateFn) {

	})
}

func TestSimpleState(t *testing.T) {

}
