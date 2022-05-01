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

// NearTopN provides a mock function with given fields: ctx, top, latitude, longitude
func (_m *IBiz) NearTopN(ctx contextx.Contextx, top int, latitude float64, longitude float64) ([]*pb.Car, int, error) {
	ret := _m.Called(ctx, top, latitude, longitude)

	var r0 []*pb.Car
	if rf, ok := ret.Get(0).(func(contextx.Contextx, int, float64, float64) []*pb.Car); ok {
		r0 = rf(ctx, top, latitude, longitude)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*pb.Car)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(contextx.Contextx, int, float64, float64) int); ok {
		r1 = rf(ctx, top, latitude, longitude)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(contextx.Contextx, int, float64, float64) error); ok {
		r2 = rf(ctx, top, latitude, longitude)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}