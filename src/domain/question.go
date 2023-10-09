package domain

type Question struct {
	ID       int    `json:"id" gorm:"primary_key"`
	QuizId   string `json:"quiz_id"`
	Contents string `json:"contents"`
}
