package service

import (
	"auth-service/internal/domain/entity"
	"auth-service/internal/domain/service/mocks"
	"auth-service/internal/repo/model"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestUser_GetAll(t *testing.T) {
	//mockUserRepo := new(mocks.UserRepository)

	mockUser := model.User{
		ID:          "1",
		Username:    "miko",
		Password:    "lux12345",
		FirstName:   "Miko",
		LastName:    "Aio",
		Email:       "miko@gmail.com",
		Phone:       "+77007007070",
		IsActive:    false,
		IsSuperuser: false,
		IsStaff:     false,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		IsDeleted:   false,
	}

	mockListUser := make([]model.User, 1)
	mockListUser[0] = mockUser

	t.Run("success", func(t *testing.T) {
		//mockUserRepo.On("GetAll", mock.Anything).
		//	Return(mockListUser, nil).
		//	Once()
		//
		//u := NewUser(mockUserRepo)
		//all, err := u.GetAll(context.TODO())
		//if err != nil {
		//	return
		//}
		//
		//assert.NotNil(t, all)
		//assert.NoError(t, err)
		//assert.Len(t, all, len(mockListUser))
		//
		//mockUserRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {

	})

}

func TestUser_GetById(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUser := model.User{
		ID:          "1",
		Username:    "miko",
		Password:    "lux12345",
		FirstName:   "Miko",
		LastName:    "Aio",
		Email:       "miko@gmail.com",
		Phone:       "+77007007070",
		IsActive:    false,
		IsSuperuser: false,
		IsStaff:     false,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		IsDeleted:   false,
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("GetById", mock.Anything, mock.AnythingOfType("string")).
			Return(mockUser, nil).
			Once()

		u := NewUser(mockUserRepo)

		a, err := u.GetById(context.TODO(), mockUser.ID)

		assert.NoError(t, err)
		assert.NotNil(t, a)

		mockUserRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockUserRepo.On("GetById", mock.Anything, mock.AnythingOfType("string")).
			Return(model.User{}, errors.New("unexpected")).
			Once()

		u := NewUser(mockUserRepo)
		a, err := u.GetById(context.TODO(), mockUser.ID)
		assert.Error(t, err)
		assert.Equal(t, entity.User{}, a)

		mockUserRepo.AssertExpectations(t)
	})

}

func TestUser_Delete(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("Delete", mock.Anything, mock.AnythingOfType("string")).
			Return(nil).
			Once()
		u := NewUser(mockUserRepo)
		err := u.Delete(context.TODO(), "1")
		assert.NoError(t, err)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockUserRepo.On("Delete", mock.Anything, mock.AnythingOfType("string")).
			Return(errors.New("unexpected")).
			Once()
		u := NewUser(mockUserRepo)
		err := u.Delete(context.TODO(), "1")
		assert.Error(t, err)
	})
}
