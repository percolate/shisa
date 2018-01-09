// generated by "charlatan -output=./provider_charlatan.go Provider".  DO NOT EDIT.

package env

import "reflect"
import "github.com/ansel1/merry"

// ProviderGetInvocation represents a single call of FakeProvider.Get
type ProviderGetInvocation struct {
	Parameters struct {
		Ident1 string
	}
	Results struct {
		Ident2 string
		Ident3 merry.Error
	}
}

// ProviderGetIntInvocation represents a single call of FakeProvider.GetInt
type ProviderGetIntInvocation struct {
	Parameters struct {
		Ident1 string
	}
	Results struct {
		Ident2 int
		Ident3 merry.Error
	}
}

// ProviderGetBoolInvocation represents a single call of FakeProvider.GetBool
type ProviderGetBoolInvocation struct {
	Parameters struct {
		Ident1 string
	}
	Results struct {
		Ident2 bool
		Ident3 merry.Error
	}
}

// ProviderMonitorInvocation represents a single call of FakeProvider.Monitor
type ProviderMonitorInvocation struct {
	Parameters struct {
		Ident1 string
		Ident2 <-chan Value
	}
}

// ProviderHealthcheckInvocation represents a single call of FakeProvider.Healthcheck
type ProviderHealthcheckInvocation struct {
	Results struct {
		Ident1 merry.Error
	}
}

// ProviderTestingT represents the methods of "testing".T used by charlatan Fakes.  It avoids importing the testing package.
type ProviderTestingT interface {
	Error(...interface{})
	Errorf(string, ...interface{})
	Fatal(...interface{})
	Helper()
}

/*
FakeProvider is a mock implementation of Provider for testing.
Use it in your tests as in this example:

	package example

	func TestWithProvider(t *testing.T) {
		f := &env.FakeProvider{
			GetHook: func(ident1 string) (ident2 string, ident3 merry.Error) {
				// ensure parameters meet expections, signal errors using t, etc
				return
			},
		}

		// test code goes here ...

		// assert state of FakeGet ...
		f.AssertGetCalledOnce(t)
	}

Create anonymous function implementations for only those interface methods that
should be called in the code under test.  This will force a panic if any
unexpected calls are made to FakeGet.
*/
type FakeProvider struct {
	GetHook         func(string) (string, merry.Error)
	GetIntHook      func(string) (int, merry.Error)
	GetBoolHook     func(string) (bool, merry.Error)
	MonitorHook     func(string, <-chan Value)
	HealthcheckHook func() merry.Error

	GetCalls         []*ProviderGetInvocation
	GetIntCalls      []*ProviderGetIntInvocation
	GetBoolCalls     []*ProviderGetBoolInvocation
	MonitorCalls     []*ProviderMonitorInvocation
	HealthcheckCalls []*ProviderHealthcheckInvocation
}

// NewFakeProviderDefaultPanic returns an instance of FakeProvider with all hooks configured to panic
func NewFakeProviderDefaultPanic() *FakeProvider {
	return &FakeProvider{
		GetHook: func(string) (ident2 string, ident3 merry.Error) {
			panic("Unexpected call to Provider.Get")
		},
		GetIntHook: func(string) (ident2 int, ident3 merry.Error) {
			panic("Unexpected call to Provider.GetInt")
		},
		GetBoolHook: func(string) (ident2 bool, ident3 merry.Error) {
			panic("Unexpected call to Provider.GetBool")
		},
		MonitorHook: func(string, <-chan Value) {
			panic("Unexpected call to Provider.Monitor")
		},
		HealthcheckHook: func() (ident1 merry.Error) {
			panic("Unexpected call to Provider.Healthcheck")
		},
	}
}

// NewFakeProviderDefaultFatal returns an instance of FakeProvider with all hooks configured to call t.Fatal
func NewFakeProviderDefaultFatal(t ProviderTestingT) *FakeProvider {
	return &FakeProvider{
		GetHook: func(string) (ident2 string, ident3 merry.Error) {
			t.Fatal("Unexpected call to Provider.Get")
			return
		},
		GetIntHook: func(string) (ident2 int, ident3 merry.Error) {
			t.Fatal("Unexpected call to Provider.GetInt")
			return
		},
		GetBoolHook: func(string) (ident2 bool, ident3 merry.Error) {
			t.Fatal("Unexpected call to Provider.GetBool")
			return
		},
		MonitorHook: func(string, <-chan Value) {
			t.Fatal("Unexpected call to Provider.Monitor")
			return
		},
		HealthcheckHook: func() (ident1 merry.Error) {
			t.Fatal("Unexpected call to Provider.Healthcheck")
			return
		},
	}
}

