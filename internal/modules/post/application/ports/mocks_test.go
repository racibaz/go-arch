// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package ports

import (
	"github.com/racibaz/go-arch/internal/modules/post/application/usecases/inputs"
	"github.com/racibaz/go-arch/internal/modules/post/domain"
	mock "github.com/stretchr/testify/mock"
)

// NewMockPostService creates a new instance of MockPostService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPostService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPostService {
	mock := &MockPostService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// MockPostService is an autogenerated mock type for the PostService type
type MockPostService struct {
	mock.Mock
}

type MockPostService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPostService) EXPECT() *MockPostService_Expecter {
	return &MockPostService_Expecter{mock: &_m.Mock}
}

// Create provides a mock function for the type MockPostService
func (_mock *MockPostService) Create(postDto inputs.CreatePostInput) error {
	ret := _mock.Called(postDto)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if returnFunc, ok := ret.Get(0).(func(inputs.CreatePostInput) error); ok {
		r0 = returnFunc(postDto)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// MockPostService_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockPostService_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - postDto inputs.CreatePostInput
func (_e *MockPostService_Expecter) Create(postDto interface{}) *MockPostService_Create_Call {
	return &MockPostService_Create_Call{Call: _e.mock.On("Create", postDto)}
}

func (_c *MockPostService_Create_Call) Run(run func(postDto inputs.CreatePostInput)) *MockPostService_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 inputs.CreatePostInput
		if args[0] != nil {
			arg0 = args[0].(inputs.CreatePostInput)
		}
		run(
			arg0,
		)
	})
	return _c
}

func (_c *MockPostService_Create_Call) Return(err error) *MockPostService_Create_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockPostService_Create_Call) RunAndReturn(run func(postDto inputs.CreatePostInput) error) *MockPostService_Create_Call {
	_c.Call.Return(run)
	return _c
}

// GetById provides a mock function for the type MockPostService
func (_mock *MockPostService) GetById(id string) (*domain.Post, error) {
	ret := _mock.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 *domain.Post
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(string) (*domain.Post, error)); ok {
		return returnFunc(id)
	}
	if returnFunc, ok := ret.Get(0).(func(string) *domain.Post); ok {
		r0 = returnFunc(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Post)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(string) error); ok {
		r1 = returnFunc(id)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockPostService_GetById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetById'
type MockPostService_GetById_Call struct {
	*mock.Call
}

// GetById is a helper method to define mock.On call
//   - id string
func (_e *MockPostService_Expecter) GetById(id interface{}) *MockPostService_GetById_Call {
	return &MockPostService_GetById_Call{Call: _e.mock.On("GetById", id)}
}

func (_c *MockPostService_GetById_Call) Run(run func(id string)) *MockPostService_GetById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 string
		if args[0] != nil {
			arg0 = args[0].(string)
		}
		run(
			arg0,
		)
	})
	return _c
}

func (_c *MockPostService_GetById_Call) Return(post *domain.Post, err error) *MockPostService_GetById_Call {
	_c.Call.Return(post, err)
	return _c
}

func (_c *MockPostService_GetById_Call) RunAndReturn(run func(id string) (*domain.Post, error)) *MockPostService_GetById_Call {
	_c.Call.Return(run)
	return _c
}
