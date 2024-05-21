package mocks

import (
	"auth-service/internal/domain/entity"
	"context"
	"github.com/stretchr/testify/mock"
)

type UserUseCase struct {
	mock.Mock
}

func (_m *UserUseCase) GetAll(ctx context.Context) ([]entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (_m *UserUseCase) Create(ctx context.Context, user entity.CreateUser) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (_m *UserUseCase) Update(ctx context.Context, id string, user entity.UpdateUser) error {
	//TODO implement me
	panic("implement me")
}

func (_m *UserUseCase) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

// GetById provides a mock function with given fields ctx, id
func (_m *UserUseCase) GetById(ctx context.Context, id string) (entity.User, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GET by ID")
	}

	var r0 entity.User
	var r1 error

	if rf, ok := ret.Get(0).(func(context.Context, string) (entity.User, error)); ok {
		return rf(ctx, id)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) entity.User); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
