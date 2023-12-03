package domain

type Answer struct {
	ID         int    `json:"id" gorm:"primary_key"`
	QuestionId int    `json:"question_id"`
	Contents   string `json:"contents"`
}
