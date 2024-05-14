package repo

import (
	"auth-service/internal/repo/model"
	"context"
	"github.com/stretchr/testify/assert"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
	"testing"
	"time"
)

func TestUser_GetById(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlxmock.NewRows(
		[]string{"id", "username", "password",
			"first_name", "last_name", "email",
			"phone", "is_superuser", "is_staff", "is_active",
			"created_at", "updated_at"}).AddRow(
		"1", "Test", "123456", "Ulan", "abdraman",
		"ulan@gmail.com", "+7787764654", true, false, false, time.Now(), time.Now())
	query := `SELECT (.*) FROM "user" WHERE id = \\?`
	mock.ExpectQuery(query).WillReturnRows(rows)
	a := NewUser(db)

	anUser, err := a.GetById(context.TODO(), "1")

	assert.NoError(t, err)
	assert.NotEmpty(t, anUser)

}

func TestUser_GetAll(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlxmock.NewRows([]string{"id", "username", "password",
		"first_name", "last_name", "email",
		"phone", "is_superuser", "is_staff", "is_active",
		"created_at", "updated_at"}).
		AddRow(
			"1", "Test", "123456", "Ulan", "abdraman",
			"ulan@gmail.com", "+7787764654", true, false, false, time.Now(), time.Now()).
		AddRow(
			"2", "Test2", "123456", "Ulan2", "abdraman2",
			"ulan@gmail.com", "+7787764654", true, false, false, time.Now(), time.Now())

	query := `SELECT (.*) FROM "user"`
	mock.ExpectQuery(query).WillReturnRows(rows)
	a := NewUser(db)
	anUsers, err := a.GetAll(context.TODO())

	assert.NoError(t, err)
	assert.Len(t, anUsers, 2)
}

func TestUser_Create(t *testing.T) {
	user := model.User{
		Username:    "Test",
		Password:    "123456",
		FirstName:   "Test",
		LastName:    "Test",
		Email:       "test@test.com",
		Phone:       "+787686",
		IsSuperuser: true,
		IsStaff:     true,
		IsActive:    true,
	}

	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := `INSERT INTO "user" 
    \(username, password, first_name, 
    last_name, email, phone, is_superuser, 
    is_staff, is_active\) 
    VALUES \(\?, \?, \?, \?, \?, \?, \?, \?, \?\)`

	mock.ExpectExec(query).
		WithArgs(user.Username, user.Password,
			user.FirstName, user.LastName, user.Email, user.Phone,
			user.IsSuperuser, user.IsStaff, user.IsActive).
		WillReturnResult(sqlxmock.NewResult(1, 1))

	a := NewUser(db)
	err = a.Create(context.TODO(), user)
	assert.NoError(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUser_Update(t *testing.T) {}

func TestUser_Delete(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := `DELETE FROM "user" WHERE id = \\?`
	mock.ExpectExec(query).WillReturnResult(sqlxmock.NewResult(1, 1))

	a := NewUser(db)
	err = a.Delete(context.TODO(), "1")
	assert.NoError(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
