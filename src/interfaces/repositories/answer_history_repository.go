package repositories

import (
	"errors"
	"fmt"
	entity "shortcut_master_api/src/domain"
	"sort"
	"strconv"

	"gorm.io/gorm"
)

type AnswerHistoryRepository struct {
	SqlHandler SqlHandler
}

// TODO: refactor this function
func (db *AnswerHistoryRepository) SelectAnswerHistories(uid int, quizType string) ([]entity.AnswerHistory, error) {
	answerHistories := []entity.AnswerHistory{}
	params := generateAnswerHistoryParams(uid)
	relationParams := generateAnswerHistoryRelationParams(quizType)
	res := db.SqlHandler.FindAllByParamsWithRelation(&answerHistories, params, relationParams)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []entity.AnswerHistory{}, fmt.Errorf("Record not found")
		}
		return []entity.AnswerHistory{}, fmt.Errorf("Failed to get question")
	}
	sort.SliceStable(answerHistories, func(i, j int) bool {
		return answerHistories[i].AnswerId < answerHistories[j].AnswerId
	})
	return answerHistories, nil
}

func (db *AnswerHistoryRepository) InsertAnswerHistories(answerHistories []entity.AnswerHistory) ([]entity.AnswerHistory, error) {
	res := db.SqlHandler.Create(&answerHistories)
	if err := res.Error; err != nil {
		return []entity.AnswerHistory{}, fmt.Errorf("Failed to create answer")
	}
	return answerHistories, nil
}

func generateAnswerHistoryParams(uid int) []map[string]interface{} {
	var pageParams []map[string]interface{}
	var conditionParams []map[string]interface{}
	var params []map[string]interface{}

	pageParams = append(pageParams, map[string]interface{}{
		"order": "id desc",
		"limit": entity.AnswerHistoryLatestUnitSize,
	})

	conditionParams = append(conditionParams, map[string]interface{}{
		"column":    "user_id",
		"condition": strconv.Itoa(uid),
	})

	params = append(params, map[string]interface{}{
		"conditions": conditionParams,
		"page":       pageParams,
	})

	return params
}

func generateAnswerHistoryRelationParams(quizType string) []map[string]interface{} {
	var params []map[string]interface{}
	var relationParams []map[string]interface{}
	var conditionParams []map[string]interface{}

	conditionParams = append(conditionParams, map[string]interface{}{
		"column":    "quiz_type",
		"condition": quizType,
	})

	relationParams = append(relationParams, map[string]interface{}{
		"table":        "answers",
		"where":        conditionParams,
		"relation_key": "id",
	})

	params = append(params, map[string]interface{}{
		"self": map[string]interface{}{
			"table":        "answer_histories",
			"relation_key": "answer_id",
		},
		"relation": relationParams,
	})

	return params
}
