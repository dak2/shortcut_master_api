package domain

type Question struct {
	ID       int    `json:"id" gorm:"primary_key"`
	QuizId   int    `json:"quiz_id"`
	QuizType string `json:"quiz_type"`
	Contents string `json:"contents"`
}
