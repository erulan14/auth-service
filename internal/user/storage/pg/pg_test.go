package pg

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

	sqlxmock "github.com/zhashkevych/go-sqlxmock"

	"auth-service/internal/user/model"
)

func TestStorage_Create(t *testing.T) {
	user := model.User{
		Username: "test",
		Password: "test",
	}

	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	a := NewStorage(db)
	query := `INSERT INTO "user" \(username, password\) VALUES \(\$1, \$2\) RETURNING id`

	mock.ExpectQuery(query).WithArgs(user.Username, user.Password).
		WillReturnRows(sqlxmock.NewRows([]string{"id"}).AddRow("1"))

	id, err := a.Create(context.TODO(), user)

	assert.NoError(t, err)
	assert.NotEmpty(t, id)

	if id != "1" {
		t.Errorf("expected id: 1, got: %s", id)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestStorage_Update(t *testing.T) {

}

func TestStorage_GetByID(t *testing.T) {
	userID := "123456"

	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectQuery(`SELECT \* FROM "user" WHERE is_deleted=FALSE AND id=\$1`).
		WithArgs(userID).
		WillReturnRows(sqlxmock.NewRows([]string{"id", "username", "password", "first_name",
			"last_name", "email", "phone", "is_superuser", "is_staff",
			"is_active", "created_at", "updated_at", "is_deleted"}).
			AddRow(userID, "test", "password", "test", "example",
				"test@example.com", "+77087784312", true, true, true, time.Now(), time.Now(), false))

	a := NewStorage(db)
	result, err := a.GetByID(context.TODO(), userID)

	assert.NoError(t, err)
	assert.NotNil(t, result)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestStorage_GetAll(t *testing.T) {
	users := []model.User{
		{
			ID:        "1",
			Username:  "test",
			Password:  "test",
			Email:     "email@example.com",
			Phone:     "+7707 777 77 77",
			IsSuper:   true,
			IsStaff:   true,
			IsActive:  true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			IsDeleted: false,
		},
		{
			ID:        "2",
			Username:  "test",
			Password:  "test",
			Email:     "email@example.com",
			Phone:     "+7707 777 77 77",
			IsSuper:   true,
			IsStaff:   true,
			IsActive:  true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			IsDeleted: false,
		},
	}

	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectQuery(`SELECT \* FROM "user" WHERE is_deleted=false`).
		WillReturnRows(sqlxmock.NewRows([]string{"id", "username", "password", "first_name",
			"last_name", "email", "phone", "is_superuser", "is_staff",
			"is_active", "created_at", "updated_at", "is_deleted"}).
			AddRow(users[0].ID, users[0].Username, users[0].Password, users[0].FirstName,
				users[0].LastName, users[0].Email, users[0].Phone, users[0].IsSuper, users[0].IsStaff,
				users[0].IsActive, users[0].CreatedAt, users[0].UpdatedAt, users[0].IsDeleted).
			AddRow(users[1].ID, users[1].Username, users[1].Password, users[1].FirstName,
				users[1].LastName, users[1].Email, users[1].Phone, users[1].IsSuper, users[1].IsStaff,
				users[1].IsActive, users[1].CreatedAt, users[1].UpdatedAt, users[1].IsDeleted))

	a := NewStorage(db)
	result, err := a.GetAll(context.TODO())

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 2)
	assert.Equal(t, users, result)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestStorage_Delete(t *testing.T) {
	userID := "1"

	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectExec(`UPDATE "user" SET is_deleted = \$2 WHERE id = \$1`).
		WithArgs(userID, true).
		WillReturnResult(sqlxmock.NewResult(1, 1))

	a := NewStorage(db)
	err = a.Delete(context.TODO(), userID)

	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
