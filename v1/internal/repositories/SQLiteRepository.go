package repositories

import (
	"errors"
	"log"
	"strconv"

	"github.com/maneulf/guarapo_lab_test/internal/handlers/models"
	"github.com/maneulf/guarapo_lab_test/internal/handlers/models/req"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLiteRepository struct {
	db *gorm.DB
}

func NewSQLiteRepository(database string) *SQLiteRepository {
	r := &SQLiteRepository{}
	r.Connect(database)
	return r
}

func (SQLr *SQLiteRepository) Connect(database string) {
	var err error
	SQLr.db, err = gorm.Open(sqlite.Open(database), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database, Err: %s", err)
	}
	SQLr.db.Migrator().DropTable(&models.DbTask{})
	SQLr.db.AutoMigrate(&models.DbTask{})

}

func (SQLr *SQLiteRepository) GetTasks(token string) ([]req.Task, error) {
	var dbTasks []models.DbTask
	result := SQLr.db.Where("user_token = ?", token).Find(&dbTasks)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("no data")
	}

	var tasks []req.Task

	for _, v := range dbTasks {
		i, error := strconv.Atoi(v.ID)

		if error != nil {
			return tasks, errors.New("unable to convert id to int")
		}
		tasks = append(tasks, req.Task{
			ID:        i,
			Title:     v.Title,
			Completed: v.Completed,
			Owner:     v.Owner,
		})

	}

	return tasks, nil
}

func (SQLr *SQLiteRepository) GetTask(id int, token string) (req.Task, error) {
	var dbTask models.DbTask
	result := SQLr.db.Where("id = ? and user_token = ?", id, token).First(&dbTask)
	var task req.Task

	if result.Error != nil {
		return task, result.Error
	}

	if result.RowsAffected == 0 {
		return task, errors.New("no data")
	}

	i, err := strconv.Atoi(dbTask.ID)

	if err != nil {
		return task, errors.New("unable to convert task id")

	}

	task = req.Task{
		ID:        i,
		Title:     dbTask.Title,
		Completed: dbTask.Completed,
		Owner:     dbTask.Owner,
	}
	return task, nil
}

func (SQLr *SQLiteRepository) Save(task req.Task, token string) error {

	str := strconv.Itoa(task.ID)
	dbTask := models.DbTask{
		ID:        str,
		Title:     task.Title,
		Completed: task.Completed,
		Owner:     task.Owner,
		UserToken: token,
	}

	result := SQLr.db.Create(&dbTask)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (SQLr *SQLiteRepository) Update(task req.Task, id int, token string) error {
	result := SQLr.db.Model(&models.DbTask{}).Where("id = ? and user_token = ?", id, token).Updates(
		map[string]interface{}{
			"id":         task.ID,
			"title":      task.Title,
			"completed":  task.Completed,
			"owner":      task.Owner,
			"user_token": token,
		},
	)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no data updated")
	}

	return nil
}

func (SQLr *SQLiteRepository) Delete(id int, token string) error {
	//result := SQLr.db.Delete(&models.DbTask{}, id)
	result := SQLr.db.Where("id = ? and user_token = ?", id, token).Delete(&models.DbTask{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
