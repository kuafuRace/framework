// Code generated by mockery. DO NOT EDIT.

package orm

import (
	orm "github.com/goravel/framework/contracts/database/orm"
	mock "github.com/stretchr/testify/mock"
)

// ObserverWithSaved is an autogenerated mock type for the ObserverWithSaved type
type ObserverWithSaved struct {
	mock.Mock
}

type ObserverWithSaved_Expecter struct {
	mock *mock.Mock
}

func (_m *ObserverWithSaved) EXPECT() *ObserverWithSaved_Expecter {
	return &ObserverWithSaved_Expecter{mock: &_m.Mock}
}

// Saved provides a mock function with given fields: _a0
func (_m *ObserverWithSaved) Saved(_a0 orm.Event) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Saved")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(orm.Event) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ObserverWithSaved_Saved_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Saved'
type ObserverWithSaved_Saved_Call struct {
	*mock.Call
}

// Saved is a helper method to define mock.On call
//   - _a0 orm.Event
func (_e *ObserverWithSaved_Expecter) Saved(_a0 interface{}) *ObserverWithSaved_Saved_Call {
	return &ObserverWithSaved_Saved_Call{Call: _e.mock.On("Saved", _a0)}
}

func (_c *ObserverWithSaved_Saved_Call) Run(run func(_a0 orm.Event)) *ObserverWithSaved_Saved_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(orm.Event))
	})
	return _c
}

func (_c *ObserverWithSaved_Saved_Call) Return(_a0 error) *ObserverWithSaved_Saved_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ObserverWithSaved_Saved_Call) RunAndReturn(run func(orm.Event) error) *ObserverWithSaved_Saved_Call {
	_c.Call.Return(run)
	return _c
}

// NewObserverWithSaved creates a new instance of ObserverWithSaved. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewObserverWithSaved(t interface {
	mock.TestingT
	Cleanup(func())
}) *ObserverWithSaved {
	mock := &ObserverWithSaved{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
