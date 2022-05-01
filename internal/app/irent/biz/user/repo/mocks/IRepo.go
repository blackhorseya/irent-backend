// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	contextx "github.com/blackhorseya/irent/internal/pkg/base/contextx"
	mock "github.com/stretchr/testify/mock"

	pb "github.com/blackhorseya/irent/pb"
)

// IRepo is an autogenerated mock type for the IRepo type
type IRepo struct {
	mock.Mock
}

// Login provides a mock function with given fields: ctx, id, password
func (_m *IRepo) Login(ctx contextx.Contextx, id string, password string) (*pb.Profile, error) {
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