package repository

import (
	"context"
	"errors"
	"github.com/andrewshostak/awesome-service/errs"
	"github.com/andrewshostak/awesome-service/model"
	"gorm.io/gorm"
	"strings"
)

const ErrDuplicateUnique = "duplicate key value violates unique constraint"

type QuestionRepo interface {
	Create(ctx context.Context, question model.Question) (*model.Question, error)
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context) ([]model.Question, error)
	One(ctx context.Context, id uint) (*model.Question, error)
	Update(ctx context.Context, question model.Question) (*model.Question, error)
}

type questionRepo struct {
	db *gorm.DB
}

func NewQuestionRepo(db *gorm.DB) QuestionRepo {
	return &questionRepo{db: db}
}

func (r *questionRepo) Create(ctx context.Context, question model.Question) (*model.Question, error) {
	if result := r.db.WithContext(ctx).Create(&question); result.Error != nil {
		if strings.Contains(result.Error.Error(), ErrDuplicateUnique) {
			return nil, errs.New(result.Error, errs.UserError)
		}
		return nil, result.Error
	}

	return &question, nil
}

func (r *questionRepo) Delete(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Delete(&model.Question{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errs.New(errors.New("question doesn't exist"), errs.UserError)
	}

	return nil
}

func (r *questionRepo) List(ctx context.Context) ([]model.Question, error) {
	var questions []model.Question
	if result := r.db.WithContext(ctx).Find(&questions); result.Error != nil {
		return nil, result.Error
	}

	return questions, nil
}

func (r *questionRepo) One(ctx context.Context, id uint) (*model.Question, error) {
	var question model.Question
	if result := r.db.WithContext(ctx).First(&question, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errs.New(result.Error, errs.UserError)
		}
		return nil, result.Error
	}

	return &question, nil
}

func (r *questionRepo) Update(ctx context.Context, question model.Question) (*model.Question, error) {
	if result := r.db.WithContext(ctx).Save(&question); result.Error != nil {
		if strings.Contains(result.Error.Error(), ErrDuplicateUnique) {
			return nil, errs.New(result.Error, errs.UserError)
		}
		return nil, result.Error
	}

	return &question, nil
}
