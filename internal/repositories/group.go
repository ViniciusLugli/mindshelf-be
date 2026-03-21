package repositories

import (
	"github.com/ViniciusLugli/mindshelf/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GroupRepository struct {
	db *gorm.DB
}

func NewGroupRepository(db *gorm.DB) *GroupRepository {
	return &GroupRepository{db: db}
}

func (r *GroupRepository) Create(group *models.Group) error {
	return r.db.Create(group).Error
}

func (r *GroupRepository) Update(group *models.Group) error {
	return r.db.Model(group).Updates(group).Error
}

func (r *GroupRepository) Delete(group *models.Group) error {
	return r.db.Delete(group).Error
}

func (r *GroupRepository) GetByID(id uuid.UUID) (models.Group, error) {
	var group models.Group
	err := r.db.First(&group, id).Error
	return group, err
}

func (r *GroupRepository) GetAllByName(name string, limit, offset int) ([]models.Group, int64, error) {
	var groups []models.Group
	var count int64

	err := r.db.Where("name LIKE ?", "%"+name+"%").Count(&count).Error

	err = r.db.Where("name LIKE ?", "%"+name+"%").Limit(limit).Offset(offset).Find(&groups).Error
	return groups, count, err
}

func (r *GroupRepository) GetAll(limit, offset int) ([]models.Group, int64, error) {
	var groups []models.Group
	var count int64

	err := r.db.Model(groups).Count(&count).Error
	if err != nil {
		return groups, count, err
	}

	err = r.db.Limit(limit).Offset(offset).Find(&groups).Error
	if err != nil {
		return groups, count, err
	}

	return groups, count, nil
}
