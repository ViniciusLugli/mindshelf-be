package services

import (
	"math"

	"github.com/ViniciusLugli/mindshelf/internal/dtos/requests"
	"github.com/ViniciusLugli/mindshelf/internal/dtos/responses"
	"github.com/ViniciusLugli/mindshelf/internal/models"
	"github.com/ViniciusLugli/mindshelf/internal/repositories"
	"github.com/google/uuid"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Update(dto requests.UpdateUserRequest, id uuid.UUID) error {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if dto.Name != "" {
		user.Name = dto.Name
	}

	if dto.Email != "" {
		user.Email = dto.Email
	}

	if dto.Password != "" {
		hashedPassword, err := HashPassword(dto.Password)
		if err != nil {
			return err
		}

		user.Password = hashedPassword
	}

	return s.repo.Update(&user)
}

func (s *UserService) Delete(id uuid.UUID) error {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(&user)
}

func (s *UserService) GetUserByID(id uuid.UUID) (responses.UserResponse, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return responses.UserResponse{}, err
	}

	return responses.NewUserResponse(user), nil
}

func (s *UserService) GetAllUsers(dto requests.GetAllUsers) (responses.PaginatedResponse[responses.UserResponse], error) {
	offset := (dto.Page - 1) * dto.Limit

	var (
		users []models.User
		count int64
		err   error
	)

	if dto.Name != "" {
		users, count, err = s.repo.GetAllByName(dto.Name, dto.Limit, offset)
	} else {
		users, count, err = s.repo.GetAll(dto.Limit, offset)
	}

	if err != nil {
		return responses.PaginatedResponse[responses.UserResponse]{}, err
	}

	totalPages := math.Ceil(float64(count) / float64(dto.Limit))
	return responses.NewPaginatedResponse(users, responses.NewUserResponse, count, dto.Page, dto.Limit, int(totalPages)), nil
}

func (s *UserService) SendFriendRequest(userID uuid.UUID, dto requests.FriendRequest) error {
	return s.repo.SendFriendRequest(userID, dto.FriendID)
}

func (s *UserService) AcceptFriendRequest(userID uuid.UUID, dto requests.FriendRequest) error {
	return s.repo.AcceptFriendRequest(userID, dto.FriendID)
}

func (s *UserService) RejectFriendRequest(userID uuid.UUID, dto requests.FriendRequest) error {
	return s.repo.RejectFriendRequest(userID, dto.FriendID)
}

func (s *UserService) RemoveFriend(userID uuid.UUID, dto requests.FriendRequest) error {
	return s.repo.RemoveFriend(userID, dto.FriendID)
}

func (s *UserService) GetFriends(userID uuid.UUID) ([]responses.UserResponse, error) {
	friends, err := s.repo.GetFriends(userID)
	if err != nil {
		return nil, err
	}

	friendsDto := make([]responses.UserResponse, len(friends))
	for i, friend := range friends {
		friendsDto[i] = responses.NewUserResponse(friend)
	}

	return friendsDto, nil
}

func (s *UserService) GetPendingFriendRequests(userID uuid.UUID) ([]responses.ReceivedFriendRequestResponse, error) {
	friendRequests, err := s.repo.GetPendingFriendRequests(userID)
	if err != nil {
		return nil, err
	}

	friendRequestsDto := make([]responses.ReceivedFriendRequestResponse, len(friendRequests))
	for i, friendship := range friendRequests {
		friendRequestsDto[i] = responses.NewReceivedFriendRequestResponse(friendship)
	}

	return friendRequestsDto, nil
}
