package repository

import (
	"context"
	"github.com/stretchr/testify/mock"
)

var _ QuestionRepo = &QuestionRepoMock{}

type QuestionRepoMock struct {
	mock.Mock
}

func (m *QuestionRepoMock) Create(ctx context.Context, question Question) (*Question, error) {
	panic("implement me")
}

func (m *QuestionRepoMock) Delete(ctx context.Context, id uint) error {
	panic("implement me")
}

func (m *QuestionRepoMock) List(ctx context.Context) ([]Question, error) {
	panic("implement me")
}

func (m *QuestionRepoMock) One(ctx context.Context, id uint) (*Question, error) {
	args := m.Called(ctx, id)

	arg := args.Get(0)
	var question *Question
	if arg != nil {
		question = arg.(*Question)
	}

	return question, args.Error(1)
}

func (m *QuestionRepoMock) Update(ctx context.Context, question Question) (*Question, error) {
	args := m.Called(ctx, question)

	arg := args.Get(0)
	var q *Question
	if arg != nil {
		q = arg.(*Question)
	}

	return q, args.Error(1)
}
