// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	contextx "github.com/blackhorseya/gocommon/pkg/contextx"
	mock "github.com/stretchr/testify/mock"

	pb "github.com/blackhorseya/irent/pb"
)

// IRepo is an autogenerated mock type for the IRepo type
type IRepo struct {
	mock.Mock
}

// QueryArrears provides a mock function with given fields: ctx, user
func (_m *IRepo) QueryArrears(ctx contextx.Contextx, user *pb.Profile) (*pb.Arrears, error) {
	ret := _m.Called(ctx, user)

	var r0 *pb.Arrears
	if rf, ok := ret.Get(0).(func(contextx.Contextx, *pb.Profile) *pb.Arrears); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Arrears)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, *pb.Profile) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