// NewFakeProviderDefaultError returns an instance of FakeProvider with all hooks configured to call t.Error
func NewFakeProviderDefaultError(t ProviderTestingT) *FakeProvider {
	return &FakeProvider{
		GetHook: func(string) (ident2 string, ident3 merry.Error) {
			t.Error("Unexpected call to Provider.Get")
			return
		},
		GetIntHook: func(string) (ident2 int, ident3 merry.Error) {
			t.Error("Unexpected call to Provider.GetInt")
			return
		},
		GetBoolHook: func(string) (ident2 bool, ident3 merry.Error) {
			t.Error("Unexpected call to Provider.GetBool")
			return
		},
		MonitorHook: func(string, <-chan Value) {
			t.Error("Unexpected call to Provider.Monitor")
			return
		},
		HealthcheckHook: func() (ident1 merry.Error) {
			t.Error("Unexpected call to Provider.Healthcheck")
			return
		},
	}
}

func (f *FakeProvider) Reset() {
	f.GetCalls = []*ProviderGetInvocation{}
	f.GetIntCalls = []*ProviderGetIntInvocation{}
	f.GetBoolCalls = []*ProviderGetBoolInvocation{}
	f.MonitorCalls = []*ProviderMonitorInvocation{}
	f.HealthcheckCalls = []*ProviderHealthcheckInvocation{}
}

func (_f1 *FakeProvider) Get(ident1 string) (ident2 string, ident3 merry.Error) {
	invocation := new(ProviderGetInvocation)

	invocation.Parameters.Ident1 = ident1

	ident2, ident3 = _f1.GetHook(ident1)

	invocation.Results.Ident2 = ident2
	invocation.Results.Ident3 = ident3

	_f1.GetCalls = append(_f1.GetCalls, invocation)

	return
}

// GetCalled returns true if FakeProvider.Get was called
func (f *FakeProvider) GetCalled() bool {
	return len(f.GetCalls) != 0
}

// AssertGetCalled calls t.Error if FakeProvider.Get was not called
func (f *FakeProvider) AssertGetCalled(t ProviderTestingT) {
	t.Helper()
	if len(f.GetCalls) == 0 {
		t.Error("FakeProvider.Get not called, expected at least one")
	}
}

// GetNotCalled returns true if FakeProvider.Get was not called
func (f *FakeProvider) GetNotCalled() bool {
	return len(f.GetCalls) == 0
}

// AssertGetNotCalled calls t.Error if FakeProvider.Get was called
func (f *FakeProvider) AssertGetNotCalled(t ProviderTestingT) {
	t.Helper()
	if len(f.GetCalls) != 0 {
		t.Error("FakeProvider.Get called, expected none")
	}
}

// GetCalledOnce returns true if FakeProvider.Get was called exactly once
func (f *FakeProvider) GetCalledOnce() bool {
	return len(f.GetCalls) == 1
}

// AssertGetCalledOnce calls t.Error if FakeProvider.Get was not called exactly once
func (f *FakeProvider) AssertGetCalledOnce(t ProviderTestingT) {
	t.Helper()
	if len(f.GetCalls) != 1 {
		t.Errorf("FakeProvider.Get called %d times, expected 1", len(f.GetCalls))
	}
}

// GetCalledN returns true if FakeProvider.Get was called at least n times
func (f *FakeProvider) GetCalledN(n int) bool {
	return len(f.GetCalls) >= n
}

// AssertGetCalledN calls t.Error if FakeProvider.Get was called less than n times
func (f *FakeProvider) AssertGetCalledN(t ProviderTestingT, n int) {
	t.Helper()
	if len(f.GetCalls) < n {
		t.Errorf("FakeProvider.Get called %d times, expected >= %d", len(f.GetCalls), n)
	}
}

// GetCalledWith returns true if FakeProvider.Get was called with the given values
func (_f2 *FakeProvider) GetCalledWith(ident1 string) (found bool) {
	for _, call := range _f2.GetCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			found = true
			break
		}
	}

	return
}

