// Code generated by go-mockgen 1.3.2; DO NOT EDIT.
//
// This file was generated by running `sg generate` (or `go-mockgen`) at the root of
// this repository. To add additional mocks to this or another package, add a new entry
// to the mockgen.yaml file in the root of this repository.

package uploads

import (
	"context"
	"sync"
	"time"

	store "github.com/sourcegraph/sourcegraph/internal/codeintel/uploads/internal/store"
	shared "github.com/sourcegraph/sourcegraph/internal/codeintel/uploads/shared"
)

// MockStore is a mock implementation of the Store interface (from the
// package
// github.com/sourcegraph/sourcegraph/internal/codeintel/uploads/internal/store)
// used for unit testing.
type MockStore struct {
	// DeleteIndexesWithoutRepositoryFunc is an instance of a mock function
	// object controlling the behavior of the method
	// DeleteIndexesWithoutRepository.
	DeleteIndexesWithoutRepositoryFunc *StoreDeleteIndexesWithoutRepositoryFunc
	// DeleteSourcedCommitsFunc is an instance of a mock function object
	// controlling the behavior of the method DeleteSourcedCommits.
	DeleteSourcedCommitsFunc *StoreDeleteSourcedCommitsFunc
	// DeleteUploadsWithoutRepositoryFunc is an instance of a mock function
	// object controlling the behavior of the method
	// DeleteUploadsWithoutRepository.
	DeleteUploadsWithoutRepositoryFunc *StoreDeleteUploadsWithoutRepositoryFunc
	// ListFunc is an instance of a mock function object controlling the
	// behavior of the method List.
	ListFunc *StoreListFunc
	// StaleSourcedCommitsFunc is an instance of a mock function object
	// controlling the behavior of the method StaleSourcedCommits.
	StaleSourcedCommitsFunc *StoreStaleSourcedCommitsFunc
	// UpdateSourcedCommitsFunc is an instance of a mock function object
	// controlling the behavior of the method UpdateSourcedCommits.
	UpdateSourcedCommitsFunc *StoreUpdateSourcedCommitsFunc
}

// NewMockStore creates a new mock of the Store interface. All methods
// return zero values for all results, unless overwritten.
func NewMockStore() *MockStore {
	return &MockStore{
		DeleteIndexesWithoutRepositoryFunc: &StoreDeleteIndexesWithoutRepositoryFunc{
			defaultHook: func(context.Context, time.Time) (r0 map[int]int, r1 error) {
				return
			},
		},
		DeleteSourcedCommitsFunc: &StoreDeleteSourcedCommitsFunc{
			defaultHook: func(context.Context, int, string, time.Duration, time.Time) (r0 int, r1 int, r2 int, r3 error) {
				return
			},
		},
		DeleteUploadsWithoutRepositoryFunc: &StoreDeleteUploadsWithoutRepositoryFunc{
			defaultHook: func(context.Context, time.Time) (r0 map[int]int, r1 error) {
				return
			},
		},
		ListFunc: &StoreListFunc{
			defaultHook: func(context.Context, store.ListOpts) (r0 []shared.Upload, r1 error) {
				return
			},
		},
		StaleSourcedCommitsFunc: &StoreStaleSourcedCommitsFunc{
			defaultHook: func(context.Context, time.Duration, int, time.Time) (r0 []shared.SourcedCommits, r1 error) {
				return
			},
		},
		UpdateSourcedCommitsFunc: &StoreUpdateSourcedCommitsFunc{
			defaultHook: func(context.Context, int, string, time.Time) (r0 int, r1 int, r2 error) {
				return
			},
		},
	}
}

