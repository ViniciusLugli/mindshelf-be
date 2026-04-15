package services

import (
	"errors"
	"strings"

	"github.com/ViniciusLugli/mindshelf/internal/dtos/requests"
	"github.com/ViniciusLugli/mindshelf/internal/dtos/responses"
	"github.com/ViniciusLugli/mindshelf/internal/models"
	"github.com/ViniciusLugli/mindshelf/internal/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var ErrSharedTaskNotFound = errors.New("task not found")

type MessageService struct {
	repo     *repositories.MessageRepository
	taskRepo *repositories.TaskRepository
}

func NewMessageService(repo *repositories.MessageRepository, taskRepo *repositories.TaskRepository) *MessageService {
	return &MessageService{repo: repo, taskRepo: taskRepo}
}

func (s *MessageService) GetMessages(userID, correspondentID uuid.UUID, page, limit int) (responses.PaginatedResponse[responses.MessageResponse], error) {
	msgs, total, err := s.repo.GetMessages(userID, correspondentID, page, limit)
	if err != nil {
		var empty responses.PaginatedResponse[responses.MessageResponse]
		return empty, err
	}

	totalPages := 0
	if limit > 0 {
		totalPages = int((total + int64(limit) - 1) / int64(limit))
	}

	pag := responses.NewPaginatedResponse(msgs, func(m models.Message) responses.MessageResponse {
		return responses.NewMessageResponse(m)
	}, total, page, limit, totalPages)

	return pag, nil
}

func (s *MessageService) SendMessage(senderID uuid.UUID, dto requests.SendChatRequest) (responses.MessageResponse, error) {
	msg := models.Message{
		Type:       models.MessageTypeText,
		SenderID:   senderID,
		ReceiverID: dto.ToUserID,
		Content:    dto.Content,
	}

	return s.createAndLoadMessage(&msg)
}

func (s *MessageService) ShareTask(senderID uuid.UUID, dto requests.ShareTaskRequest) (responses.MessageResponse, error) {
	task, err := s.taskRepo.GetByID(dto.TaskID, senderID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return responses.MessageResponse{}, ErrSharedTaskNotFound
		}

		return responses.MessageResponse{}, err
	}

	messageContent := strings.TrimSpace(dto.Content)
	if messageContent == "" {
		messageContent = "Shared a task"
	}

	msg := models.Message{
		Type:                 models.MessageTypeSharedTask,
		SenderID:             senderID,
		ReceiverID:           dto.ToUserID,
		Content:              messageContent,
		SharedTaskSourceID:   &task.ID,
		SharedTaskTitle:      task.Title,
		SharedTaskNotes:      task.Notes,
		SharedTaskGroupName:  task.Group.Name,
		SharedTaskGroupColor: task.Group.Color,
	}

	return s.createAndLoadMessage(&msg)
}

func (s *MessageService) createAndLoadMessage(msg *models.Message) (responses.MessageResponse, error) {
	created, err := s.repo.Create(msg)
	if err != nil {
		return responses.MessageResponse{}, err
	}

	fullMsg, err := s.repo.GetByID(created.ID)
	if err == nil {
		return responses.NewMessageResponse(fullMsg), nil
	}

	return responses.NewMessageResponse(*created), nil
}

func (s *MessageService) GetConversation(userID, withUserID uuid.UUID) ([]responses.MessageResponse, error) {
	msgs, err := s.repo.GetConversation(userID, withUserID)
	if err != nil {
		return nil, err
	}

	out := make([]responses.MessageResponse, len(msgs))
	for i, m := range msgs {
		out[i] = responses.NewMessageResponse(m)
	}

	return out, nil
}

func (s *MessageService) GetChats(userID uuid.UUID) ([]responses.ChatResponse, error) {
	summaries, err := s.repo.GetChats(userID)
	if err != nil {
		return nil, err
	}

	out := make([]responses.ChatResponse, len(summaries))
	for i, ssum := range summaries {
		out[i] = responses.ChatResponse{
			Friend:      responses.NewUserResponse(ssum.Friend),
			LastMessage: responses.NewMessageResponse(ssum.LastMessage),
			UnreadCount: ssum.UnreadCount,
		}
	}

	return out, nil
}

func (s *MessageService) MarkMessagesAsRead(userID uuid.UUID, dto requests.MarkMessagesReadRequest) (responses.MarkMessagesReadResponse, error) {
	updated, readAt, err := s.repo.MarkConversationAsRead(userID, dto.WithUserID, dto.UpToMessageID)
	if err != nil {
		return responses.MarkMessagesReadResponse{}, err
	}

	return responses.MarkMessagesReadResponse{
		WithUserID: dto.WithUserID,
		Updated:    updated,
		ReadAt:     readAt,
	}, nil
}
