package mocks

import (
	"auth-service/internal/repo/model"
	"context"
	"github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}

func (_m *UserRepository) GetByUserName(ctx context.Context, username string) (model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (_m *UserRepository) GetAll(ctx context.Context) ([]model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (_m *UserRepository) Create(ctx context.Context, user model.CreateUser) (string, error) {
	// Create the mock
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 string
	var r1 error

	if rf, ok := ret.Get(0).(func(context.Context, model.CreateUser) (string, error)); ok {
		return rf(ctx, user)
	}

	if rf, ok := ret.Get(0).(func(context.Context, model.CreateUser) string); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.CreateUser) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *UserRepository) Update(ctx context.Context, id string, user model.UpdateUser) error {
	// Update the mock
	ret := _m.Called(ctx, id, user)
	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, model.UpdateUser) error); ok {
		r0 = rf(ctx, id, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *UserRepository) Delete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)
	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *UserRepository) GetById(ctx context.Context, id string) (model.User, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (model.User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) model.User); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//func NewUser(t interface {
//	mock.TestingT
//	Cleanup(func())
//}) *UserRepository {
//	mock := &UserRepository{}
//	mock.Mock.Test(t)
//
//	t.Cleanup(func() { mock.AssertExpectations(t) })
//	return mock
//}
