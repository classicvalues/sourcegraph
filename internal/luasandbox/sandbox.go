package luasandbox

import (
	"context"
	"sync"
	"time"

	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"

	"github.com/sourcegraph/sourcegraph/internal/observation"
)

type Sandbox struct {
	// note: operations around vm state are kept isolated in
	// the run function, which ensures mutual access with the
	// mutex.
	state *lua.LState
	m     sync.Mutex

	operations *operations
}

// Close releases resources occupied by the underlying Lua VM.
// No calls to the sandbox should be made after closing it.
func (s *Sandbox) Close() {
	s.state.Close()
}

// RunScript runs the given Lua script text in the sandbox.
func (s *Sandbox) RunScript(ctx context.Context, opts RunOptions, script string) (retValue lua.LValue, err error) {
	ctx, endObservation := s.operations.runScript.With(ctx, &err, observation.Args{})
	defer endObservation(1, observation.Args{})

	f := func(ctx context.Context, state *lua.LState) error {
		if err := state.DoString(script); err != nil {
			return err
		}

		retValue = state.Get(lua.MultRet)
		return nil
	}
	err = s.RunGoCallback(ctx, opts, f)
	return
}

// Call invokes the given function bound to this sandbox within the sandbox.
func (s *Sandbox) Call(ctx context.Context, opts RunOptions, luaFunction *lua.LFunction, args ...interface{}) (retValue lua.LValue, err error) {
	ctx, endObservation := s.operations.call.With(ctx, &err, observation.Args{})
	defer endObservation(1, observation.Args{})

	f := func(ctx context.Context, state *lua.LState) error {
		state.Push(luaFunction)
		for _, arg := range args {
			state.Push(luar.New(s.state, arg))
		}

		if err := state.PCall(len(args), lua.MultRet, nil); err != nil {
			return err
		}

		retValue = state.Get(lua.MultRet)
		return nil
	}
	err = s.RunGoCallback(ctx, opts, f)
	return
}

// CallGenerator invokes the given coroutine bound to this sandbox within the sandbox.
// Each yield from the coroutine will be collected in the output slide and returned to
// the caller. This method does not pass values back into the coroutine when resuming
// execution.
func (s *Sandbox) CallGenerator(ctx context.Context, opts RunOptions, luaFunction *lua.LFunction, args ...interface{}) (retValues []lua.LValue, err error) {
	ctx, endObservation := s.operations.callGenerator.With(ctx, &err, observation.Args{})
	defer endObservation(1, observation.Args{})

	f := func(ctx context.Context, state *lua.LState) error {
		luaArgs := make([]lua.LValue, 0, len(args))
		for _, arg := range args {
			luaArgs = append(luaArgs, luar.New(s.state, arg))
		}

		co, _ := state.NewThread()

	loop:
		for {
			state, err, yieldedValues := state.Resume(co, luaFunction, luaArgs...)
			switch state {
			case lua.ResumeError:
				return err

			case lua.ResumeYield:
				retValues = append(retValues, yieldedValues...)
				continue

			case lua.ResumeOK:
				retValues = append(retValues, yieldedValues...)
				break loop
			}
		}

		return nil
	}
	err = s.RunGoCallback(ctx, opts, f)
	return
}

type RunOptions struct {
	Timeout time.Duration
}

const DefaultTimeout = time.Millisecond * 200

// RunGoCallback invokes the given Go callback with exclusive access to the state of the
// sandbox.
func (s *Sandbox) RunGoCallback(ctx context.Context, opts RunOptions, f func(ctx context.Context, state *lua.LState) error) (err error) {
	ctx, endObservation := s.operations.runGoCallback.With(ctx, &err, observation.Args{})
	defer endObservation(1, observation.Args{})

	s.m.Lock()
	defer s.m.Unlock()

	if opts.Timeout == 0 {
		opts.Timeout = DefaultTimeout
	}
	ctx, cancel := context.WithTimeout(ctx, opts.Timeout)
	defer cancel()

	s.state.SetContext(ctx)
	defer s.state.RemoveContext()

	return f(ctx, s.state)
}