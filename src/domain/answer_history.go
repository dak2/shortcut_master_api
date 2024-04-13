package domain

type AnswerHistory struct {
	UserId    int    `json:"user_id"`
	AnswerId  int    `json:"answer_id"`
	IsCorrect bool   `json:"is_correct"`
	Contents  string `json:"contents"`
}

type AnswerHistoryUpdateRequest struct {
	UserId   int    `json:"user_id"`
	AnswerId int    `json:"answer_id"`
	Contents string `json:"contents"`
}

var AnswerHistoryLatestUnitSize = 10
