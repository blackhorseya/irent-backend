// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	contextx "github.com/blackhorseya/irent/internal/pkg/base/contextx"
	mock "github.com/stretchr/testify/mock"

	pb "github.com/blackhorseya/irent/pb"
)

// IBiz is an autogenerated mock type for the IBiz type
type IBiz struct {
	mock.Mock
}

// GetByAccessToken provides a mock function with given fields: ctx, token
func (_m *IBiz) GetByAccessToken(ctx contextx.Contextx, token string) (*pb.Profile, error) {
	ret := _m.Called(ctx, token)

	var r0 *pb.Profile
	if rf, ok := ret.Get(0).(func(contextx.Contextx, string) *pb.Profile); ok {
		r0 = rf(ctx, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Profile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, string) error); ok {
		r1 = rf(ctx, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *IBiz) GetByID(ctx contextx.Contextx, id string) (*pb.Profile, error) {
	ret := _m.Called(ctx, id)

	var r0 *pb.Profile
	if rf, ok := ret.Get(0).(func(contextx.Contextx, string) *pb.Profile); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Profile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: ctx, id, password
func (_m *IBiz) Login(ctx contextx.Contextx, id string, password string) (*pb.Profile, error) {
	ret := _m.Called(ctx, id, password)

	var r0 *pb.Profile
	if rf, ok := ret.Get(0).(func(contextx.Contextx, string, string) *pb.Profile); ok {
		r0 = rf(ctx, id, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Profile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, string, string) error); ok {
		r1 = rf(ctx, id, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Logout provides a mock function with given fields: ctx, _a1
func (_m *IBiz) Logout(ctx contextx.Contextx, _a1 *pb.Profile) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(contextx.Contextx, *pb.Profile) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RefreshToken provides a mock function with given fields: ctx, _a1
func (_m *IBiz) RefreshToken(ctx contextx.Contextx, _a1 *pb.Profile) (*pb.Profile, error) {
	ret := _m.Called(ctx, _a1)

	var r0 *pb.Profile
	if rf, ok := ret.Get(0).(func(contextx.Contextx, *pb.Profile) *pb.Profile); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Profile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, *pb.Profile) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
