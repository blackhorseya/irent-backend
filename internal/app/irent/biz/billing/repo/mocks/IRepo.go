// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	contextx "github.com/blackhorseya/gocommon/pkg/contextx"
	mock "github.com/stretchr/testify/mock"

	pb "github.com/blackhorseya/irent/pb"

	testing "testing"

	user "github.com/blackhorseya/irent/internal/pkg/entity/user"
)

// IRepo is an autogenerated mock type for the IRepo type
type IRepo struct {
	mock.Mock
}

// QueryArrears provides a mock function with given fields: ctx, _a1
func (_m *IRepo) QueryArrears(ctx contextx.Contextx, _a1 *user.Profile) (*pb.Arrears, error) {
	ret := _m.Called(ctx, _a1)

	var r0 *pb.Arrears
	if rf, ok := ret.Get(0).(func(contextx.Contextx, *user.Profile) *pb.Arrears); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Arrears)
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

// NewIRepo creates a new instance of IRepo. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewIRepo(t testing.TB) *IRepo {
	mock := &IRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
