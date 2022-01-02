package repository

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/andrewshostak/awesome-service/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"regexp"
	"testing"
)

func TestQuestionRepo_List_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sql mock")
	}
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open db connection")
	}

	repo := NewQuestionRepo(gormDB)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "questions"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "is_answered"}).AddRow(1, "test", true))

	result, err := repo.List(context.Background())

	assert.Nil(t, err)
	assert.Equal(t, result, []model.Question{{1, "test", true}})
}

func TestQuestionRepo_List_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sql mock")
	}
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open db connection")
	}

	repo := NewQuestionRepo(gormDB)
	err = errors.New("some database error")
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "questions"`)).
		WillReturnError(err)

	result, repoErr := repo.List(context.Background())

	assert.Nil(t, result)
	assert.Equal(t, err, repoErr)
}
