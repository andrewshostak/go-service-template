package repository

type Question struct {
	ID         uint   `gorm:"primaryKey" gorm:"id"`
	Title      string `gorm:"title"`
	IsAnswered bool   `gorm:"is_answered"`
}
