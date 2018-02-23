// generated by "charlatan -output=./server_charlatan.go Server".  DO NOT EDIT.

package auxiliary

import "reflect"

import "time"

// ServerNameInvocation represents a single call of FakeServer.Name
type ServerNameInvocation struct {
	Results struct {
		Ident1 string
	}
}

// ServerAddressInvocation represents a single call of FakeServer.Address
type ServerAddressInvocation struct {
	Results struct {
		Ident1 string
	}
}

// ServerListenInvocation represents a single call of FakeServer.Listen
type ServerListenInvocation struct {
	Results struct {
		Ident1 error
	}
}

// ServerServeInvocation represents a single call of FakeServer.Serve
type ServerServeInvocation struct {
	Results struct {
		Ident1 error
	}
}

// ServerShutdownInvocation represents a single call of FakeServer.Shutdown
type ServerShutdownInvocation struct {
	Parameters struct {
		Ident1 time.Duration
	}
	Results struct {
		Ident2 error
	}
}

// ServerTestingT represents the methods of "testing".T used by charlatan Fakes.  It avoids importing the testing package.
type ServerTestingT interface {
	Error(...interface{})
	Errorf(string, ...interface{})
	Fatal(...interface{})
	Helper()
}

/*
FakeServer is a mock implementation of Server for testing.
Use it in your tests as in this example:

	package example

	func TestWithServer(t *testing.T) {
		f := &auxiliary.FakeServer{
			NameHook: func() (ident1 string) {
				// ensure parameters meet expections, signal errors using t, etc
				return
			},
		}

		// test code goes here ...

		// assert state of FakeName ...
		f.AssertNameCalledOnce(t)
	}

Create anonymous function implementations for only those interface methods that
should be called in the code under test.  This will force a panic if any
unexpected calls are made to FakeName.
*/
type FakeServer struct {
	NameHook     func() string
	AddressHook  func() string
	ListenHook   func() error
	ServeHook    func() error
	ShutdownHook func(time.Duration) error

	NameCalls     []*ServerNameInvocation
	AddressCalls  []*ServerAddressInvocation
	ListenCalls   []*ServerListenInvocation
	ServeCalls    []*ServerServeInvocation
	ShutdownCalls []*ServerShutdownInvocation
}

// NewFakeServerDefaultPanic returns an instance of FakeServer with all hooks configured to panic
func NewFakeServerDefaultPanic() *FakeServer {
	return &FakeServer{
		NameHook: func() (ident1 string) {
			panic("Unexpected call to Server.Name")
		},
		AddressHook: func() (ident1 string) {
			panic("Unexpected call to Server.Address")
		},
		ListenHook: func() (ident1 error) {
			panic("Unexpected call to Server.Listen")
		},
		ServeHook: func() (ident1 error) {
			panic("Unexpected call to Server.Serve")
		},
		ShutdownHook: func(time.Duration) (ident2 error) {
			panic("Unexpected call to Server.Shutdown")
		},
	}
}

// NewFakeServerDefaultFatal returns an instance of FakeServer with all hooks configured to call t.Fatal
func NewFakeServerDefaultFatal(t ServerTestingT) *FakeServer {
	return &FakeServer{
		NameHook: func() (ident1 string) {
			t.Fatal("Unexpected call to Server.Name")
			return
		},
		AddressHook: func() (ident1 string) {
			t.Fatal("Unexpected call to Server.Address")
			return
		},
		ListenHook: func() (ident1 error) {
			t.Fatal("Unexpected call to Server.Listen")
			return
		},
		ServeHook: func() (ident1 error) {
			t.Fatal("Unexpected call to Server.Serve")
			return
		},
		ShutdownHook: func(time.Duration) (ident2 error) {
			t.Fatal("Unexpected call to Server.Shutdown")
			return
		},
	}
}

// NewFakeServerDefaultError returns an instance of FakeServer with all hooks configured to call t.Error
func NewFakeServerDefaultError(t ServerTestingT) *FakeServer {
	return &FakeServer{
		NameHook: func() (ident1 string) {
			t.Error("Unexpected call to Server.Name")
			return
		},
		AddressHook: func() (ident1 string) {
			t.Error("Unexpected call to Server.Address")
			return
		},
		ListenHook: func() (ident1 error) {
			t.Error("Unexpected call to Server.Listen")
			return
		},
		ServeHook: func() (ident1 error) {
			t.Error("Unexpected call to Server.Serve")
			return
		},
		ShutdownHook: func(time.Duration) (ident2 error) {
			t.Error("Unexpected call to Server.Shutdown")
			return
		},
	}
}

