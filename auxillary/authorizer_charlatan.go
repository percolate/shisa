// generated by "charlatan -output=./authorizer_charlatan.go Authorizer".  DO NOT EDIT.

package auxillary

import "reflect"
import "github.com/ansel1/merry"
import "github.com/percolate/shisa/context"
import "github.com/percolate/shisa/service"

// AuthorizerAuthorizeInvocation represents a single call of FakeAuthorizer.Authorize
type AuthorizerAuthorizeInvocation struct {
	Parameters struct {
		Ident1 context.Context
		Ident2 *service.Request
	}
	Results struct {
		Ident3 bool
		Ident4 merry.Error
	}
}

// AuthorizerTestingT represents the methods of "testing".T used by charlatan Fakes.  It avoids importing the testing package.
type AuthorizerTestingT interface {
	Error(...interface{})
	Errorf(string, ...interface{})
	Fatal(...interface{})
	Helper()
}

/*
FakeAuthorizer is a mock implementation of Authorizer for testing.
Use it in your tests as in this example:

	package example

	func TestWithAuthorizer(t *testing.T) {
		f := &auxillary.FakeAuthorizer{
			AuthorizeHook: func(ident1 context.Context, ident2 *service.Request) (ident3 bool, ident4 merry.Error) {
				// ensure parameters meet expections, signal errors using t, etc
				return
			},
		}

		// test code goes here ...

		// assert state of FakeAuthorize ...
		f.AssertAuthorizeCalledOnce(t)
	}

Create anonymous function implementations for only those interface methods that
should be called in the code under test.  This will force a panic if any
unexpected calls are made to FakeAuthorize.
*/
type FakeAuthorizer struct {
	AuthorizeHook func(context.Context, *service.Request) (bool, merry.Error)

	AuthorizeCalls []*AuthorizerAuthorizeInvocation
}

// NewFakeAuthorizerDefaultPanic returns an instance of FakeAuthorizer with all hooks configured to panic
func NewFakeAuthorizerDefaultPanic() *FakeAuthorizer {
	return &FakeAuthorizer{
		AuthorizeHook: func(context.Context, *service.Request) (ident3 bool, ident4 merry.Error) {
			panic("Unexpected call to Authorizer.Authorize")
		},
	}
}

// NewFakeAuthorizerDefaultFatal returns an instance of FakeAuthorizer with all hooks configured to call t.Fatal
func NewFakeAuthorizerDefaultFatal(t AuthorizerTestingT) *FakeAuthorizer {
	return &FakeAuthorizer{
		AuthorizeHook: func(context.Context, *service.Request) (ident3 bool, ident4 merry.Error) {
			t.Fatal("Unexpected call to Authorizer.Authorize")
			return
		},
	}
}

// NewFakeAuthorizerDefaultError returns an instance of FakeAuthorizer with all hooks configured to call t.Error
func NewFakeAuthorizerDefaultError(t AuthorizerTestingT) *FakeAuthorizer {
	return &FakeAuthorizer{
		AuthorizeHook: func(context.Context, *service.Request) (ident3 bool, ident4 merry.Error) {
			t.Error("Unexpected call to Authorizer.Authorize")
			return
		},
	}
}

func (f *FakeAuthorizer) Reset() {
	f.AuthorizeCalls = []*AuthorizerAuthorizeInvocation{}
}

func (_f1 *FakeAuthorizer) Authorize(ident1 context.Context, ident2 *service.Request) (ident3 bool, ident4 merry.Error) {
	invocation := new(AuthorizerAuthorizeInvocation)

	invocation.Parameters.Ident1 = ident1
	invocation.Parameters.Ident2 = ident2

	ident3, ident4 = _f1.AuthorizeHook(ident1, ident2)

	invocation.Results.Ident3 = ident3
	invocation.Results.Ident4 = ident4

	_f1.AuthorizeCalls = append(_f1.AuthorizeCalls, invocation)

	return
}

// AuthorizeCalled returns true if FakeAuthorizer.Authorize was called
func (f *FakeAuthorizer) AuthorizeCalled() bool {
	return len(f.AuthorizeCalls) != 0
}

// AssertAuthorizeCalled calls t.Error if FakeAuthorizer.Authorize was not called
func (f *FakeAuthorizer) AssertAuthorizeCalled(t AuthorizerTestingT) {
	t.Helper()
	if len(f.AuthorizeCalls) == 0 {
		t.Error("FakeAuthorizer.Authorize not called, expected at least one")
	}
}

