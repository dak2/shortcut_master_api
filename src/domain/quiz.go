package domain

type Quiz struct {
	ID   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	Type string `json:"type"`
}
