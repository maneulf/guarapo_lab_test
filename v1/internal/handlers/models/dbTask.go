package models

import "github.com/maneulf/guarapo_lab_test/internal/handlers/models/req"

type DbTask struct {
	req.Task
	UserToken string `json:"user_token" gorm:"primaryKey"`
}