// AssertGetCalledWith calls t.Error if FakeProvider.Get was not called with the given values
func (_f3 *FakeProvider) AssertGetCalledWith(t ProviderTestingT, ident1 string) {
	t.Helper()
	var found bool
	for _, call := range _f3.GetCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			found = true
			break
		}
	}

	if !found {
		t.Error("FakeProvider.Get not called with expected parameters")
	}
}

// GetCalledOnceWith returns true if FakeProvider.Get was called exactly once with the given values
func (_f4 *FakeProvider) GetCalledOnceWith(ident1 string) bool {
	var count int
	for _, call := range _f4.GetCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			count++
		}
	}

	return count == 1
}

// AssertGetCalledOnceWith calls t.Error if FakeProvider.Get was not called exactly once with the given values
func (_f5 *FakeProvider) AssertGetCalledOnceWith(t ProviderTestingT, ident1 string) {
	t.Helper()
	var count int
	for _, call := range _f5.GetCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			count++
		}
	}

	if count != 1 {
		t.Errorf("FakeProvider.Get called %d times with expected parameters, expected one", count)
	}
}

// GetResultsForCall returns the result values for the first call to FakeProvider.Get with the given values
func (_f6 *FakeProvider) GetResultsForCall(ident1 string) (ident2 string, ident3 merry.Error, found bool) {
	for _, call := range _f6.GetCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			ident2 = call.Results.Ident2
			ident3 = call.Results.Ident3
			found = true
			break
		}
	}

	return
}

func (_f7 *FakeProvider) GetInt(ident1 string) (ident2 int, ident3 merry.Error) {
	invocation := new(ProviderGetIntInvocation)

	invocation.Parameters.Ident1 = ident1

	ident2, ident3 = _f7.GetIntHook(ident1)

	invocation.Results.Ident2 = ident2
	invocation.Results.Ident3 = ident3

	_f7.GetIntCalls = append(_f7.GetIntCalls, invocation)

	return
}

// GetIntCalled returns true if FakeProvider.GetInt was called
func (f *FakeProvider) GetIntCalled() bool {
	return len(f.GetIntCalls) != 0
}

// AssertGetIntCalled calls t.Error if FakeProvider.GetInt was not called
func (f *FakeProvider) AssertGetIntCalled(t ProviderTestingT) {
	t.Helper()
	if len(f.GetIntCalls) == 0 {
		t.Error("FakeProvider.GetInt not called, expected at least one")
	}
}

// GetIntNotCalled returns true if FakeProvider.GetInt was not called
func (f *FakeProvider) GetIntNotCalled() bool {
	return len(f.GetIntCalls) == 0
}

// AssertGetIntNotCalled calls t.Error if FakeProvider.GetInt was called
func (f *FakeProvider) AssertGetIntNotCalled(t ProviderTestingT) {
	t.Helper()
	if len(f.GetIntCalls) != 0 {
		t.Error("FakeProvider.GetInt called, expected none")
	}
}

// GetIntCalledOnce returns true if FakeProvider.GetInt was called exactly once
func (f *FakeProvider) GetIntCalledOnce() bool {
	return len(f.GetIntCalls) == 1
}

// AssertGetIntCalledOnce calls t.Error if FakeProvider.GetInt was not called exactly once
func (f *FakeProvider) AssertGetIntCalledOnce(t ProviderTestingT) {
	t.Helper()
	if len(f.GetIntCalls) != 1 {
		t.Errorf("FakeProvider.GetInt called %d times, expected 1", len(f.GetIntCalls))
	}
}

// GetIntCalledN returns true if FakeProvider.GetInt was called at least n times
func (f *FakeProvider) GetIntCalledN(n int) bool {
	return len(f.GetIntCalls) >= n
}

// AssertGetIntCalledN calls t.Error if FakeProvider.GetInt was called less than n times
func (f *FakeProvider) AssertGetIntCalledN(t ProviderTestingT, n int) {
	t.Helper()
	if len(f.GetIntCalls) < n {
		t.Errorf("FakeProvider.GetInt called %d times, expected >= %d", len(f.GetIntCalls), n)
	}
}

// GetIntCalledWith returns true if FakeProvider.GetInt was called with the given values
func (_f8 *FakeProvider) GetIntCalledWith(ident1 string) (found bool) {
	for _, call := range _f8.GetIntCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			found = true
			break
		}
	}

	return
}

