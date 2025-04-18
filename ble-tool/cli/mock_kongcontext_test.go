// Code generated by pegomock. DO NOT EDIT.
// Source: ble-tool/cli (interfaces: KongContext)

package cli_test

import (
	pegomock "github.com/petergtz/pegomock/v4"
	"reflect"
	"time"
)

type MockKongContext struct {
	fail func(message string, callerSkip ...int)
}

func NewMockKongContext(options ...pegomock.Option) *MockKongContext {
	mock := &MockKongContext{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockKongContext) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockKongContext) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockKongContext) Command() string {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockKongContext().")
	}
	_params := []pegomock.Param{}
	_result := pegomock.GetGenericMockFrom(mock).Invoke("Command", _params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem()})
	var _ret0 string
	if len(_result) != 0 {
		if _result[0] != nil {
			_ret0 = _result[0].(string)
		}
	}
	return _ret0
}

func (mock *MockKongContext) VerifyWasCalledOnce() *VerifierMockKongContext {
	return &VerifierMockKongContext{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockKongContext) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockKongContext {
	return &VerifierMockKongContext{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockKongContext) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockKongContext {
	return &VerifierMockKongContext{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockKongContext) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockKongContext {
	return &VerifierMockKongContext{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockKongContext struct {
	mock                   *MockKongContext
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockKongContext) Command() *MockKongContext_Command_OngoingVerification {
	_params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Command", _params, verifier.timeout)
	return &MockKongContext_Command_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockKongContext_Command_OngoingVerification struct {
	mock              *MockKongContext
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockKongContext_Command_OngoingVerification) GetCapturedArguments() {
}

func (c *MockKongContext_Command_OngoingVerification) GetAllCapturedArguments() {
}
