package domain

type User struct {
	ID   					int    `json:"id" gorm:"primary_key"`
	GoogleUserId	string `json:"sub"`
	Name 					string `json:"name"`
	Email 				string `json:"email"`
}
