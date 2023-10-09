package database

import (
	entity "shortcut_master_api/src/domain"
)

type QuestionRepository struct {
	SqlHandler SqlHandler
}

func (db *QuestionRepository) Select() []entity.Question {
	question := []entity.Question{}
	db.SqlHandler.FindAll(&question)
	return question
}

func (db *QuestionRepository) SelectByQuiz(id string) []entity.Question {
	question := []entity.Question{}
	db.SqlHandler.FindByParams(&question, "quiz_id", id)
	return question
}