// NewStrictMockStore creates a new mock of the Store interface. All methods
// panic on invocation, unless overwritten.
func NewStrictMockStore() *MockStore {
	return &MockStore{
		DeleteIndexesWithoutRepositoryFunc: &StoreDeleteIndexesWithoutRepositoryFunc{
			defaultHook: func(context.Context, time.Time) (map[int]int, error) {
				panic("unexpected invocation of MockStore.DeleteIndexesWithoutRepository")
			},
		},
		DeleteSourcedCommitsFunc: &StoreDeleteSourcedCommitsFunc{
			defaultHook: func(context.Context, int, string, time.Duration, time.Time) (int, int, int, error) {
				panic("unexpected invocation of MockStore.DeleteSourcedCommits")
			},
		},
		DeleteUploadsWithoutRepositoryFunc: &StoreDeleteUploadsWithoutRepositoryFunc{
			defaultHook: func(context.Context, time.Time) (map[int]int, error) {
				panic("unexpected invocation of MockStore.DeleteUploadsWithoutRepository")
			},
		},
		ListFunc: &StoreListFunc{
			defaultHook: func(context.Context, store.ListOpts) ([]shared.Upload, error) {
				panic("unexpected invocation of MockStore.List")
			},
		},
		StaleSourcedCommitsFunc: &StoreStaleSourcedCommitsFunc{
			defaultHook: func(context.Context, time.Duration, int, time.Time) ([]shared.SourcedCommits, error) {
				panic("unexpected invocation of MockStore.StaleSourcedCommits")
			},
		},
		UpdateSourcedCommitsFunc: &StoreUpdateSourcedCommitsFunc{
			defaultHook: func(context.Context, int, string, time.Time) (int, int, error) {
				panic("unexpected invocation of MockStore.UpdateSourcedCommits")
			},
		},
	}
}

// NewMockStoreFrom creates a new mock of the MockStore interface. All
// methods delegate to the given implementation, unless overwritten.
func NewMockStoreFrom(i store.Store) *MockStore {
	return &MockStore{
		DeleteIndexesWithoutRepositoryFunc: &StoreDeleteIndexesWithoutRepositoryFunc{
			defaultHook: i.DeleteIndexesWithoutRepository,
		},
		DeleteSourcedCommitsFunc: &StoreDeleteSourcedCommitsFunc{
			defaultHook: i.DeleteSourcedCommits,
		},
		DeleteUploadsWithoutRepositoryFunc: &StoreDeleteUploadsWithoutRepositoryFunc{
			defaultHook: i.DeleteUploadsWithoutRepository,
		},
		ListFunc: &StoreListFunc{
			defaultHook: i.List,
		},
		StaleSourcedCommitsFunc: &StoreStaleSourcedCommitsFunc{
			defaultHook: i.StaleSourcedCommits,
		},
		UpdateSourcedCommitsFunc: &StoreUpdateSourcedCommitsFunc{
			defaultHook: i.UpdateSourcedCommits,
		},
	}
}

// StoreDeleteIndexesWithoutRepositoryFunc describes the behavior when the
// DeleteIndexesWithoutRepository method of the parent MockStore instance is
// invoked.
type StoreDeleteIndexesWithoutRepositoryFunc struct {
	defaultHook func(context.Context, time.Time) (map[int]int, error)
	hooks       []func(context.Context, time.Time) (map[int]int, error)
	history     []StoreDeleteIndexesWithoutRepositoryFuncCall
	mutex       sync.Mutex
}