// AssertGetIntCalledWith calls t.Error if FakeProvider.GetInt was not called with the given values
func (_f9 *FakeProvider) AssertGetIntCalledWith(t ProviderTestingT, ident1 string) {
	t.Helper()
	var found bool
	for _, call := range _f9.GetIntCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			found = true
			break
		}
	}

	if !found {
		t.Error("FakeProvider.GetInt not called with expected parameters")
	}
}

// GetIntCalledOnceWith returns true if FakeProvider.GetInt was called exactly once with the given values
func (_f10 *FakeProvider) GetIntCalledOnceWith(ident1 string) bool {
	var count int
	for _, call := range _f10.GetIntCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			count++
		}
	}

	return count == 1
}

// AssertGetIntCalledOnceWith calls t.Error if FakeProvider.GetInt was not called exactly once with the given values
func (_f11 *FakeProvider) AssertGetIntCalledOnceWith(t ProviderTestingT, ident1 string) {
	t.Helper()
	var count int
	for _, call := range _f11.GetIntCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			count++
		}
	}

	if count != 1 {
		t.Errorf("FakeProvider.GetInt called %d times with expected parameters, expected one", count)
	}
}

// GetIntResultsForCall returns the result values for the first call to FakeProvider.GetInt with the given values
func (_f12 *FakeProvider) GetIntResultsForCall(ident1 string) (ident2 int, ident3 merry.Error, found bool) {
	for _, call := range _f12.GetIntCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			ident2 = call.Results.Ident2
			ident3 = call.Results.Ident3
			found = true
			break
		}
	}

	return
}

func (_f13 *FakeProvider) GetBool(ident1 string) (ident2 bool, ident3 merry.Error) {
	invocation := new(ProviderGetBoolInvocation)

	invocation.Parameters.Ident1 = ident1

	ident2, ident3 = _f13.GetBoolHook(ident1)

	invocation.Results.Ident2 = ident2
	invocation.Results.Ident3 = ident3

	_f13.GetBoolCalls = append(_f13.GetBoolCalls, invocation)

	return
}

// GetBoolCalled returns true if FakeProvider.GetBool was called
func (f *FakeProvider) GetBoolCalled() bool {
	return len(f.GetBoolCalls) != 0
}

// AssertGetBoolCalled calls t.Error if FakeProvider.GetBool was not called
func (f *FakeProvider) AssertGetBoolCalled(t ProviderTestingT) {
	t.Helper()
	if len(f.GetBoolCalls) == 0 {
		t.Error("FakeProvider.GetBool not called, expected at least one")
	}
}

// GetBoolNotCalled returns true if FakeProvider.GetBool was not called
func (f *FakeProvider) GetBoolNotCalled() bool {
	return len(f.GetBoolCalls) == 0
}

// AssertGetBoolNotCalled calls t.Error if FakeProvider.GetBool was called
func (f *FakeProvider) AssertGetBoolNotCalled(t ProviderTestingT) {
	t.Helper()
	if len(f.GetBoolCalls) != 0 {
		t.Error("FakeProvider.GetBool called, expected none")
	}
}

// GetBoolCalledOnce returns true if FakeProvider.GetBool was called exactly once
func (f *FakeProvider) GetBoolCalledOnce() bool {
	return len(f.GetBoolCalls) == 1
}

// AssertGetBoolCalledOnce calls t.Error if FakeProvider.GetBool was not called exactly once
func (f *FakeProvider) AssertGetBoolCalledOnce(t ProviderTestingT) {
	t.Helper()
	if len(f.GetBoolCalls) != 1 {
		t.Errorf("FakeProvider.GetBool called %d times, expected 1", len(f.GetBoolCalls))
	}
}

// GetBoolCalledN returns true if FakeProvider.GetBool was called at least n times
func (f *FakeProvider) GetBoolCalledN(n int) bool {
	return len(f.GetBoolCalls) >= n
}

// AssertGetBoolCalledN calls t.Error if FakeProvider.GetBool was called less than n times
func (f *FakeProvider) AssertGetBoolCalledN(t ProviderTestingT, n int) {
	t.Helper()
	if len(f.GetBoolCalls) < n {
		t.Errorf("FakeProvider.GetBool called %d times, expected >= %d", len(f.GetBoolCalls), n)
	}
}

// GetBoolCalledWith returns true if FakeProvider.GetBool was called with the given values
func (_f14 *FakeProvider) GetBoolCalledWith(ident1 string) (found bool) {
	for _, call := range _f14.GetBoolCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			found = true
			break
		}
	}

	return
}

