package service

import (
	"context"
	"github.com/andrewshostak/awesome-service/model"
	"github.com/andrewshostak/awesome-service/repository"
)

type QuestionService interface {
	Create(ctx context.Context, question model.QuestionCreate) (*model.Question, error)
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context) ([]model.Question, error)
	One(ctx context.Context, id uint) (*model.Question, error)
	Update(ctx context.Context, id uint, question model.QuestionUpdate) (*model.Question, error)
}

type questionService struct {
	qr repository.QuestionRepo
}

func NewQuestionService(qr repository.QuestionRepo) QuestionService {
	return &questionService{qr: qr}
}

func (s *questionService) Create(ctx context.Context, question model.QuestionCreate) (*model.Question, error) {
	mapped := question.MapToQuestion()
	return s.qr.Create(ctx, mapped)
}

func (s *questionService) Delete(ctx context.Context, id uint) error {
	return s.qr.Delete(ctx, id)
}

func (s *questionService) List(ctx context.Context) ([]model.Question, error) {
	return s.qr.List(ctx)
}

func (s *questionService) One(ctx context.Context, id uint) (*model.Question, error) {
	return s.qr.One(ctx, id)
}

func (s *questionService) Update(ctx context.Context, id uint, question model.QuestionUpdate) (*model.Question, error) {
	_, err := s.qr.One(ctx, id)
	if err != nil {
		return nil, err
	}

	mapped := question.MapToQuestion(id)
	return s.qr.Update(ctx, mapped)
}
