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

	a := NewUser(db)

	rows := sqlxmock.NewRows(
		[]string{"id", "username", "password",
			"first_name", "last_name", "email",
			"phone", "is_superuser", "is_staff", "is_active",
			"created_at", "updated_at", "is_deleted"}).AddRow(
		"12", "Test", "123456", "Ulan", "abdraman",
		"ulan@gmail.com", "+7787764654", true, false, false, time.Now(), time.Now(), false)

	query := `SELECT (.*) FROM "user" WHERE is_deleted=FALSE AND id=\\?`
	mock.ExpectQuery(query).WillReturnRows(rows)

	anUser, err := a.GetById(context.TODO(), "12")

	assert.NoError(t, err)
	assert.NotEmpty(t, anUser)
}

func TestUser_GetAll(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	a := NewUser(db)

	rows := sqlxmock.NewRows([]string{"id", "username", "password",
		"first_name", "last_name", "email",
		"phone", "is_superuser", "is_staff", "is_active",
		"created_at", "updated_at", "is_deleted"}).
		AddRow(
			"1", "Test", "123456", "Ulan", "abdraman",
			"ulan@gmail.com", "+7787764654", true, false, false, time.Now(), time.Now(), false).
		AddRow(
			"2", "Test2", "123456", "Ulan2", "abdraman2",
			"ulan@gmail.com", "+7787764654", true, false, false, time.Now(), time.Now(), false)

	query := `SELECT (.*) FROM "user" WHERE is_deleted=false`
	mock.ExpectQuery(query).WillReturnRows(rows)

	anUsers, err := a.GetAll(context.TODO())

	assert.NoError(t, err)
	assert.Len(t, anUsers, 2)
}

func TestUser_Create(t *testing.T) {
	user := model.CreateUser{
		Username: "Test",
		Password: "123456",
	}

	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	a := NewUser(db)

	query := `INSERT INTO "user" \(username, password\) VALUES \(\$1, \$2\) RETURNING id`

	mock.
		ExpectQuery(query).
		WithArgs(user.Username, user.Password).
		WillReturnRows(sqlxmock.NewRows([]string{"id"}).AddRow("Hasdkasld;akasl;"))

	id, err := a.Create(context.TODO(), user)

	assert.NoError(t, err)
	assert.NotEmpty(t, id)

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

	a := NewUser(db)

	query := `UPDATE "user" SET is_deleted = \$2 WHERE id = \$1`
	mock.ExpectExec(query).WillReturnResult(sqlxmock.NewResult(1, 1))

	err = a.Delete(context.TODO(), "1")
	assert.NoError(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