// AssertGetBoolCalledWith calls t.Error if FakeProvider.GetBool was not called with the given values
func (_f15 *FakeProvider) AssertGetBoolCalledWith(t ProviderTestingT, ident1 string) {
	t.Helper()
	var found bool
	for _, call := range _f15.GetBoolCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			found = true
			break
		}
	}

	if !found {
		t.Error("FakeProvider.GetBool not called with expected parameters")
	}
}

// GetBoolCalledOnceWith returns true if FakeProvider.GetBool was called exactly once with the given values
func (_f16 *FakeProvider) GetBoolCalledOnceWith(ident1 string) bool {
	var count int
	for _, call := range _f16.GetBoolCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			count++
		}
	}

	return count == 1
}

// AssertGetBoolCalledOnceWith calls t.Error if FakeProvider.GetBool was not called exactly once with the given values
func (_f17 *FakeProvider) AssertGetBoolCalledOnceWith(t ProviderTestingT, ident1 string) {
	t.Helper()
	var count int
	for _, call := range _f17.GetBoolCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			count++
		}
	}

	if count != 1 {
		t.Errorf("FakeProvider.GetBool called %d times with expected parameters, expected one", count)
	}
}

// GetBoolResultsForCall returns the result values for the first call to FakeProvider.GetBool with the given values
func (_f18 *FakeProvider) GetBoolResultsForCall(ident1 string) (ident2 bool, ident3 merry.Error, found bool) {
	for _, call := range _f18.GetBoolCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			ident2 = call.Results.Ident2
			ident3 = call.Results.Ident3
			found = true
			break
		}
	}

	return
}

func (_f19 *FakeProvider) Monitor(ident1 string, ident2 <-chan Value) {
	invocation := new(ProviderMonitorInvocation)

	invocation.Parameters.Ident1 = ident1
	invocation.Parameters.Ident2 = ident2

	_f19.MonitorHook(ident1, ident2)

	_f19.MonitorCalls = append(_f19.MonitorCalls, invocation)

	return
}

// MonitorCalled returns true if FakeProvider.Monitor was called
func (f *FakeProvider) MonitorCalled() bool {
	return len(f.MonitorCalls) != 0
}

// AssertMonitorCalled calls t.Error if FakeProvider.Monitor was not called
func (f *FakeProvider) AssertMonitorCalled(t ProviderTestingT) {
	t.Helper()
	if len(f.MonitorCalls) == 0 {
		t.Error("FakeProvider.Monitor not called, expected at least one")
	}
}

// MonitorNotCalled returns true if FakeProvider.Monitor was not called
func (f *FakeProvider) MonitorNotCalled() bool {
	return len(f.MonitorCalls) == 0
}

// AssertMonitorNotCalled calls t.Error if FakeProvider.Monitor was called
func (f *FakeProvider) AssertMonitorNotCalled(t ProviderTestingT) {
	t.Helper()
	if len(f.MonitorCalls) != 0 {
		t.Error("FakeProvider.Monitor called, expected none")
	}
}

// MonitorCalledOnce returns true if FakeProvider.Monitor was called exactly once
func (f *FakeProvider) MonitorCalledOnce() bool {
	return len(f.MonitorCalls) == 1
}

// AssertMonitorCalledOnce calls t.Error if FakeProvider.Monitor was not called exactly once
func (f *FakeProvider) AssertMonitorCalledOnce(t ProviderTestingT) {
	t.Helper()
	if len(f.MonitorCalls) != 1 {
		t.Errorf("FakeProvider.Monitor called %d times, expected 1", len(f.MonitorCalls))
	}
}

// MonitorCalledN returns true if FakeProvider.Monitor was called at least n times
func (f *FakeProvider) MonitorCalledN(n int) bool {
	return len(f.MonitorCalls) >= n
}

// AssertMonitorCalledN calls t.Error if FakeProvider.Monitor was called less than n times
func (f *FakeProvider) AssertMonitorCalledN(t ProviderTestingT, n int) {
	t.Helper()
	if len(f.MonitorCalls) < n {
		t.Errorf("FakeProvider.Monitor called %d times, expected >= %d", len(f.MonitorCalls), n)
	}
}

