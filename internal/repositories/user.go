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

func (r *UserRepository) Create(user *models.User) (*models.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Update(user *models.User) error {
	return r.db.Model(user).Updates(user).Error
}

func (r *UserRepository) Delete(user *models.User) error {
	return r.db.Delete(user).Error
}

func (r *UserRepository) GetByID(id uuid.UUID) (models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return user, err
}

func (r *UserRepository) GetByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (r *UserRepository) GetAllByName(name string, limit, offset int) ([]models.User, int64, error) {
	var users []models.User
	var count int64

	err := r.db.Model(&models.User{}).Where("name LIKE ?", "%"+name+"%").Count(&count).Error
	if err != nil {
		return users, count, err
	}

	err = r.db.Where("name LIKE ?", "%"+name+"%").Limit(limit).Offset(offset).Find(&users).Error
	if err != nil {
		return users, count, err
	}

	return users, count, nil
}

func (r *UserRepository) GetAll(limit, offset int) ([]models.User, int64, error) {
	var users []models.User
	var count int64

	err := r.db.Model(&models.User{}).Count(&count).Error
	if err != nil {
		return users, count, err
	}

	err = r.db.Limit(limit).Offset(offset).Find(&users).Error
	if err != nil {
		return users, count, err
	}

	return users, count, nil
}

func (r *UserRepository) SendFriendRequest(UserID, friendID uuid.UUID) error {
	friendship := models.UserFriend{
		UserID:   UserID,
		FriendID: friendID,
		Status:   "pending",
	}

	return r.db.Create(&friendship).Error
}

func (r *UserRepository) AcceptFriendRequest(userID, friendID uuid.UUID) error {
	return r.db.Model(&models.UserFriend{}).
		Where("user_id = ? AND friend_id = ?", friendID, userID).
		Update("status", "accepted").Error
}

func (r *UserRepository) RejectFriendRequest(userID, friendID uuid.UUID) error {
	return r.db.Model(&models.UserFriend{}).
		Where("user_id = ? AND friend_id = ?", friendID, userID).
		Update("status", "rejected").Error
}

func (r *UserRepository) RemoveFriend(userID, friendID uuid.UUID) error {
	// Inverted OR because friendship may be created in whatever direction
	return r.db.Where(
		"(user_id = ? AND friend_id = ?) OR (user_id = ? AND friend_id = ?)",
		userID, friendID, friendID, userID).
		Delete(&models.UserFriend{}).Error
}

func (r *UserRepository) GetFriends(userID uuid.UUID) ([]models.User, error) {
	var friends []models.User
	err := r.db.
		Distinct("users.id", "users.name", "users.email", "users.password", "users.avatar_url", "users.created_at", "users.updated_at", "users.deleted_at").
		Joins("JOIN user_friends ON user_friends.friend_id = users.id OR user_friends.user_id = users.id").
		Where("(user_friends.user_id = ? OR user_friends.friend_id = ?) AND user_friends.status = ? AND users.id != ?",
			userID, userID, models.Accepted, userID).
		Find(&friends).Error
	return friends, err
}

func (r *UserRepository) GetPendingFriendRequests(userID uuid.UUID) ([]models.UserFriend, error) {
	var friendRequests []models.UserFriend

	err := r.db.
		Preload("User").
		Where("friend_id = ? AND status = ?", userID, models.Pending).
		Order("created_at desc").
		Find(&friendRequests).Error

	return friendRequests, err
}
