package repositories

import (
	"github.com/ViniciusLugli/mindshelf/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(task *models.Task) error {
	return r.db.Create(task).Error
}

func (r *TaskRepository) Update(task *models.Task) error {
	return r.db.Model(task).Updates(task).Error
}

func (r *TaskRepository) Delete(task *models.Task) error {
	return r.db.Delete(task).Error
}

func (r *TaskRepository) GetByID(id uuid.UUID) (models.Task, error) {
	var task models.Task
	err := r.db.First(&task, id).Error
	return task, err
}

func (r *TaskRepository) GetByTitle(title string) (models.Task, error) {
	var task models.Task
	err := r.db.Where("title = ?", title).First(&task).Error
	return task, err
}

func (r *TaskRepository) GetAllByTitle(title string, limit, offset int) ([]models.Task, int64, error) {
	var tasks []models.Task
	var count int64

	err := r.db.Model(&models.Task{}).Where("title LIKE ?", "%"+title+"%").Count(&count).Error
	if err != nil {
		return tasks, count, err
	}

	err = r.db.Where("title LIKE ?", "%"+title+"%").Limit(limit).Offset(offset).Find(&tasks).Error
	if err != nil {
		return tasks, count, err
	}

	return tasks, count, nil
}

func (r *TaskRepository) GetAll(limit, offset int) ([]models.Task, int64, error) {
	var tasks []models.Task
	var count int64

	err := r.db.Model(&models.Task{}).Count(&count).Error
	if err != nil {
		return tasks, count, err
	}

	err = r.db.Limit(limit).Offset(offset).Find(&tasks).Error
	if err != nil {
		return tasks, count, err
	}

	return tasks, count, nil
}
