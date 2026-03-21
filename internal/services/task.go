package services

import (
	"errors"
	"math"

	"github.com/ViniciusLugli/mindshelf/internal/dtos/requests"
	"github.com/ViniciusLugli/mindshelf/internal/dtos/responses"
	"github.com/ViniciusLugli/mindshelf/internal/repositories"
	"github.com/google/uuid"
)

type TaskService struct {
	repo *repositories.TaskRepository
}

func NewTaskService(repo *repositories.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) Create(dto requests.CreateTaskRequest) error {
	task := dto.ToModel()
	return s.repo.Create(&task)
}

func (s *TaskService) Update(dto requests.UpdateTaskRequest) error {
	task, err := s.repo.GetByID(dto.ID)
	if err != nil {
		return err
	}

	task.Title = dto.Title
	task.Notes = dto.Notes

	return s.repo.Update(&task)
}

func (s *TaskService) Delete(dto requests.DeleteTaskRequest) error {
	task, err := s.repo.GetByID(dto.ID)
	if err != nil {
		return err
	}

	return s.repo.Delete(&task)
}

func (s *TaskService) GetTask(dto requests.GetTask) (responses.TaskResponse, error) {
	if dto.ID != uuid.Nil {
		task, err := s.repo.GetByID(dto.ID)
		if err != nil {
			return responses.TaskResponse{}, err
		}

		return responses.NewTaskResponse(task), nil
	}

	if dto.Title != "" {
		task, err := s.repo.GetByTitle(dto.Title)
		if err != nil {
			return responses.TaskResponse{}, err
		}

		return responses.NewTaskResponse(task), nil
	}

	return responses.TaskResponse{}, errors.New("no valid params was passed")
}

func (s *TaskService) GetAllTasks(dto requests.GetAllTasks) (responses.PaginatedResponse[responses.TaskResponse], error) {
	offset := (dto.Page - 1) * dto.Limit

	tasks, count, err := s.repo.GetAll(dto.Limit, offset)
	if err != nil {
		return responses.PaginatedResponse[responses.TaskResponse]{}, err
	}

	totalPages := math.Ceil(float64(count) / float64(dto.Limit))
	return responses.NewPaginatedResponse(tasks, responses.NewTaskResponse, count, dto.Page, dto.Limit, int(totalPages)), nil
}
