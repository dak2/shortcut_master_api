package repositories

import (
	"errors"
	"fmt"
	entity "shortcut_master_api/src/domain"

	"gorm.io/gorm"
)

type AnswerHistoryRepository struct {
	SqlHandler SqlHandler
}

func (db *AnswerHistoryRepository) SelectAnswerHistories(quizType string) ([]entity.AnswerHistory, error) {
	answerHistories := []entity.AnswerHistory{}
	params := generateAnswerHistoryParams()
	relationParams := generateAnswerHistoryRelationParams(quizType)

	res := db.SqlHandler.FindAllByParamsWithRelation(&answerHistories, params, relationParams)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []entity.AnswerHistory{}, fmt.Errorf("Record not found")
		}
		return []entity.AnswerHistory{}, fmt.Errorf("Failed to get question")
	}
	return answerHistories, nil
}

func (db *AnswerHistoryRepository) InsertAnswerHistories(answerHistories []entity.AnswerHistory) ([]entity.AnswerHistory, error) {
	res := db.SqlHandler.Create(&answerHistories)
	if err := res.Error; err != nil {
		return []entity.AnswerHistory{}, fmt.Errorf("Failed to create answer")
	}
	return answerHistories, nil
}

func generateAnswerHistoryParams() []map[string]interface{} {
	var params []map[string]interface{}

	params = append(params, map[string]interface{}{
		"order": "id asc",
		"limit": entity.AnswerHistoryLatestUnitSize,
	})

	return params
}

func generateAnswerHistoryRelationParams(quizType string) []map[string]interface{} {
	var params []map[string]interface{}
	var relationParams []map[string]interface{}
	var conditionParams []map[string]interface{}

	params = append(params, map[string]interface{}{
		"self": map[string]interface{}{
			"table":        "answer_histories",
			"relation_key": "answer_id",
		},
		"relation": append(relationParams, map[string]interface{}{
			"table": "answers",
			"where": append(conditionParams, map[string]interface{}{
				"column":    "quiz_type",
				"condition": quizType,
			}),
			"relation_key": "id",
		}),
	})

	return params
}
