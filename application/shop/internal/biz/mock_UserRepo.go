// Code generated by mockery v2.41.0. DO NOT EDIT.

package biz

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockUserRepo is an autogenerated mock type for the UserRepo type
type MockUserRepo struct {
	mock.Mock
}

type MockUserRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUserRepo) EXPECT() *MockUserRepo_Expecter {
	return &MockUserRepo_Expecter{mock: &_m.Mock}
}

// CheckPassword provides a mock function with given fields: ctx, password, encryptedPassword
func (_m *MockUserRepo) CheckPassword(ctx context.Context, password string, encryptedPassword string) (bool, error) {
	ret := _m.Called(ctx, password, encryptedPassword)

	if len(ret) == 0 {
		panic("no return value specified for CheckPassword")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (bool, error)); ok {
		return rf(ctx, password, encryptedPassword)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) bool); ok {
		r0 = rf(ctx, password, encryptedPassword)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, password, encryptedPassword)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserRepo_CheckPassword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckPassword'
type MockUserRepo_CheckPassword_Call struct {
	*mock.Call
}

// CheckPassword is a helper method to define mock.On call
//   - ctx context.Context
//   - password string
//   - encryptedPassword string
func (_e *MockUserRepo_Expecter) CheckPassword(ctx interface{}, password interface{}, encryptedPassword interface{}) *MockUserRepo_CheckPassword_Call {
	return &MockUserRepo_CheckPassword_Call{Call: _e.mock.On("CheckPassword", ctx, password, encryptedPassword)}
}

func (_c *MockUserRepo_CheckPassword_Call) Run(run func(ctx context.Context, password string, encryptedPassword string)) *MockUserRepo_CheckPassword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockUserRepo_CheckPassword_Call) Return(_a0 bool, _a1 error) *MockUserRepo_CheckPassword_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserRepo_CheckPassword_Call) RunAndReturn(run func(context.Context, string, string) (bool, error)) *MockUserRepo_CheckPassword_Call {
	_c.Call.Return(run)
	return _c
}

// CreateUser provides a mock function with given fields: c, u
func (_m *MockUserRepo) CreateUser(c context.Context, u *User) (*User, error) {
	ret := _m.Called(c, u)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 *User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *User) (*User, error)); ok {
		return rf(c, u)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *User) *User); ok {
		r0 = rf(c, u)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *User) error); ok {
		r1 = rf(c, u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserRepo_CreateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUser'
type MockUserRepo_CreateUser_Call struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//   - c context.Context
//   - u *User
func (_e *MockUserRepo_Expecter) CreateUser(c interface{}, u interface{}) *MockUserRepo_CreateUser_Call {
	return &MockUserRepo_CreateUser_Call{Call: _e.mock.On("CreateUser", c, u)}
}

func (_c *MockUserRepo_CreateUser_Call) Run(run func(c context.Context, u *User)) *MockUserRepo_CreateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*User))
	})
	return _c
}

func (_c *MockUserRepo_CreateUser_Call) Return(_a0 *User, _a1 error) *MockUserRepo_CreateUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserRepo_CreateUser_Call) RunAndReturn(run func(context.Context, *User) (*User, error)) *MockUserRepo_CreateUser_Call {
	_c.Call.Return(run)
	return _c
}

// UserById provides a mock function with given fields: ctx, Id
func (_m *MockUserRepo) UserById(ctx context.Context, Id int64) (*User, error) {
	ret := _m.Called(ctx, Id)

	if len(ret) == 0 {
		panic("no return value specified for UserById")
	}

	var r0 *User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*User, error)); ok {
		return rf(ctx, Id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *User); ok {
		r0 = rf(ctx, Id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, Id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserRepo_UserById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UserById'
type MockUserRepo_UserById_Call struct {
	*mock.Call
}

// UserById is a helper method to define mock.On call
//   - ctx context.Context
//   - Id int64
func (_e *MockUserRepo_Expecter) UserById(ctx interface{}, Id interface{}) *MockUserRepo_UserById_Call {
	return &MockUserRepo_UserById_Call{Call: _e.mock.On("UserById", ctx, Id)}
}

func (_c *MockUserRepo_UserById_Call) Run(run func(ctx context.Context, Id int64)) *MockUserRepo_UserById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockUserRepo_UserById_Call) Return(_a0 *User, _a1 error) *MockUserRepo_UserById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserRepo_UserById_Call) RunAndReturn(run func(context.Context, int64) (*User, error)) *MockUserRepo_UserById_Call {
	_c.Call.Return(run)
	return _c
}

// UserByMobile provides a mock function with given fields: ctx, mobile
func (_m *MockUserRepo) UserByMobile(ctx context.Context, mobile string) (*User, error) {
	ret := _m.Called(ctx, mobile)

	if len(ret) == 0 {
		panic("no return value specified for UserByMobile")
	}

	var r0 *User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*User, error)); ok {
		return rf(ctx, mobile)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *User); ok {
		r0 = rf(ctx, mobile)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, mobile)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserRepo_UserByMobile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UserByMobile'
type MockUserRepo_UserByMobile_Call struct {
	*mock.Call
}

// UserByMobile is a helper method to define mock.On call
//   - ctx context.Context
//   - mobile string
func (_e *MockUserRepo_Expecter) UserByMobile(ctx interface{}, mobile interface{}) *MockUserRepo_UserByMobile_Call {
	return &MockUserRepo_UserByMobile_Call{Call: _e.mock.On("UserByMobile", ctx, mobile)}
}

func (_c *MockUserRepo_UserByMobile_Call) Run(run func(ctx context.Context, mobile string)) *MockUserRepo_UserByMobile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockUserRepo_UserByMobile_Call) Return(_a0 *User, _a1 error) *MockUserRepo_UserByMobile_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserRepo_UserByMobile_Call) RunAndReturn(run func(context.Context, string) (*User, error)) *MockUserRepo_UserByMobile_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUserRepo creates a new instance of MockUserRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUserRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUserRepo {
	mock := &MockUserRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}