// DeleteIndexesWithoutRepository delegates to the next hook function in the
// queue and stores the parameter and result values of this invocation.
func (m *MockStore) DeleteIndexesWithoutRepository(v0 context.Context, v1 time.Time) (map[int]int, error) {
	r0, r1 := m.DeleteIndexesWithoutRepositoryFunc.nextHook()(v0, v1)
	m.DeleteIndexesWithoutRepositoryFunc.appendCall(StoreDeleteIndexesWithoutRepositoryFuncCall{v0, v1, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the
// DeleteIndexesWithoutRepository method of the parent MockStore instance is
// invoked and the hook queue is empty.
func (f *StoreDeleteIndexesWithoutRepositoryFunc) SetDefaultHook(hook func(context.Context, time.Time) (map[int]int, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// DeleteIndexesWithoutRepository method of the parent MockStore instance
// invokes the hook at the front of the queue and discards it. After the
// queue is empty, the default hook function is invoked for any future
// action.
func (f *StoreDeleteIndexesWithoutRepositoryFunc) PushHook(hook func(context.Context, time.Time) (map[int]int, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *StoreDeleteIndexesWithoutRepositoryFunc) SetDefaultReturn(r0 map[int]int, r1 error) {
	f.SetDefaultHook(func(context.Context, time.Time) (map[int]int, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *StoreDeleteIndexesWithoutRepositoryFunc) PushReturn(r0 map[int]int, r1 error) {
	f.PushHook(func(context.Context, time.Time) (map[int]int, error) {
		return r0, r1
	})
}

func (f *StoreDeleteIndexesWithoutRepositoryFunc) nextHook() func(context.Context, time.Time) (map[int]int, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StoreDeleteIndexesWithoutRepositoryFunc) appendCall(r0 StoreDeleteIndexesWithoutRepositoryFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StoreDeleteIndexesWithoutRepositoryFuncCall
// objects describing the invocations of this function.
func (f *StoreDeleteIndexesWithoutRepositoryFunc) History() []StoreDeleteIndexesWithoutRepositoryFuncCall {
	f.mutex.Lock()
	history := make([]StoreDeleteIndexesWithoutRepositoryFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StoreDeleteIndexesWithoutRepositoryFuncCall is an object that describes
// an invocation of method DeleteIndexesWithoutRepository on an instance of
// MockStore.
type StoreDeleteIndexesWithoutRepositoryFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 time.Time
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 map[int]int
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StoreDeleteIndexesWithoutRepositoryFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StoreDeleteIndexesWithoutRepositoryFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// StoreDeleteSourcedCommitsFunc describes the behavior when the
// DeleteSourcedCommits method of the parent MockStore instance is invoked.
type StoreDeleteSourcedCommitsFunc struct {
	defaultHook func(context.Context, int, string, time.Duration, time.Time) (int, int, int, error)
	hooks       []func(context.Context, int, string, time.Duration, time.Time) (int, int, int, error)
	history     []StoreDeleteSourcedCommitsFuncCall
	mutex       sync.Mutex
}

// DeleteSourcedCommits delegates to the next hook function in the queue and
// stores the parameter and result values of this invocation.
func (m *MockStore) DeleteSourcedCommits(v0 context.Context, v1 int, v2 string, v3 time.Duration, v4 time.Time) (int, int, int, error) {
	r0, r1, r2, r3 := m.DeleteSourcedCommitsFunc.nextHook()(v0, v1, v2, v3, v4)
	m.DeleteSourcedCommitsFunc.appendCall(StoreDeleteSourcedCommitsFuncCall{v0, v1, v2, v3, v4, r0, r1, r2, r3})
	return r0, r1, r2, r3
}

// SetDefaultHook sets function that is called when the DeleteSourcedCommits
// method of the parent MockStore instance is invoked and the hook queue is
// empty.
func (f *StoreDeleteSourcedCommitsFunc) SetDefaultHook(hook func(context.Context, int, string, time.Duration, time.Time) (int, int, int, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// DeleteSourcedCommits method of the parent MockStore instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *StoreDeleteSourcedCommitsFunc) PushHook(hook func(context.Context, int, string, time.Duration, time.Time) (int, int, int, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *StoreDeleteSourcedCommitsFunc) SetDefaultReturn(r0 int, r1 int, r2 int, r3 error) {
	f.SetDefaultHook(func(context.Context, int, string, time.Duration, time.Time) (int, int, int, error) {
		return r0, r1, r2, r3
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *StoreDeleteSourcedCommitsFunc) PushReturn(r0 int, r1 int, r2 int, r3 error) {
	f.PushHook(func(context.Context, int, string, time.Duration, time.Time) (int, int, int, error) {
		return r0, r1, r2, r3
	})
}

func (f *StoreDeleteSourcedCommitsFunc) nextHook() func(context.Context, int, string, time.Duration, time.Time) (int, int, int, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StoreDeleteSourcedCommitsFunc) appendCall(r0 StoreDeleteSourcedCommitsFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StoreDeleteSourcedCommitsFuncCall objects
// describing the invocations of this function.
func (f *StoreDeleteSourcedCommitsFunc) History() []StoreDeleteSourcedCommitsFuncCall {
	f.mutex.Lock()
	history := make([]StoreDeleteSourcedCommitsFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StoreDeleteSourcedCommitsFuncCall is an object that describes an
// invocation of method DeleteSourcedCommits on an instance of MockStore.
type StoreDeleteSourcedCommitsFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 int
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 string
	// Arg3 is the value of the 4th argument passed to this method
	// invocation.
	Arg3 time.Duration
	// Arg4 is the value of the 5th argument passed to this method
	// invocation.
	Arg4 time.Time
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 int
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 int
	// Result2 is the value of the 3rd result returned from this method
	// invocation.
	Result2 int
	// Result3 is the value of the 4th result returned from this method
	// invocation.
	Result3 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StoreDeleteSourcedCommitsFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2, c.Arg3, c.Arg4}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StoreDeleteSourcedCommitsFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1, c.Result2, c.Result3}
}

// StoreDeleteUploadsWithoutRepositoryFunc describes the behavior when the
// DeleteUploadsWithoutRepository method of the parent MockStore instance is
// invoked.
type StoreDeleteUploadsWithoutRepositoryFunc struct {
	defaultHook func(context.Context, time.Time) (map[int]int, error)
	hooks       []func(context.Context, time.Time) (map[int]int, error)
	history     []StoreDeleteUploadsWithoutRepositoryFuncCall
	mutex       sync.Mutex
}

// DeleteUploadsWithoutRepository delegates to the next hook function in the
// queue and stores the parameter and result values of this invocation.
func (m *MockStore) DeleteUploadsWithoutRepository(v0 context.Context, v1 time.Time) (map[int]int, error) {
	r0, r1 := m.DeleteUploadsWithoutRepositoryFunc.nextHook()(v0, v1)
	m.DeleteUploadsWithoutRepositoryFunc.appendCall(StoreDeleteUploadsWithoutRepositoryFuncCall{v0, v1, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the
// DeleteUploadsWithoutRepository method of the parent MockStore instance is
// invoked and the hook queue is empty.
func (f *StoreDeleteUploadsWithoutRepositoryFunc) SetDefaultHook(hook func(context.Context, time.Time) (map[int]int, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// DeleteUploadsWithoutRepository method of the parent MockStore instance
// invokes the hook at the front of the queue and discards it. After the
// queue is empty, the default hook function is invoked for any future
// action.
func (f *StoreDeleteUploadsWithoutRepositoryFunc) PushHook(hook func(context.Context, time.Time) (map[int]int, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *StoreDeleteUploadsWithoutRepositoryFunc) SetDefaultReturn(r0 map[int]int, r1 error) {
	f.SetDefaultHook(func(context.Context, time.Time) (map[int]int, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *StoreDeleteUploadsWithoutRepositoryFunc) PushReturn(r0 map[int]int, r1 error) {
	f.PushHook(func(context.Context, time.Time) (map[int]int, error) {
		return r0, r1
	})
}

func (f *StoreDeleteUploadsWithoutRepositoryFunc) nextHook() func(context.Context, time.Time) (map[int]int, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StoreDeleteUploadsWithoutRepositoryFunc) appendCall(r0 StoreDeleteUploadsWithoutRepositoryFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StoreDeleteUploadsWithoutRepositoryFuncCall
// objects describing the invocations of this function.
func (f *StoreDeleteUploadsWithoutRepositoryFunc) History() []StoreDeleteUploadsWithoutRepositoryFuncCall {
	f.mutex.Lock()
	history := make([]StoreDeleteUploadsWithoutRepositoryFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StoreDeleteUploadsWithoutRepositoryFuncCall is an object that describes
// an invocation of method DeleteUploadsWithoutRepository on an instance of
// MockStore.
type StoreDeleteUploadsWithoutRepositoryFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 time.Time
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 map[int]int
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StoreDeleteUploadsWithoutRepositoryFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StoreDeleteUploadsWithoutRepositoryFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// StoreListFunc describes the behavior when the List method of the parent
// MockStore instance is invoked.
type StoreListFunc struct {
	defaultHook func(context.Context, store.ListOpts) ([]shared.Upload, error)
	hooks       []func(context.Context, store.ListOpts) ([]shared.Upload, error)
	history     []StoreListFuncCall
	mutex       sync.Mutex
}

// List delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockStore) List(v0 context.Context, v1 store.ListOpts) ([]shared.Upload, error) {
	r0, r1 := m.ListFunc.nextHook()(v0, v1)
	m.ListFunc.appendCall(StoreListFuncCall{v0, v1, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the List method of the
// parent MockStore instance is invoked and the hook queue is empty.
func (f *StoreListFunc) SetDefaultHook(hook func(context.Context, store.ListOpts) ([]shared.Upload, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// List method of the parent MockStore instance invokes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *StoreListFunc) PushHook(hook func(context.Context, store.ListOpts) ([]shared.Upload, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *StoreListFunc) SetDefaultReturn(r0 []shared.Upload, r1 error) {
	f.SetDefaultHook(func(context.Context, store.ListOpts) ([]shared.Upload, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *StoreListFunc) PushReturn(r0 []shared.Upload, r1 error) {
	f.PushHook(func(context.Context, store.ListOpts) ([]shared.Upload, error) {
		return r0, r1
	})
}

func (f *StoreListFunc) nextHook() func(context.Context, store.ListOpts) ([]shared.Upload, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StoreListFunc) appendCall(r0 StoreListFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StoreListFuncCall objects describing the
// invocations of this function.
func (f *StoreListFunc) History() []StoreListFuncCall {
	f.mutex.Lock()
	history := make([]StoreListFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StoreListFuncCall is an object that describes an invocation of method
// List on an instance of MockStore.
type StoreListFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 store.ListOpts
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 []shared.Upload
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StoreListFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StoreListFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// StoreStaleSourcedCommitsFunc describes the behavior when the
// StaleSourcedCommits method of the parent MockStore instance is invoked.
type StoreStaleSourcedCommitsFunc struct {
	defaultHook func(context.Context, time.Duration, int, time.Time) ([]shared.SourcedCommits, error)
	hooks       []func(context.Context, time.Duration, int, time.Time) ([]shared.SourcedCommits, error)
	history     []StoreStaleSourcedCommitsFuncCall
	mutex       sync.Mutex
}

// StaleSourcedCommits delegates to the next hook function in the queue and
// stores the parameter and result values of this invocation.
func (m *MockStore) StaleSourcedCommits(v0 context.Context, v1 time.Duration, v2 int, v3 time.Time) ([]shared.SourcedCommits, error) {
	r0, r1 := m.StaleSourcedCommitsFunc.nextHook()(v0, v1, v2, v3)
	m.StaleSourcedCommitsFunc.appendCall(StoreStaleSourcedCommitsFuncCall{v0, v1, v2, v3, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the StaleSourcedCommits
// method of the parent MockStore instance is invoked and the hook queue is
// empty.
func (f *StoreStaleSourcedCommitsFunc) SetDefaultHook(hook func(context.Context, time.Duration, int, time.Time) ([]shared.SourcedCommits, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// StaleSourcedCommits method of the parent MockStore instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *StoreStaleSourcedCommitsFunc) PushHook(hook func(context.Context, time.Duration, int, time.Time) ([]shared.SourcedCommits, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *StoreStaleSourcedCommitsFunc) SetDefaultReturn(r0 []shared.SourcedCommits, r1 error) {
	f.SetDefaultHook(func(context.Context, time.Duration, int, time.Time) ([]shared.SourcedCommits, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *StoreStaleSourcedCommitsFunc) PushReturn(r0 []shared.SourcedCommits, r1 error) {
	f.PushHook(func(context.Context, time.Duration, int, time.Time) ([]shared.SourcedCommits, error) {
		return r0, r1
	})
}

func (f *StoreStaleSourcedCommitsFunc) nextHook() func(context.Context, time.Duration, int, time.Time) ([]shared.SourcedCommits, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StoreStaleSourcedCommitsFunc) appendCall(r0 StoreStaleSourcedCommitsFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StoreStaleSourcedCommitsFuncCall objects
// describing the invocations of this function.
func (f *StoreStaleSourcedCommitsFunc) History() []StoreStaleSourcedCommitsFuncCall {
	f.mutex.Lock()
	history := make([]StoreStaleSourcedCommitsFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StoreStaleSourcedCommitsFuncCall is an object that describes an
// invocation of method StaleSourcedCommits on an instance of MockStore.
type StoreStaleSourcedCommitsFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 time.Duration
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 int
	// Arg3 is the value of the 4th argument passed to this method
	// invocation.
	Arg3 time.Time
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 []shared.SourcedCommits
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StoreStaleSourcedCommitsFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2, c.Arg3}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StoreStaleSourcedCommitsFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// StoreUpdateSourcedCommitsFunc describes the behavior when the
// UpdateSourcedCommits method of the parent MockStore instance is invoked.
type StoreUpdateSourcedCommitsFunc struct {
	defaultHook func(context.Context, int, string, time.Time) (int, int, error)
	hooks       []func(context.Context, int, string, time.Time) (int, int, error)
	history     []StoreUpdateSourcedCommitsFuncCall
	mutex       sync.Mutex
}

// UpdateSourcedCommits delegates to the next hook function in the queue and
// stores the parameter and result values of this invocation.
func (m *MockStore) UpdateSourcedCommits(v0 context.Context, v1 int, v2 string, v3 time.Time) (int, int, error) {
	r0, r1, r2 := m.UpdateSourcedCommitsFunc.nextHook()(v0, v1, v2, v3)
	m.UpdateSourcedCommitsFunc.appendCall(StoreUpdateSourcedCommitsFuncCall{v0, v1, v2, v3, r0, r1, r2})
	return r0, r1, r2
}

// SetDefaultHook sets function that is called when the UpdateSourcedCommits
// method of the parent MockStore instance is invoked and the hook queue is
// empty.
func (f *StoreUpdateSourcedCommitsFunc) SetDefaultHook(hook func(context.Context, int, string, time.Time) (int, int, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// UpdateSourcedCommits method of the parent MockStore instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *StoreUpdateSourcedCommitsFunc) PushHook(hook func(context.Context, int, string, time.Time) (int, int, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *StoreUpdateSourcedCommitsFunc) SetDefaultReturn(r0 int, r1 int, r2 error) {
	f.SetDefaultHook(func(context.Context, int, string, time.Time) (int, int, error) {
		return r0, r1, r2
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *StoreUpdateSourcedCommitsFunc) PushReturn(r0 int, r1 int, r2 error) {
	f.PushHook(func(context.Context, int, string, time.Time) (int, int, error) {
		return r0, r1, r2
	})
}

func (f *StoreUpdateSourcedCommitsFunc) nextHook() func(context.Context, int, string, time.Time) (int, int, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StoreUpdateSourcedCommitsFunc) appendCall(r0 StoreUpdateSourcedCommitsFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StoreUpdateSourcedCommitsFuncCall objects
// describing the invocations of this function.
func (f *StoreUpdateSourcedCommitsFunc) History() []StoreUpdateSourcedCommitsFuncCall {
	f.mutex.Lock()
	history := make([]StoreUpdateSourcedCommitsFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StoreUpdateSourcedCommitsFuncCall is an object that describes an
// invocation of method UpdateSourcedCommits on an instance of MockStore.
type StoreUpdateSourcedCommitsFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 int
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 string
	// Arg3 is the value of the 4th argument passed to this method
	// invocation.
	Arg3 time.Time
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 int
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 int
	// Result2 is the value of the 3rd result returned from this method
	// invocation.
	Result2 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StoreUpdateSourcedCommitsFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2, c.Arg3}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StoreUpdateSourcedCommitsFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1, c.Result2}
}