// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	contextx "github.com/blackhorseya/gocommon/pkg/contextx"
	mock "github.com/stretchr/testify/mock"

	order "github.com/blackhorseya/irent/internal/pkg/entity/order"

	user "github.com/blackhorseya/irent/internal/pkg/entity/user"
)

// IRepo is an autogenerated mock type for the IRepo type
type IRepo struct {
	mock.Mock
}

// Book provides a mock function with given fields: ctx, id, projID, from
func (_m *IRepo) Book(ctx contextx.Contextx, id string, projID string, from *user.Profile) (*order.Booking, error) {
	ret := _m.Called(ctx, id, projID, from)

	var r0 *order.Booking
	if rf, ok := ret.Get(0).(func(contextx.Contextx, string, string, *user.Profile) *order.Booking); ok {
		r0 = rf(ctx, id, projID, from)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*order.Booking)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, string, string, *user.Profile) error); ok {
		r1 = rf(ctx, id, projID, from)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CancelBooking provides a mock function with given fields: ctx, id, from
func (_m *IRepo) CancelBooking(ctx contextx.Contextx, id string, from *user.Profile) error {
	ret := _m.Called(ctx, id, from)

	var r0 error
	if rf, ok := ret.Get(0).(func(contextx.Contextx, string, *user.Profile) error); ok {
		r0 = rf(ctx, id, from)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueryBookings provides a mock function with given fields: ctx, from
func (_m *IRepo) QueryBookings(ctx contextx.Contextx, from *user.Profile) ([]*order.Info, error) {
	ret := _m.Called(ctx, from)

	var r0 []*order.Info
	if rf, ok := ret.Get(0).(func(contextx.Contextx, *user.Profile) []*order.Info); ok {
		r0 = rf(ctx, from)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*order.Info)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, *user.Profile) error); ok {
		r1 = rf(ctx, from)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewIRepoT interface {
	mock.TestingT
	Cleanup(func())
}

// NewIRepo creates a new instance of IRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIRepo(t NewIRepoT) *IRepo {
	mock := &IRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
