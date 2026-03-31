package services

import (
	"github.com/ViniciusLugli/mindshelf/internal/dtos/requests"
	"github.com/ViniciusLugli/mindshelf/internal/dtos/responses"
	"github.com/ViniciusLugli/mindshelf/internal/models"
	"github.com/ViniciusLugli/mindshelf/internal/repositories"
	"github.com/google/uuid"
)

type MessageService struct {
	repo *repositories.MessageRepository
}

func NewMessageService(repo *repositories.MessageRepository) *MessageService {
	return &MessageService{repo: repo}
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
		SenderID:   senderID,
		ReceiverID: dto.ToUserID,
		Content:    dto.Content,
	}

	created, err := s.repo.Create(&msg)
	if err != nil {
		return responses.MessageResponse{}, err
	}

	fullMsgs, _, _ := s.repo.GetMessages(senderID, dto.ToUserID, 1, 1)
	if len(fullMsgs) > 0 {
		return responses.NewMessageResponse(fullMsgs[0]), nil
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
		}
	}

	return out, nil
}