func (f *FakeServer) Reset() {
	f.NameCalls = []*ServerNameInvocation{}
	f.AddressCalls = []*ServerAddressInvocation{}
	f.ListenCalls = []*ServerListenInvocation{}
	f.ServeCalls = []*ServerServeInvocation{}
	f.ShutdownCalls = []*ServerShutdownInvocation{}
}

func (_f1 *FakeServer) Name() (ident1 string) {
	invocation := new(ServerNameInvocation)

	ident1 = _f1.NameHook()

	invocation.Results.Ident1 = ident1

	_f1.NameCalls = append(_f1.NameCalls, invocation)

	return
}

// NameCalled returns true if FakeServer.Name was called
func (f *FakeServer) NameCalled() bool {
	return len(f.NameCalls) != 0
}

// AssertNameCalled calls t.Error if FakeServer.Name was not called
func (f *FakeServer) AssertNameCalled(t ServerTestingT) {
	t.Helper()
	if len(f.NameCalls) == 0 {
		t.Error("FakeServer.Name not called, expected at least one")
	}
}

// NameNotCalled returns true if FakeServer.Name was not called
func (f *FakeServer) NameNotCalled() bool {
	return len(f.NameCalls) == 0
}

// AssertNameNotCalled calls t.Error if FakeServer.Name was called
func (f *FakeServer) AssertNameNotCalled(t ServerTestingT) {
	t.Helper()
	if len(f.NameCalls) != 0 {
		t.Error("FakeServer.Name called, expected none")
	}
}

// NameCalledOnce returns true if FakeServer.Name was called exactly once
func (f *FakeServer) NameCalledOnce() bool {
	return len(f.NameCalls) == 1
}

// AssertNameCalledOnce calls t.Error if FakeServer.Name was not called exactly once
func (f *FakeServer) AssertNameCalledOnce(t ServerTestingT) {
	t.Helper()
	if len(f.NameCalls) != 1 {
		t.Errorf("FakeServer.Name called %d times, expected 1", len(f.NameCalls))
	}
}

// NameCalledN returns true if FakeServer.Name was called at least n times
func (f *FakeServer) NameCalledN(n int) bool {
	return len(f.NameCalls) >= n
}

// AssertNameCalledN calls t.Error if FakeServer.Name was called less than n times
func (f *FakeServer) AssertNameCalledN(t ServerTestingT, n int) {
	t.Helper()
	if len(f.NameCalls) < n {
		t.Errorf("FakeServer.Name called %d times, expected >= %d", len(f.NameCalls), n)
	}
}

func (_f2 *FakeServer) Address() (ident1 string) {
	invocation := new(ServerAddressInvocation)

	ident1 = _f2.AddressHook()

	invocation.Results.Ident1 = ident1

	_f2.AddressCalls = append(_f2.AddressCalls, invocation)

	return
}

// AddressCalled returns true if FakeServer.Address was called
func (f *FakeServer) AddressCalled() bool {
	return len(f.AddressCalls) != 0
}

// AssertAddressCalled calls t.Error if FakeServer.Address was not called
func (f *FakeServer) AssertAddressCalled(t ServerTestingT) {
	t.Helper()
	if len(f.AddressCalls) == 0 {
		t.Error("FakeServer.Address not called, expected at least one")
	}
}

// AddressNotCalled returns true if FakeServer.Address was not called
func (f *FakeServer) AddressNotCalled() bool {
	return len(f.AddressCalls) == 0
}

// AssertAddressNotCalled calls t.Error if FakeServer.Address was called
func (f *FakeServer) AssertAddressNotCalled(t ServerTestingT) {
	t.Helper()
	if len(f.AddressCalls) != 0 {
		t.Error("FakeServer.Address called, expected none")
	}
}

// AddressCalledOnce returns true if FakeServer.Address was called exactly once
func (f *FakeServer) AddressCalledOnce() bool {
	return len(f.AddressCalls) == 1
}

// AssertAddressCalledOnce calls t.Error if FakeServer.Address was not called exactly once
func (f *FakeServer) AssertAddressCalledOnce(t ServerTestingT) {
	t.Helper()
	if len(f.AddressCalls) != 1 {
		t.Errorf("FakeServer.Address called %d times, expected 1", len(f.AddressCalls))
	}
}

// AddressCalledN returns true if FakeServer.Address was called at least n times
func (f *FakeServer) AddressCalledN(n int) bool {
	return len(f.AddressCalls) >= n
}

