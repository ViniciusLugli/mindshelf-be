package services

import (
	"errors"
	"math"

	"github.com/ViniciusLugli/mindshelf/internal/dtos/requests"
	"github.com/ViniciusLugli/mindshelf/internal/dtos/responses"
	"github.com/ViniciusLugli/mindshelf/internal/repositories"
	"github.com/google/uuid"
)

var ErrTaskGroupNotFound = errors.New("group not found")

type TaskService struct {
	repo      *repositories.TaskRepository
	groupRepo *repositories.GroupRepository
}

func NewTaskService(repo *repositories.TaskRepository, groupRepo *repositories.GroupRepository) *TaskService {
	return &TaskService{repo: repo, groupRepo: groupRepo}
}

func (s *TaskService) Create(dto requests.CreateTaskRequest, userID uuid.UUID) error {
	exists, err := s.groupRepo.ExistsByIDAndUserID(dto.GroupID, userID)
	if err != nil {
		return err
	}

	if !exists {
		return ErrTaskGroupNotFound
	}

	task := dto.ToModel()
	return s.repo.Create(&task)
}

func (s *TaskService) Update(dto requests.UpdateTaskRequest, userID uuid.UUID) error {
	task, err := s.repo.GetByID(dto.ID, userID)
	if err != nil {
		return err
	}

	if dto.Title != "" {
		task.Title = dto.Title
	}

	if dto.Notes != "" {
		task.Notes = dto.Notes
	}

	return s.repo.Update(&task)
}

func (s *TaskService) Delete(dto requests.DeleteTaskRequest, userID uuid.UUID) error {
	task, err := s.repo.GetByID(dto.ID, userID)
	if err != nil {
		return err
	}

	return s.repo.Delete(&task)
}

func (s *TaskService) GetTask(dto requests.GetTask, userID uuid.UUID) (responses.TaskResponse, error) {
	task, err := s.repo.GetByID(dto.ID, userID)
	if err != nil {
		return responses.TaskResponse{}, err
	}

	return responses.NewTaskResponse(task), nil
}

func (s *TaskService) GetAllTasks(dto requests.GetAllTasks, userID uuid.UUID) (responses.PaginatedResponse[responses.TaskResponse], error) {
	offset := (dto.Page - 1) * dto.Limit

	tasks, count, err := s.repo.GetAll(dto.Limit, offset, userID)
	if err != nil {
		return responses.PaginatedResponse[responses.TaskResponse]{}, err
	}

	totalPages := math.Ceil(float64(count) / float64(dto.Limit))
	return responses.NewPaginatedResponse(tasks, responses.NewTaskResponse, count, dto.Page, dto.Limit, int(totalPages)), nil
}

func (s *TaskService) GetAllTasksByTitle(dto requests.GetAllTasksByTitle, userID uuid.UUID) (responses.PaginatedResponse[responses.TaskResponse], error) {
	offset := (dto.Page - 1) * dto.Limit

	tasks, count, err := s.repo.GetAllByTitle(dto.Title, dto.Limit, offset, userID)
	if err != nil {
		return responses.PaginatedResponse[responses.TaskResponse]{}, err
	}

	totalPages := math.Ceil(float64(count) / float64(dto.Limit))
	return responses.NewPaginatedResponse(tasks, responses.NewTaskResponse, count, dto.Page, dto.Limit, int(totalPages)), nil
}
