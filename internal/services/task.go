package services

import (
	"math"

	"github.com/ViniciusLugli/mindshelf/internal/dtos/requests"
	"github.com/ViniciusLugli/mindshelf/internal/dtos/responses"
	"github.com/ViniciusLugli/mindshelf/internal/repositories"
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
	task, err := s.repo.GetByID(dto.ID)
	if err != nil {
		return responses.TaskResponse{}, err
	}

	return responses.NewTaskResponse(task), nil
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

func (s *TaskService) GetAllTasksByTitle(dto requests.GetAllTasksByTitle) (responses.PaginatedResponse[responses.TaskResponse], error) {
	offset := (dto.Page - 1) * dto.Limit

	tasks, count, err := s.repo.GetAllByTitle(dto.Title, dto.Limit, offset)
	if err != nil {
		return responses.PaginatedResponse[responses.TaskResponse]{}, err
	}

	totalPages := math.Ceil(float64(count) / float64(dto.Limit))
	return responses.NewPaginatedResponse(tasks, responses.NewTaskResponse, count, dto.Page, dto.Limit, int(totalPages)), nil
}
