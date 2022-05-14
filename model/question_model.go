package model

type Question struct {
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	IsAnswered bool   `json:"is_answered"`
}

type QuestionCreate struct {
	Title      string `json:"title" binding:"required" validators:"min=3,max=10"`
	IsAnswered bool   `json:"is_answered"`
}

func (q QuestionCreate) MapToQuestion() Question {
	return Question{
		Title:      q.Title,
		IsAnswered: q.IsAnswered,
	}
}

type QuestionUpdate struct {
	Title      string `json:"title" binding:"required" validators:"min=3,max=10"`
	IsAnswered bool   `json:"is_answered" binding:"required"`
}

func (q QuestionUpdate) MapToQuestion(id uint) Question {
	return Question{
		ID:         id,
		Title:      q.Title,
		IsAnswered: q.IsAnswered,
	}
}
