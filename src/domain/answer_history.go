package domain

type AnswerHistory struct {
	AnswerId  int    `json:"answer_id"`
	IsCorrect bool   `json:"is_correct"`
	Contents  string `json:"contents"`
}

type AnswerHistoryUpdateRequest struct {
	AnswerId int    `json:"answer_id"`
	Contents string `json:"contents"`
}
