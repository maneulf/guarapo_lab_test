package repositories

import (
	"log"

	"github.com/maneulf/guarapo_lab_test/internal/handlers/models/req"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLiteRepository struct {
	db *gorm.DB
}

func NewSQLiteRepository() *SQLiteRepository {
	r := &SQLiteRepository{}
	r.Connect()
	return r
}

func (SQLr *SQLiteRepository) Connect() {
	var err error
	SQLr.db, err = gorm.Open(sqlite.Open("Tasks.db"), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database, Err: %s", err)
	}
	SQLr.db.AutoMigrate(&req.Task{})

}

func (SQLr *SQLiteRepository) GetTasks(token string) ([]req.Task, error) {
	SQLr.db.Where("id = ?")
	return []req.Task{}, nil

}

func (SQLr *SQLiteRepository) GetTask(id int, token string) (req.Task, error) {
	var task req.Task
	result := SQLr.db.Where("id = ?", id).First(&task)
	if result.Error != nil {
		return task, result.Error
	}
	return task, nil
}

func (SQLr *SQLiteRepository) Save(task req.Task, token string) error {
	result := SQLr.db.Create(&task)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (SQLr *SQLiteRepository) Update(task req.Task, id int, token string) error {
	return nil
}

func (SQLr *SQLiteRepository) Delete(id int, token string) error {
	return nil
}
