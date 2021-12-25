package repository

import (
	"context"
	"github.com/andrewshostak/awesome-service/model"
	"github.com/stretchr/testify/mock"
)

var _ QuestionRepo = &QuestionRepoMock{}

type QuestionRepoMock struct {
	mock.Mock
}

func (m *QuestionRepoMock) Create(ctx context.Context, question model.Question) (*model.Question, error) {
	panic("implement me")
}

func (m *QuestionRepoMock) Delete(ctx context.Context, id uint) error {
	panic("implement me")
}

func (m *QuestionRepoMock) List(ctx context.Context) ([]model.Question, error) {
	panic("implement me")
}

func (m *QuestionRepoMock) One(ctx context.Context, id uint) (*model.Question, error) {
	args := m.Called(ctx, id)

	arg := args.Get(0)
	var question *model.Question
	if arg != nil {
		question = arg.(*model.Question)
	}

	return question, args.Error(1)
}

func (m *QuestionRepoMock) Update(ctx context.Context, question model.Question) (*model.Question, error) {
	args := m.Called(ctx, question)

	arg := args.Get(0)
	var q *model.Question
	if arg != nil {
		q = arg.(*model.Question)
	}

	return q, args.Error(1)
}
