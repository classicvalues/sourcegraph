// Code generated by go-mockgen 1.1.2; DO NOT EDIT.

package workerutil

import (
	"context"
	"sync"
)

// MockWithHooks is a mock implementation of the WithHooks interface (from
// the package github.com/sourcegraph/sourcegraph/internal/workerutil) used
// for unit testing.
type MockWithHooks struct {
	// PostHandleFunc is an instance of a mock function object controlling
	// the behavior of the method PostHandle.
	PostHandleFunc *WithHooksPostHandleFunc
	// PreHandleFunc is an instance of a mock function object controlling
	// the behavior of the method PreHandle.
	PreHandleFunc *WithHooksPreHandleFunc
}

// NewMockWithHooks creates a new mock of the WithHooks interface. All
// methods return zero values for all results, unless overwritten.
func NewMockWithHooks() *MockWithHooks {
	return &MockWithHooks{
		PostHandleFunc: &WithHooksPostHandleFunc{
			defaultHook: func(context.Context, Record) {
				return
			},
		},
		PreHandleFunc: &WithHooksPreHandleFunc{
			defaultHook: func(context.Context, Record) {
				return
			},
		},
	}
}

// NewStrictMockWithHooks creates a new mock of the WithHooks interface. All
// methods panic on invocation, unless overwritten.
func NewStrictMockWithHooks() *MockWithHooks {
	return &MockWithHooks{
		PostHandleFunc: &WithHooksPostHandleFunc{
			defaultHook: func(context.Context, Record) {
				panic("unexpected invocation of MockWithHooks.PostHandle")
			},
		},
		PreHandleFunc: &WithHooksPreHandleFunc{
			defaultHook: func(context.Context, Record) {
				panic("unexpected invocation of MockWithHooks.PreHandle")
			},
		},
	}
}

// NewMockWithHooksFrom creates a new mock of the MockWithHooks interface.
// All methods delegate to the given implementation, unless overwritten.
func NewMockWithHooksFrom(i WithHooks) *MockWithHooks {
	return &MockWithHooks{
		PostHandleFunc: &WithHooksPostHandleFunc{
			defaultHook: i.PostHandle,
		},
		PreHandleFunc: &WithHooksPreHandleFunc{
			defaultHook: i.PreHandle,
		},
	}
}

// WithHooksPostHandleFunc describes the behavior when the PostHandle method
// of the parent MockWithHooks instance is invoked.
type WithHooksPostHandleFunc struct {
	defaultHook func(context.Context, Record)
	hooks       []func(context.Context, Record)
	history     []WithHooksPostHandleFuncCall
	mutex       sync.Mutex
}

// PostHandle delegates to the next hook function in the queue and stores
// the parameter and result values of this invocation.
func (m *MockWithHooks) PostHandle(v0 context.Context, v1 Record) {
	m.PostHandleFunc.nextHook()(v0, v1)
	m.PostHandleFunc.appendCall(WithHooksPostHandleFuncCall{v0, v1})
	return
}

// SetDefaultHook sets function that is called when the PostHandle method of
// the parent MockWithHooks instance is invoked and the hook queue is empty.
func (f *WithHooksPostHandleFunc) SetDefaultHook(hook func(context.Context, Record)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// PostHandle method of the parent MockWithHooks instance invokes the hook
// at the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *WithHooksPostHandleFunc) PushHook(hook func(context.Context, Record)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *WithHooksPostHandleFunc) SetDefaultReturn() {
	f.SetDefaultHook(func(context.Context, Record) {
		return
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *WithHooksPostHandleFunc) PushReturn() {
	f.PushHook(func(context.Context, Record) {
		return
	})
}

func (f *WithHooksPostHandleFunc) nextHook() func(context.Context, Record) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *WithHooksPostHandleFunc) appendCall(r0 WithHooksPostHandleFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of WithHooksPostHandleFuncCall objects
// describing the invocations of this function.
func (f *WithHooksPostHandleFunc) History() []WithHooksPostHandleFuncCall {
	f.mutex.Lock()
	history := make([]WithHooksPostHandleFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// WithHooksPostHandleFuncCall is an object that describes an invocation of
// method PostHandle on an instance of MockWithHooks.
type WithHooksPostHandleFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 Record
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c WithHooksPostHandleFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c WithHooksPostHandleFuncCall) Results() []interface{} {
	return []interface{}{}
}

// WithHooksPreHandleFunc describes the behavior when the PreHandle method
// of the parent MockWithHooks instance is invoked.
type WithHooksPreHandleFunc struct {
	defaultHook func(context.Context, Record)
	hooks       []func(context.Context, Record)
	history     []WithHooksPreHandleFuncCall
	mutex       sync.Mutex
}

// PreHandle delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockWithHooks) PreHandle(v0 context.Context, v1 Record) {
	m.PreHandleFunc.nextHook()(v0, v1)
	m.PreHandleFunc.appendCall(WithHooksPreHandleFuncCall{v0, v1})
	return
}

// SetDefaultHook sets function that is called when the PreHandle method of
// the parent MockWithHooks instance is invoked and the hook queue is empty.
func (f *WithHooksPreHandleFunc) SetDefaultHook(hook func(context.Context, Record)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// PreHandle method of the parent MockWithHooks instance invokes the hook at
// the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *WithHooksPreHandleFunc) PushHook(hook func(context.Context, Record)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *WithHooksPreHandleFunc) SetDefaultReturn() {
	f.SetDefaultHook(func(context.Context, Record) {
		return
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *WithHooksPreHandleFunc) PushReturn() {
	f.PushHook(func(context.Context, Record) {
		return
	})
}

func (f *WithHooksPreHandleFunc) nextHook() func(context.Context, Record) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *WithHooksPreHandleFunc) appendCall(r0 WithHooksPreHandleFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of WithHooksPreHandleFuncCall objects
// describing the invocations of this function.
func (f *WithHooksPreHandleFunc) History() []WithHooksPreHandleFuncCall {
	f.mutex.Lock()
	history := make([]WithHooksPreHandleFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// WithHooksPreHandleFuncCall is an object that describes an invocation of
// method PreHandle on an instance of MockWithHooks.
type WithHooksPreHandleFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 Record
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c WithHooksPreHandleFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c WithHooksPreHandleFuncCall) Results() []interface{} {
	return []interface{}{}
}