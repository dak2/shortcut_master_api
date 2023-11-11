package domain

type AnswerHistory struct {
	ID        int    `json:"id" gorm:"primary_key"`
	AnswerId  int    `json:"answer_id"`
	IsCorrect bool   `json:"is_correct"`
}
