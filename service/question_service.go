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
	toCreate := mapToRepositoryModel(question.MapToQuestion())

	created, err := s.qr.Create(ctx, toCreate)
	if err != nil {
		return nil, err
	}

	return mapFromRepositoryModel(created), nil
}

func (s *questionService) Delete(ctx context.Context, id uint) error {
	return s.qr.Delete(ctx, id)
}

func (s *questionService) List(ctx context.Context) ([]model.Question, error) {
	questions, err := s.qr.List(ctx)
	if err != nil {
		return nil, err
	}

	return mapFromRepositoryModels(questions), nil
}

func (s *questionService) One(ctx context.Context, id uint) (*model.Question, error) {
	question, err := s.qr.One(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapFromRepositoryModel(question), nil
}

func (s *questionService) Update(ctx context.Context, id uint, question model.QuestionUpdate) (*model.Question, error) {
	_, err := s.qr.One(ctx, id)
	if err != nil {
		return nil, err
	}

	toUpdate := mapToRepositoryModel(question.MapToQuestion(id))

	updated, err := s.qr.Update(ctx, toUpdate)
	if err != nil {
		return nil, err
	}

	return mapFromRepositoryModel(updated), nil
}

func mapToRepositoryModel(question model.Question) repository.Question {
	return repository.Question{
		ID:         question.ID,
		Title:      question.Title,
		IsAnswered: question.IsAnswered,
	}
}

func mapFromRepositoryModel(question *repository.Question) *model.Question {
	return &model.Question{
		ID:         question.ID,
		Title:      question.Title,
		IsAnswered: question.IsAnswered,
	}
}

func mapFromRepositoryModels(questions []repository.Question) []model.Question {
	mapped := make([]model.Question, 0, len(questions))
	for i := range questions {
		mapped = append(mapped, *mapFromRepositoryModel(&questions[i]))
	}

	return mapped
}
