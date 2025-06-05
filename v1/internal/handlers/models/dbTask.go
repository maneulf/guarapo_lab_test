package models

type DbTask struct {
	ID        string `json:"id" gorm:"primaryKey"`
	Title     string `json:"title" binding:"required"`
	Completed bool   `json:"completed"`
	Owner     string `json:"owner"`
	UserToken string `json:"user_token" gorm:"primaryKey"`
}
