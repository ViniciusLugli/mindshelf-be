package repositories

import (
	"github.com/ViniciusLugli/mindshelf/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (r *TaskRepository) CreateTx(tx *gorm.DB, task *models.Task) error {
	return tx.Create(task).Error
}

func (r *TaskRepository) Update(task *models.Task) error {
	return r.db.
		Model(&models.Task{}).
		Where("id = ?", task.ID).
		Omit(clause.Associations).
		Updates(map[string]any{
			"title": task.Title,
			"notes": task.Notes,
		}).Error
}

func (r *TaskRepository) Delete(task *models.Task) error {
	return r.db.Delete(task).Error
}

func (r *TaskRepository) userGroupsSubquery(db *gorm.DB, userID uuid.UUID) *gorm.DB {
	return db.Model(&models.Group{}).Select("id").Where("user_id = ?", userID)
}

func (r *TaskRepository) GetByID(id uuid.UUID, userID uuid.UUID) (models.Task, error) {
	return r.GetByIDTx(r.db, id, userID)

}

func (r *TaskRepository) GetByIDTx(tx *gorm.DB, id uuid.UUID, userID uuid.UUID) (models.Task, error) {
	var task models.Task
	err := tx.
		Preload("Group").
		Where("tasks.id = ? AND tasks.group_id IN (?)", id, r.userGroupsSubquery(tx, userID)).
		First(&task).Error
	return task, err
}

func (r *TaskRepository) GetByTitle(title string, userID uuid.UUID) (models.Task, error) {
	var task models.Task
	err := r.db.
		Preload("Group").
		Where("tasks.title = ? AND tasks.group_id IN (?)", title, r.userGroupsSubquery(r.db, userID)).
		First(&task).Error

	return task, err
}

func (r *TaskRepository) GetAllByTitle(title string, limit, offset int, userID uuid.UUID) ([]models.Task, int64, error) {
	var tasks []models.Task
	var count int64

	base := r.db.Model(&models.Task{}).
		Where("tasks.title LIKE ? AND tasks.group_id IN (?)", "%"+title+"%", r.userGroupsSubquery(r.db, userID))

	if err := base.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := base.Preload("Group").Limit(limit).Offset(offset).Find(&tasks).Error; err != nil {
		return nil, 0, err
	}

	return tasks, count, nil
}

func (r *TaskRepository) GetAll(limit, offset int, userID uuid.UUID) ([]models.Task, int64, error) {
	var tasks []models.Task
	var count int64

	base := r.db.Model(&models.Task{}).Where("tasks.group_id IN (?)", r.userGroupsSubquery(r.db, userID))

	if err := base.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := base.Preload("Group").Limit(limit).Offset(offset).Find(&tasks).Error; err != nil {
		return nil, 0, err
	}

	return tasks, count, nil
}

func (r *TaskRepository) GetAllFiltered(title string, groupID *uuid.UUID, limit, offset int, userID uuid.UUID) ([]models.Task, int64, error) {
	var tasks []models.Task
	var count int64

	base := r.db.Model(&models.Task{}).Where("tasks.group_id IN (?)", r.userGroupsSubquery(r.db, userID))

	if title != "" {
		base = base.Where("tasks.title LIKE ?", "%"+title+"%")
	}

	if groupID != nil {
		base = base.Where("tasks.group_id = ?", *groupID)
	}

	if err := base.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := base.Preload("Group").Limit(limit).Offset(offset).Find(&tasks).Error; err != nil {
		return nil, 0, err
	}

	return tasks, count, nil
}

func (r *TaskRepository) GetAllByGroupID(groupID uuid.UUID, limit, offset int, userID uuid.UUID) ([]models.Task, int64, error) {
	var tasks []models.Task
	var count int64

	base := r.db.Model(&models.Task{}).
		Where("tasks.group_id = ? AND tasks.group_id IN (?)", groupID, r.userGroupsSubquery(r.db, userID))

	if err := base.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := base.Preload("Group").Limit(limit).Offset(offset).Find(&tasks).Error; err != nil {
		return nil, 0, err
	}

	return tasks, count, nil
}