// AuthorizeNotCalled returns true if FakeAuthorizer.Authorize was not called
func (f *FakeAuthorizer) AuthorizeNotCalled() bool {
	return len(f.AuthorizeCalls) == 0
}

// AssertAuthorizeNotCalled calls t.Error if FakeAuthorizer.Authorize was called
func (f *FakeAuthorizer) AssertAuthorizeNotCalled(t AuthorizerTestingT) {
	t.Helper()
	if len(f.AuthorizeCalls) != 0 {
		t.Error("FakeAuthorizer.Authorize called, expected none")
	}
}

// AuthorizeCalledOnce returns true if FakeAuthorizer.Authorize was called exactly once
func (f *FakeAuthorizer) AuthorizeCalledOnce() bool {
	return len(f.AuthorizeCalls) == 1
}

// AssertAuthorizeCalledOnce calls t.Error if FakeAuthorizer.Authorize was not called exactly once
func (f *FakeAuthorizer) AssertAuthorizeCalledOnce(t AuthorizerTestingT) {
	t.Helper()
	if len(f.AuthorizeCalls) != 1 {
		t.Errorf("FakeAuthorizer.Authorize called %d times, expected 1", len(f.AuthorizeCalls))
	}
}

// AuthorizeCalledN returns true if FakeAuthorizer.Authorize was called at least n times
func (f *FakeAuthorizer) AuthorizeCalledN(n int) bool {
	return len(f.AuthorizeCalls) >= n
}

// AssertAuthorizeCalledN calls t.Error if FakeAuthorizer.Authorize was called less than n times
func (f *FakeAuthorizer) AssertAuthorizeCalledN(t AuthorizerTestingT, n int) {
	t.Helper()
	if len(f.AuthorizeCalls) < n {
		t.Errorf("FakeAuthorizer.Authorize called %d times, expected >= %d", len(f.AuthorizeCalls), n)
	}
}

// AuthorizeCalledWith returns true if FakeAuthorizer.Authorize was called with the given values
func (_f2 *FakeAuthorizer) AuthorizeCalledWith(ident1 context.Context, ident2 *service.Request) (found bool) {
	for _, call := range _f2.AuthorizeCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) && reflect.DeepEqual(call.Parameters.Ident2, ident2) {
			found = true
			break
		}
	}

	return
}

// AssertAuthorizeCalledWith calls t.Error if FakeAuthorizer.Authorize was not called with the given values
func (_f3 *FakeAuthorizer) AssertAuthorizeCalledWith(t AuthorizerTestingT, ident1 context.Context, ident2 *service.Request) {
	t.Helper()
	var found bool
	for _, call := range _f3.AuthorizeCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) && reflect.DeepEqual(call.Parameters.Ident2, ident2) {
			found = true
			break
		}
	}

	if !found {
		t.Error("FakeAuthorizer.Authorize not called with expected parameters")
	}
}

// AuthorizeCalledOnceWith returns true if FakeAuthorizer.Authorize was called exactly once with the given values
func (_f4 *FakeAuthorizer) AuthorizeCalledOnceWith(ident1 context.Context, ident2 *service.Request) bool {
	var count int
	for _, call := range _f4.AuthorizeCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) && reflect.DeepEqual(call.Parameters.Ident2, ident2) {
			count++
		}
	}

	return count == 1
}

// AssertAuthorizeCalledOnceWith calls t.Error if FakeAuthorizer.Authorize was not called exactly once with the given values
func (_f5 *FakeAuthorizer) AssertAuthorizeCalledOnceWith(t AuthorizerTestingT, ident1 context.Context, ident2 *service.Request) {
	t.Helper()
	var count int
	for _, call := range _f5.AuthorizeCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) && reflect.DeepEqual(call.Parameters.Ident2, ident2) {
			count++
		}
	}

	if count != 1 {
		t.Errorf("FakeAuthorizer.Authorize called %d times with expected parameters, expected one", count)
	}
}

// AuthorizeResultsForCall returns the result values for the first call to FakeAuthorizer.Authorize with the given values
func (_f6 *FakeAuthorizer) AuthorizeResultsForCall(ident1 context.Context, ident2 *service.Request) (ident3 bool, ident4 merry.Error, found bool) {
	for _, call := range _f6.AuthorizeCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) && reflect.DeepEqual(call.Parameters.Ident2, ident2) {
			ident3 = call.Results.Ident3
			ident4 = call.Results.Ident4
			found = true
			break
		}
	}

	return
}