// AssertAddressCalledN calls t.Error if FakeServer.Address was called less than n times
func (f *FakeServer) AssertAddressCalledN(t ServerTestingT, n int) {
	t.Helper()
	if len(f.AddressCalls) < n {
		t.Errorf("FakeServer.Address called %d times, expected >= %d", len(f.AddressCalls), n)
	}
}

func (_f3 *FakeServer) Listen() (ident1 error) {
	invocation := new(ServerListenInvocation)

	ident1 = _f3.ListenHook()

	invocation.Results.Ident1 = ident1

	_f3.ListenCalls = append(_f3.ListenCalls, invocation)

	return
}

// ListenCalled returns true if FakeServer.Listen was called
func (f *FakeServer) ListenCalled() bool {
	return len(f.ListenCalls) != 0
}

// AssertListenCalled calls t.Error if FakeServer.Listen was not called
func (f *FakeServer) AssertListenCalled(t ServerTestingT) {
	t.Helper()
	if len(f.ListenCalls) == 0 {
		t.Error("FakeServer.Listen not called, expected at least one")
	}
}

// ListenNotCalled returns true if FakeServer.Listen was not called
func (f *FakeServer) ListenNotCalled() bool {
	return len(f.ListenCalls) == 0
}

// AssertListenNotCalled calls t.Error if FakeServer.Listen was called
func (f *FakeServer) AssertListenNotCalled(t ServerTestingT) {
	t.Helper()
	if len(f.ListenCalls) != 0 {
		t.Error("FakeServer.Listen called, expected none")
	}
}

// ListenCalledOnce returns true if FakeServer.Listen was called exactly once
func (f *FakeServer) ListenCalledOnce() bool {
	return len(f.ListenCalls) == 1
}

// AssertListenCalledOnce calls t.Error if FakeServer.Listen was not called exactly once
func (f *FakeServer) AssertListenCalledOnce(t ServerTestingT) {
	t.Helper()
	if len(f.ListenCalls) != 1 {
		t.Errorf("FakeServer.Listen called %d times, expected 1", len(f.ListenCalls))
	}
}

// ListenCalledN returns true if FakeServer.Listen was called at least n times
func (f *FakeServer) ListenCalledN(n int) bool {
	return len(f.ListenCalls) >= n
}

// AssertListenCalledN calls t.Error if FakeServer.Listen was called less than n times
func (f *FakeServer) AssertListenCalledN(t ServerTestingT, n int) {
	t.Helper()
	if len(f.ListenCalls) < n {
		t.Errorf("FakeServer.Listen called %d times, expected >= %d", len(f.ListenCalls), n)
	}
}

func (_f4 *FakeServer) Serve() (ident1 error) {
	invocation := new(ServerServeInvocation)

	ident1 = _f4.ServeHook()

	invocation.Results.Ident1 = ident1

	_f4.ServeCalls = append(_f4.ServeCalls, invocation)

	return
}

// ServeCalled returns true if FakeServer.Serve was called
func (f *FakeServer) ServeCalled() bool {
	return len(f.ServeCalls) != 0
}

// AssertServeCalled calls t.Error if FakeServer.Serve was not called
func (f *FakeServer) AssertServeCalled(t ServerTestingT) {
	t.Helper()
	if len(f.ServeCalls) == 0 {
		t.Error("FakeServer.Serve not called, expected at least one")
	}
}

// ServeNotCalled returns true if FakeServer.Serve was not called
func (f *FakeServer) ServeNotCalled() bool {
	return len(f.ServeCalls) == 0
}

// AssertServeNotCalled calls t.Error if FakeServer.Serve was called
func (f *FakeServer) AssertServeNotCalled(t ServerTestingT) {
	t.Helper()
	if len(f.ServeCalls) != 0 {
		t.Error("FakeServer.Serve called, expected none")
	}
}

// ServeCalledOnce returns true if FakeServer.Serve was called exactly once
func (f *FakeServer) ServeCalledOnce() bool {
	return len(f.ServeCalls) == 1
}

// AssertServeCalledOnce calls t.Error if FakeServer.Serve was not called exactly once
func (f *FakeServer) AssertServeCalledOnce(t ServerTestingT) {
	t.Helper()
	if len(f.ServeCalls) != 1 {
		t.Errorf("FakeServer.Serve called %d times, expected 1", len(f.ServeCalls))
	}
}

// ServeCalledN returns true if FakeServer.Serve was called at least n times
func (f *FakeServer) ServeCalledN(n int) bool {
	return len(f.ServeCalls) >= n
}

// AssertServeCalledN calls t.Error if FakeServer.Serve was called less than n times
func (f *FakeServer) AssertServeCalledN(t ServerTestingT, n int) {
	t.Helper()
	if len(f.ServeCalls) < n {
		t.Errorf("FakeServer.Serve called %d times, expected >= %d", len(f.ServeCalls), n)
	}
}

