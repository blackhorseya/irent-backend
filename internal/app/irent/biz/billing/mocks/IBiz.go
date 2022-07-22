// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	contextx "github.com/blackhorseya/gocommon/pkg/contextx"
	mock "github.com/stretchr/testify/mock"

	user "github.com/blackhorseya/irent/internal/pkg/entity/user"
)

// IBiz is an autogenerated mock type for the IBiz type
type IBiz struct {
	mock.Mock
}

// GetArrears provides a mock function with given fields: ctx, _a1
func (_m *IBiz) GetArrears(ctx contextx.Contextx, _a1 *user.Profile) (*user.Arrears, error) {
	ret := _m.Called(ctx, _a1)

	var r0 *user.Arrears
	if rf, ok := ret.Get(0).(func(contextx.Contextx, *user.Profile) *user.Arrears); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.Arrears)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, *user.Profile) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIBiz interface {
	mock.TestingT
	Cleanup(func())
}

// NewIBiz creates a new instance of IBiz. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIBiz(t mockConstructorTestingTNewIBiz) *IBiz {
	mock := &IBiz{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
