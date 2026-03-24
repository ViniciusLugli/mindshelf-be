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

func (r *GroupRepository) GetByID(id uuid.UUID, userID uuid.UUID) (models.Group, error) {
	var group models.Group
	err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&group).Error
	return group, err
}

func (r *GroupRepository) GetAllByName(name string, limit, offset int, userID uuid.UUID) ([]models.Group, int64, error) {
	var groups []models.Group
	var count int64

	base := r.db.Where("name LIKE ? and user_id = ?", "%"+name+"%", userID)

	if err := base.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := base.Limit(limit).Offset(offset).Find(&groups).Error; err != nil {
		return nil, 0, err
	}

	return groups, count, nil
}

func (r *GroupRepository) GetAll(limit, offset int, userID uuid.UUID) ([]models.Group, int64, error) {
	var groups []models.Group
	var count int64

	base := r.db.Model(groups).Where("user_id = ?", userID)

	if err := base.Count(&count).Error; err != nil {
		return groups, count, err
	}

	if err := base.Limit(limit).Offset(offset).Find(&groups).Error; err != nil {
		return groups, count, err
	}

	return groups, count, nil
}