// MonitorCalledWith returns true if FakeProvider.Monitor was called with the given values
func (_f20 *FakeProvider) MonitorCalledWith(ident1 string, ident2 <-chan Value) (found bool) {
	for _, call := range _f20.MonitorCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) && reflect.DeepEqual(call.Parameters.Ident2, ident2) {
			found = true
			break
		}
	}

	return
}

// AssertMonitorCalledWith calls t.Error if FakeProvider.Monitor was not called with the given values
func (_f21 *FakeProvider) AssertMonitorCalledWith(t ProviderTestingT, ident1 string, ident2 <-chan Value) {
	t.Helper()
	var found bool
	for _, call := range _f21.MonitorCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) && reflect.DeepEqual(call.Parameters.Ident2, ident2) {
			found = true
			break
		}
	}

	if !found {
		t.Error("FakeProvider.Monitor not called with expected parameters")
	}
}

// MonitorCalledOnceWith returns true if FakeProvider.Monitor was called exactly once with the given values
func (_f22 *FakeProvider) MonitorCalledOnceWith(ident1 string, ident2 <-chan Value) bool {
	var count int
	for _, call := range _f22.MonitorCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) && reflect.DeepEqual(call.Parameters.Ident2, ident2) {
			count++
		}
	}

	return count == 1
}

// AssertMonitorCalledOnceWith calls t.Error if FakeProvider.Monitor was not called exactly once with the given values
func (_f23 *FakeProvider) AssertMonitorCalledOnceWith(t ProviderTestingT, ident1 string, ident2 <-chan Value) {
	t.Helper()
	var count int
	for _, call := range _f23.MonitorCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) && reflect.DeepEqual(call.Parameters.Ident2, ident2) {
			count++
		}
	}

	if count != 1 {
		t.Errorf("FakeProvider.Monitor called %d times with expected parameters, expected one", count)
	}
}

func (_f24 *FakeProvider) Healthcheck() (ident1 merry.Error) {
	invocation := new(ProviderHealthcheckInvocation)

	ident1 = _f24.HealthcheckHook()

	invocation.Results.Ident1 = ident1

	_f24.HealthcheckCalls = append(_f24.HealthcheckCalls, invocation)

	return
}

// HealthcheckCalled returns true if FakeProvider.Healthcheck was called
func (f *FakeProvider) HealthcheckCalled() bool {
	return len(f.HealthcheckCalls) != 0
}

// AssertHealthcheckCalled calls t.Error if FakeProvider.Healthcheck was not called
func (f *FakeProvider) AssertHealthcheckCalled(t ProviderTestingT) {
	t.Helper()
	if len(f.HealthcheckCalls) == 0 {
		t.Error("FakeProvider.Healthcheck not called, expected at least one")
	}
}

// HealthcheckNotCalled returns true if FakeProvider.Healthcheck was not called
func (f *FakeProvider) HealthcheckNotCalled() bool {
	return len(f.HealthcheckCalls) == 0
}

// AssertHealthcheckNotCalled calls t.Error if FakeProvider.Healthcheck was called
func (f *FakeProvider) AssertHealthcheckNotCalled(t ProviderTestingT) {
	t.Helper()
	if len(f.HealthcheckCalls) != 0 {
		t.Error("FakeProvider.Healthcheck called, expected none")
	}
}

// HealthcheckCalledOnce returns true if FakeProvider.Healthcheck was called exactly once
func (f *FakeProvider) HealthcheckCalledOnce() bool {
	return len(f.HealthcheckCalls) == 1
}

// AssertHealthcheckCalledOnce calls t.Error if FakeProvider.Healthcheck was not called exactly once
func (f *FakeProvider) AssertHealthcheckCalledOnce(t ProviderTestingT) {
	t.Helper()
	if len(f.HealthcheckCalls) != 1 {
		t.Errorf("FakeProvider.Healthcheck called %d times, expected 1", len(f.HealthcheckCalls))
	}
}

// HealthcheckCalledN returns true if FakeProvider.Healthcheck was called at least n times
func (f *FakeProvider) HealthcheckCalledN(n int) bool {
	return len(f.HealthcheckCalls) >= n
}

// AssertHealthcheckCalledN calls t.Error if FakeProvider.Healthcheck was called less than n times
func (f *FakeProvider) AssertHealthcheckCalledN(t ProviderTestingT, n int) {
	t.Helper()
	if len(f.HealthcheckCalls) < n {
		t.Errorf("FakeProvider.Healthcheck called %d times, expected >= %d", len(f.HealthcheckCalls), n)
	}
}
