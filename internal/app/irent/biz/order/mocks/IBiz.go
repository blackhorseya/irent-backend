// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	contextx "github.com/blackhorseya/gocommon/pkg/contextx"
	mock "github.com/stretchr/testify/mock"

	order "github.com/blackhorseya/irent/internal/pkg/entity/order"

	user "github.com/blackhorseya/irent/internal/pkg/entity/user"
)

// IBiz is an autogenerated mock type for the IBiz type
type IBiz struct {
	mock.Mock
}

// BookCar provides a mock function with given fields: ctx, id, projID, from, circularly
func (_m *IBiz) BookCar(ctx contextx.Contextx, id string, projID string, from *user.Profile, circularly bool) (*order.Booking, error) {
	ret := _m.Called(ctx, id, projID, from, circularly)

	var r0 *order.Booking
	if rf, ok := ret.Get(0).(func(contextx.Contextx, string, string, *user.Profile, bool) *order.Booking); ok {
		r0 = rf(ctx, id, projID, from, circularly)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*order.Booking)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, string, string, *user.Profile, bool) error); ok {
		r1 = rf(ctx, id, projID, from, circularly)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CancelBooking provides a mock function with given fields: ctx, id, from
func (_m *IBiz) CancelBooking(ctx contextx.Contextx, id string, from *user.Profile) error {
	ret := _m.Called(ctx, id, from)

	var r0 error
	if rf, ok := ret.Get(0).(func(contextx.Contextx, string, *user.Profile) error); ok {
		r0 = rf(ctx, id, from)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByID provides a mock function with given fields: ctx, id, from
func (_m *IBiz) GetByID(ctx contextx.Contextx, id string, from *user.Profile) (*order.Info, error) {
	ret := _m.Called(ctx, id, from)

	var r0 *order.Info
	if rf, ok := ret.Get(0).(func(contextx.Contextx, string, *user.Profile) *order.Info); ok {
		r0 = rf(ctx, id, from)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*order.Info)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, string, *user.Profile) error); ok {
		r1 = rf(ctx, id, from)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, start, end, from
func (_m *IBiz) List(ctx contextx.Contextx, start int, end int, from *user.Profile) ([]*order.Info, error) {
	ret := _m.Called(ctx, start, end, from)

	var r0 []*order.Info
	if rf, ok := ret.Get(0).(func(contextx.Contextx, int, int, *user.Profile) []*order.Info); ok {
		r0 = rf(ctx, start, end, from)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*order.Info)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, int, int, *user.Profile) error); ok {
		r1 = rf(ctx, start, end, from)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPremiumBookings provides a mock function with given fields: ctx
func (_m *IBiz) ListPremiumBookings(ctx contextx.Contextx) (map[*user.Profile]*order.Booking, error) {
	ret := _m.Called(ctx)

	var r0 map[*user.Profile]*order.Booking
	if rf, ok := ret.Get(0).(func(contextx.Contextx) map[*user.Profile]*order.Booking); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[*user.Profile]*order.Booking)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReBookCar provides a mock function with given fields: ctx, no, id, projID, from
func (_m *IBiz) ReBookCar(ctx contextx.Contextx, no string, id string, projID string, from *user.Profile) (*order.Booking, error) {
	ret := _m.Called(ctx, no, id, projID, from)

	var r0 *order.Booking
	if rf, ok := ret.Get(0).(func(contextx.Contextx, string, string, string, *user.Profile) *order.Booking); ok {
		r0 = rf(ctx, no, id, projID, from)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*order.Booking)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, string, string, string, *user.Profile) error); ok {
		r1 = rf(ctx, no, id, projID, from)
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