func (_f5 *FakeServer) Shutdown(ident1 time.Duration) (ident2 error) {
	invocation := new(ServerShutdownInvocation)

	invocation.Parameters.Ident1 = ident1

	ident2 = _f5.ShutdownHook(ident1)

	invocation.Results.Ident2 = ident2

	_f5.ShutdownCalls = append(_f5.ShutdownCalls, invocation)

	return
}

// ShutdownCalled returns true if FakeServer.Shutdown was called
func (f *FakeServer) ShutdownCalled() bool {
	return len(f.ShutdownCalls) != 0
}

// AssertShutdownCalled calls t.Error if FakeServer.Shutdown was not called
func (f *FakeServer) AssertShutdownCalled(t ServerTestingT) {
	t.Helper()
	if len(f.ShutdownCalls) == 0 {
		t.Error("FakeServer.Shutdown not called, expected at least one")
	}
}

// ShutdownNotCalled returns true if FakeServer.Shutdown was not called
func (f *FakeServer) ShutdownNotCalled() bool {
	return len(f.ShutdownCalls) == 0
}

// AssertShutdownNotCalled calls t.Error if FakeServer.Shutdown was called
func (f *FakeServer) AssertShutdownNotCalled(t ServerTestingT) {
	t.Helper()
	if len(f.ShutdownCalls) != 0 {
		t.Error("FakeServer.Shutdown called, expected none")
	}
}

// ShutdownCalledOnce returns true if FakeServer.Shutdown was called exactly once
func (f *FakeServer) ShutdownCalledOnce() bool {
	return len(f.ShutdownCalls) == 1
}

// AssertShutdownCalledOnce calls t.Error if FakeServer.Shutdown was not called exactly once
func (f *FakeServer) AssertShutdownCalledOnce(t ServerTestingT) {
	t.Helper()
	if len(f.ShutdownCalls) != 1 {
		t.Errorf("FakeServer.Shutdown called %d times, expected 1", len(f.ShutdownCalls))
	}
}

// ShutdownCalledN returns true if FakeServer.Shutdown was called at least n times
func (f *FakeServer) ShutdownCalledN(n int) bool {
	return len(f.ShutdownCalls) >= n
}

// AssertShutdownCalledN calls t.Error if FakeServer.Shutdown was called less than n times
func (f *FakeServer) AssertShutdownCalledN(t ServerTestingT, n int) {
	t.Helper()
	if len(f.ShutdownCalls) < n {
		t.Errorf("FakeServer.Shutdown called %d times, expected >= %d", len(f.ShutdownCalls), n)
	}
}

// ShutdownCalledWith returns true if FakeServer.Shutdown was called with the given values
func (_f6 *FakeServer) ShutdownCalledWith(ident1 time.Duration) (found bool) {
	for _, call := range _f6.ShutdownCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			found = true
			break
		}
	}

	return
}

// AssertShutdownCalledWith calls t.Error if FakeServer.Shutdown was not called with the given values
func (_f7 *FakeServer) AssertShutdownCalledWith(t ServerTestingT, ident1 time.Duration) {
	t.Helper()
	var found bool
	for _, call := range _f7.ShutdownCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			found = true
			break
		}
	}

	if !found {
		t.Error("FakeServer.Shutdown not called with expected parameters")
	}
}

// ShutdownCalledOnceWith returns true if FakeServer.Shutdown was called exactly once with the given values
func (_f8 *FakeServer) ShutdownCalledOnceWith(ident1 time.Duration) bool {
	var count int
	for _, call := range _f8.ShutdownCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			count++
		}
	}

	return count == 1
}

// AssertShutdownCalledOnceWith calls t.Error if FakeServer.Shutdown was not called exactly once with the given values
func (_f9 *FakeServer) AssertShutdownCalledOnceWith(t ServerTestingT, ident1 time.Duration) {
	t.Helper()
	var count int
	for _, call := range _f9.ShutdownCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			count++
		}
	}

	if count != 1 {
		t.Errorf("FakeServer.Shutdown called %d times with expected parameters, expected one", count)
	}
}

// ShutdownResultsForCall returns the result values for the first call to FakeServer.Shutdown with the given values
func (_f10 *FakeServer) ShutdownResultsForCall(ident1 time.Duration) (ident2 error, found bool) {
	for _, call := range _f10.ShutdownCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			ident2 = call.Results.Ident2
			found = true
			break
		}
	}

	return
}
