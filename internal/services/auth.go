package services

import (
	"errors"

	"github.com/ViniciusLugli/mindshelf/internal/dtos/requests"
	"github.com/ViniciusLugli/mindshelf/internal/dtos/responses"
	"github.com/ViniciusLugli/mindshelf/internal/models"
	"github.com/ViniciusLugli/mindshelf/internal/repositories"
	"github.com/ViniciusLugli/mindshelf/internal/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repositories.UserRepository
}

func NewAuthService(repo *repositories.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func CheckPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *AuthService) Register(dto requests.CreateUserRequest) (responses.AuthResponse, error) {
	if err := utils.ValidateJWTConfig(); err != nil {
		return responses.AuthResponse{}, err
	}

	hashedPassword, err := HashPassword(dto.Password)
	if err != nil {
		return responses.AuthResponse{}, err
	}

	user := &models.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: hashedPassword,
	}

	created, err := s.repo.Create(user)
	if err != nil {
		return responses.AuthResponse{}, err
	}

	token, err := utils.GenerateToken(created.ID, created.Email)
	if err != nil {
		return responses.AuthResponse{}, err
	}

	return responses.NewAuthResponse(token, *created), nil
}

func (s *AuthService) Login(dto requests.LoginRequest) (responses.AuthResponse, error) {
	user, err := s.repo.GetByEmail(dto.Email)
	if err != nil {
		return responses.AuthResponse{}, err
	}

	if !CheckPassword(dto.Password, user.Password) {
		return responses.AuthResponse{}, errors.New("Invalid Email or Password")
	}

	token, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		return responses.AuthResponse{}, err
	}

	return responses.NewAuthResponse(token, user), nil
}

func (s *AuthService) GetProfile(userID uuid.UUID) (responses.UserResponse, error) {
	user, err := s.repo.GetByID(userID)
	if err != nil {
		return responses.UserResponse{}, err
	}

	return responses.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
