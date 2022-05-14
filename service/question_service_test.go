package service

import (
	"context"
	"errors"
	"github.com/andrewshostak/awesome-service/errs"
	"github.com/andrewshostak/awesome-service/model"
	"github.com/andrewshostak/awesome-service/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuestionService_Update_Success(t *testing.T) {
	ctx := context.Background()
	repo := &repository.QuestionRepoMock{}
	service := NewQuestionService(repo)

	id := uint(123)
	toUpdate := model.QuestionUpdate{Title: "new title", IsAnswered: true}
	question := mapToRepositoryModel(toUpdate.MapToQuestion(id))
	expected := repository.Question{ID: id, Title: question.Title, IsAnswered: question.IsAnswered}

	repo.On("One", ctx, id).Return(&repository.Question{}, nil).Once()
	repo.On("Update", ctx, question).Return(&expected, nil).Once()

	result, err := service.Update(ctx, id, toUpdate)
	assert.Nil(t, err)
	assert.Equal(t, *mapFromRepositoryModel(&expected), *result)
}

func TestQuestionService_Update_One_Error(t *testing.T) {
	ctx := context.Background()
	repo := &repository.QuestionRepoMock{}
	service := NewQuestionService(repo)

	id := uint(123)
	toUpdate := model.QuestionUpdate{Title: "new title", IsAnswered: true}
	repoErr := errs.New(errors.New("repository error"), errs.UserError)

	repo.On("One", ctx, id).Return(nil, repoErr).Once()

	result, err := service.Update(ctx, id, toUpdate)
	assert.Equal(t, repoErr, err)
	assert.Nil(t, result)
}

func TestQuestionService_Update_Error(t *testing.T) {
	ctx := context.Background()
	repo := &repository.QuestionRepoMock{}
	service := NewQuestionService(repo)

	id := uint(123)
	toUpdate := model.QuestionUpdate{Title: "new title", IsAnswered: true}
	question := mapToRepositoryModel(toUpdate.MapToQuestion(id))
	repoErr := errs.New(errors.New("update error"), errs.UserError)

	repo.On("One", ctx, id).Return(&repository.Question{}, nil).Once()
	repo.On("Update", ctx, question).Return(nil, repoErr).Once()

	result, err := service.Update(ctx, id, toUpdate)
	assert.Equal(t, repoErr, err)
	assert.Nil(t, result)
}
