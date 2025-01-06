package cli_test

import (
	"maps"
	"testing"

	"github.com/cl1ckname/cdf/internal/cli"
	"github.com/cl1ckname/cdf/internal/handler"
)

func TestCli(t *testing.T) {
	t.Parallel()

	code := handler.CodeAdd
	command := "add"
	program := "cdf"
	path := "/home"
	kwarg := "--tmp=/tmp/cdf-2127"

	tests := []struct {
		name  string
		args  []string
		call  handler.Call
		error error
	}{
		{
			name: "only cmd",
			args: []string{program, command},
			call: handler.Call{
				Code: code,
			},
			error: nil,
		},
		{
			name:  "no program",
			error: cli.ErrInvalidArgs,
		},
		{
			name: "one argument",
			args: []string{program, command, path},
			call: handler.Call{
				Code: code,
				Args: []string{path},
			},
			error: nil,
		},
		{
			name: "one kwarg",
			args: []string{program, command, kwarg},
			call: handler.Call{
				Code:   code,
				Kwargs: handler.Kwargs{"tmp": "/tmp/cdf-2127"},
			},
		},
		{
			name: "arg and kwarg",
			args: []string{program, command, path, kwarg},
			call: handler.Call{
				Code:   code,
				Kwargs: handler.Kwargs{"tmp": "/tmp/cdf-2127"},
				Args:   handler.Args{path},
			},
		},
		{
			name: "arg and kwarg reverse",
			args: []string{program, command, kwarg, path},
			call: handler.Call{
				Code:   code,
				Kwargs: handler.Kwargs{"tmp": "/tmp/cdf-2127"},
				Args:   handler.Args{path},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			call, err := cli.ParseCall(test.args)
			if err != nil {
				if test.error == nil {
					t.Fatal("unexpected error")
				}
				return
			}
			if test.error != nil {
				t.Fatal("expected error")
			}
			if expected, actual := test.call.Code, call.Code; expected != actual {
				t.Fatalf("calls code difference: %d vs %d", expected, actual)
			}
			if expected, actual := test.call.Args, call.Args; !arrayEq(expected, actual) {
				t.Fatalf("args difference: %v vs %v", expected, actual)
			}
			if expected, actual := test.call.Kwargs, call.Kwargs; !maps.Equal(expected, actual) {
				t.Fatalf("kwargs differencs: %v vs %v", expected, actual)
			}
		})
	}
}

func arrayEq[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
