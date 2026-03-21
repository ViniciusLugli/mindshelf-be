package services

import (
	"errors"
	"math"

	"github.com/ViniciusLugli/mindshelf/internal/dtos/requests"
	"github.com/ViniciusLugli/mindshelf/internal/dtos/responses"
	"github.com/ViniciusLugli/mindshelf/internal/repositories"
	"github.com/google/uuid"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(dto requests.CreateUserRequest) error {
	user := dto.ToModel()
	return s.repo.Create(&user)
}

func (s *UserService) Update(dto requests.UpdateUserRequest) error {
	user, err := s.repo.GetByID(dto.ID)
	if err != nil {
		return err
	}

	user.Name = dto.Name
	user.Email = dto.Email
	user.Password = dto.Password

	return s.repo.Update(&user)
}

func (s *UserService) Delete(dto requests.DeleteUserRequest) error {
	user, err := s.repo.GetByID(dto.ID)
	if err != nil {
		return err
	}

	return s.repo.Delete(&user)
}

func (s *UserService) GetUser(dto requests.GetUser) (responses.UserResponse, error) {
	if dto.ID != uuid.Nil {
		user, err := s.repo.GetByID(dto.ID)
		if err != nil {
			return responses.UserResponse{}, err
		}

		return responses.NewUserResponse(user), nil
	}

	if dto.Email != "" {
		user, err := s.repo.GetByEmail(dto.Email)
		if err != nil {
			return responses.UserResponse{}, err
		}

		return responses.NewUserResponse(user), err
	}

	return responses.UserResponse{}, errors.New("no valid params was passed")
}

func (s *UserService) GetAllUsers(dto requests.GetAllUsers) (responses.PaginatedUserResponse, error) {
	offset := (dto.Page - 1) * dto.Limit

	users, count, err := s.repo.GetAll(dto.Limit, offset)
	if err != nil {
		return responses.PaginatedUserResponse{}, err
	}

	total_pages := math.Ceil(float64(count) / float64(dto.Limit))
	return responses.NewPaginatedUserResponse(users, count, dto.Page, dto.Limit, int(total_pages)), nil
}
