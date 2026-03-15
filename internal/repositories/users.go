package repositories

import (
	"github.com/ViniciusLugli/mindshelf/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.Users) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) Update(user *models.Users) error {
	return r.db.Model(user).Updates(user).Error
}

func (r *UserRepository) Delete(user *models.Users) error {
	return r.db.Delete(user).Error
}

func (r *UserRepository) GetByID(id uuid.UUID) (models.Users, error) {
	var user models.Users
	err := r.db.First(&user, id).Error
	return user, err
}

func (r *UserRepository) GetByEmail(email string) (models.Users, error) {
	var user models.Users
	err := r.db.First(&user, email).Error
	return user, err
}

func (r *UserRepository) GetAll(limit, offset int) ([]models.Users, int64, error) {
	var users []models.Users
	var count int64

	err := r.db.Model(&models.Users{}).Count(&count).Error
	if err != nil {
		return users, count, err
	}

	err = r.db.Limit(limit).Offset(offset).Find(&users).Error
	if err != nil {
		return users, count, err
	}

	return users, count, nil
}
