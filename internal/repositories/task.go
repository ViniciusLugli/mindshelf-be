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

func (r *TaskRepository) GetByID(id uuid.UUID, userID uuid.UUID) (models.Task, error) {
	var task models.Task
	err := r.db.
		Joins("Group").
		Where(`tasks.id = ? AND "Group"."user_id" = ?`, id, userID).
		First(&task).Error
	return task, err
}

func (r *TaskRepository) GetByTitle(title string, userID uuid.UUID) (models.Task, error) {
	var task models.Task
	err := r.db.
		Joins("Group").
		Where(`tasks.title = ? AND "Group"."user_id" = ?`, title, userID).
		First(&task).Error

	return task, err
}

func (r *TaskRepository) GetAllByTitle(title string, limit, offset int, userID uuid.UUID) ([]models.Task, int64, error) {
	var tasks []models.Task
	var count int64

	base := r.db.Model(&models.Task{}).
		Joins("Group").
		Where(`tasks.title LIKE ? AND "Group"."user_id" = ?`, "%"+title+"%", userID)

	if err := base.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := base.Limit(limit).Offset(offset).Find(&tasks).Error; err != nil {
		return nil, 0, err
	}

	return tasks, count, nil
}

func (r *TaskRepository) GetAll(limit, offset int, userID uuid.UUID) ([]models.Task, int64, error) {
	var tasks []models.Task
	var count int64

	base := r.db.Model(&models.Task{}).Joins("Group").Where(`"Group"."user_id" = ?`, userID)

	if err := base.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := base.Limit(limit).Offset(offset).Find(&tasks).Error; err != nil {
		return nil, 0, err
	}

	return tasks, count, nil
}

func (r *TaskRepository) GetAllByGroupID(groupID uuid.UUID, limit, offset int, userID uuid.UUID) ([]models.Task, int64, error) {
	var tasks []models.Task
	var count int64

	base := r.db.Model(&models.Task{}).
		Joins("Group").
		Where(`tasks.group_id = ? AND "Group"."user_id" = ?`, groupID, userID)

	if err := base.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := base.Limit(limit).Offset(offset).Find(&tasks).Error; err != nil {
		return nil, 0, err
	}

	return tasks, count, nil
}
