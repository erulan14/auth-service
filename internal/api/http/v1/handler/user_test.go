package handler

import (
	"auth-service/internal/api/http/v1/handler/convertor"
	"auth-service/internal/api/http/v1/handler/model"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"auth-service/internal/api/http/v1/handler/mocks"
	"auth-service/internal/domain/entity"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/go-faker/faker/v4"
)

func TestUser_GetById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	var mockUser entity.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)

	mockUCase := new(mocks.UserUseCase)
	num := mockUser.ID

	mockUCase.On("GetById", mock.Anything, mock.AnythingOfType("string")).Return(mockUser, nil)
	w := httptest.NewRecorder()

	c, r := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodGet, "/users/:id", nil)
	c.Params = append(c.Params, gin.Param{Key: "id", Value: num})
	handler := NewUser(mockUCase)

	r.GET("/users/:id", handler.GetById)
	r.ServeHTTP(w, c.Request)

	responseData, err := ioutil.ReadAll(w.Body)
	assert.NoError(t, err)

	var respUser model.User
	_ = json.Unmarshal(responseData, &respUser)

	assert.Equal(t, convertor.ToHandlerModel(mockUser), respUser)
	assert.Equal(t, http.StatusOK, w.Code)
	mockUCase.AssertExpectations(t)
}
