package services

import (
	"math"

	"github.com/ViniciusLugli/mindshelf/internal/dtos/requests"
	"github.com/ViniciusLugli/mindshelf/internal/dtos/responses"
	"github.com/ViniciusLugli/mindshelf/internal/repositories"
)

type GroupService struct {
	repo *repositories.GroupRepository
}

func NewGroupService(repo *repositories.GroupRepository) *GroupService {
	return &GroupService{repo: repo}
}

func (s *GroupService) Create(dto requests.CreateGroupRequest) error {
	group := dto.ToModel()
	return s.repo.Create(&group)
}

func (s *GroupService) Update(dto requests.UpdateGroupRequest) error {
	group, err := s.repo.GetByID(dto.ID)
	if err != nil {
		return err
	}

	group.Name = dto.Name
	return s.repo.Update(&group)
}

func (s *GroupService) Delete(dto requests.DeleteGroupRequest) error {
	group, err := s.repo.GetByID(dto.ID)
	if err != nil {
		return err
	}

	return s.repo.Delete(&group)
}

func (s *GroupService) GetGroupByID(dto requests.GetGroupByID) (responses.GroupResponse, error) {
	group, err := s.repo.GetByID(dto.ID)
	if err != nil {
		return responses.GroupResponse{}, err
	}

	return responses.NewGroupRespone(group), nil
}

func (s *GroupService) GetGroupByName(dto requests.GetAllGroupsByName) (responses.PaginatedResponse[responses.GroupResponse], error) {
	offset := (dto.Page - 1) * dto.Limit

	groups, count, err := s.repo.GetAllByName(dto.Name, dto.Limit, offset)
	if err != nil {
		return responses.PaginatedResponse[responses.GroupResponse]{}, err
	}

	totalPages := math.Ceil(float64(count) / float64(dto.Limit))

	return responses.NewPaginatedResponse(groups, responses.NewGroupRespone, count, dto.Page, dto.Limit, int(totalPages)), nil
}

func (s *GroupService) GetAll(dto requests.GetAllGroups) (responses.PaginatedResponse[responses.GroupResponse], error) {
	offset := (dto.Page - 1) * dto.Limit

	groups, count, err := s.repo.GetAll(dto.Limit, offset)
	if err != nil {
		return responses.PaginatedResponse[responses.GroupResponse]{}, err
	}

	totalPages := math.Ceil(float64(count) / float64(dto.Limit))

	return responses.NewPaginatedResponse(groups, responses.NewGroupRespone, count, dto.Page, dto.Limit, int(totalPages)), nil
}
