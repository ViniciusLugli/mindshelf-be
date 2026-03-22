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

func (s *UserService) Update(dto requests.UpdateUserRequest, id uuid.UUID) error {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	user.Name = dto.Name
	user.Email = dto.Email
	user.Password = dto.Password

	return s.repo.Update(&user)
}

func (s *UserService) Delete(id uuid.UUID) error {
	user, err := s.repo.GetByID(id)
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

func (s *UserService) GetAllUsers(dto requests.GetAllUsers) (responses.PaginatedResponse[responses.UserResponse], error) {
	offset := (dto.Page - 1) * dto.Limit

	users, count, err := s.repo.GetAll(dto.Limit, offset)
	if err != nil {
		return responses.PaginatedResponse[responses.UserResponse]{}, err
	}

	totalPages := math.Ceil(float64(count) / float64(dto.Limit))
	return responses.NewPaginatedResponse(users, responses.NewUserResponse, count, dto.Page, dto.Limit, int(totalPages)), nil
}

func (s *UserService) GetAllUsersByName(dto requests.GetAllUsersByName) (responses.PaginatedResponse[responses.UserResponse], error) {
	offset := (dto.Page - 1) * dto.Limit

	users, count, err := s.repo.GetAllByName(dto.Name, dto.Limit, offset)
	if err != nil {
		return responses.PaginatedResponse[responses.UserResponse]{}, err
	}

	totalPages := math.Ceil(float64(count) / float64(dto.Limit))
	return responses.NewPaginatedResponse(users, responses.NewUserResponse, count, dto.Page, dto.Limit, int(totalPages)), nil
}
