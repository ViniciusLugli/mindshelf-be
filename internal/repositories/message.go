package repositories

import (
	"time"

	"github.com/ViniciusLugli/mindshelf/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{db: db}
}

func (r *MessageRepository) Create(message *models.Message) (*models.Message, error) {
	if err := r.db.Create(message).Error; err != nil {
		return nil, err
	}

	return message, nil
}

func (r *MessageRepository) GetByID(id uuid.UUID) (models.Message, error) {
	var message models.Message
	err := r.db.Preload("Sender").Preload("Receiver").First(&message, "id = ?", id).Error
	return message, err
}

func (r *MessageRepository) WithTransaction(fn func(tx *gorm.DB) error) error {
	return r.db.Transaction(fn)
}

func (r *MessageRepository) GetSharedTaskMessageForReceiverTx(tx *gorm.DB, id, receiverID uuid.UUID) (models.Message, error) {
	var message models.Message
	err := tx.
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("id = ? AND receiver_id = ? AND type = ?", id, receiverID, models.MessageTypeSharedTask).
		First(&message).Error
	return message, err
}

func (r *MessageRepository) SetImportedTaskTx(tx *gorm.DB, messageID, taskID uuid.UUID) error {
	return tx.Model(&models.Message{}).
		Where("id = ?", messageID).
		Updates(map[string]any{"imported_task_id": taskID}).Error
}

func (r *MessageRepository) GetMessages(userID, correspondentID uuid.UUID, page, limit int) ([]models.Message, int64, error) {
	var messages []models.Message
	var total int64

	if page < 1 {
		page = 1
	}
	if limit <= 0 {
		limit = 20
	}

	offset := (page - 1) * limit

	baseQuery := r.db.Model(&models.Message{}).Where("(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)",
		userID, correspondentID, correspondentID, userID)

	if err := baseQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := baseQuery.Preload("Sender").Preload("Receiver").Order("created_at desc").Limit(limit).Offset(offset).Find(&messages).Error; err != nil {
		return nil, 0, err
	}

	return messages, total, nil
}

type ChatSummary struct {
	Friend      models.User
	LastMessage models.Message
	UnreadCount int64
}

func (r *MessageRepository) GetConversation(userID, withUserID uuid.UUID) ([]models.Message, error) {
	var messages []models.Message
	err := r.db.Preload("Sender").Preload("Receiver").Where("(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)",
		userID, withUserID, withUserID, userID).Order("created_at asc").Find(&messages).Error
	return messages, err
}

func (r *MessageRepository) GetChats(userID uuid.UUID) ([]ChatSummary, error) {
	type idRow struct {
		CorrespondentID uuid.UUID `gorm:"column:correspondent_id"`
	}

	var rows []idRow
	raw := `SELECT DISTINCT CASE WHEN sender_id = ? THEN receiver_id ELSE sender_id END AS correspondent_id FROM messages WHERE sender_id = ? OR receiver_id = ?`
	if err := r.db.Raw(raw, userID, userID, userID).Scan(&rows).Error; err != nil {
		return nil, err
	}

	summaries := make([]ChatSummary, 0, len(rows))
	for _, row := range rows {
		var friend models.User
		if err := r.db.First(&friend, "id = ?", row.CorrespondentID).Error; err != nil {
			continue
		}

		var last models.Message
		if err := r.db.Preload("Sender").Preload("Receiver").Where("(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)",
			userID, row.CorrespondentID, row.CorrespondentID, userID).Order("created_at desc").Limit(1).Find(&last).Error; err != nil {
			continue
		}

		var unreadCount int64
		if err := r.db.Model(&models.Message{}).
			Where("sender_id = ? AND receiver_id = ? AND read_at IS NULL", row.CorrespondentID, userID).
			Count(&unreadCount).Error; err != nil {
			continue
		}

		summaries = append(summaries, ChatSummary{Friend: friend, LastMessage: last, UnreadCount: unreadCount})
	}

	return summaries, nil
}

func (r *MessageRepository) MarkConversationAsRead(readerID, withUserID uuid.UUID, upToMessageID *uuid.UUID) (int64, *time.Time, error) {
	query := r.db.Model(&models.Message{}).
		Where("sender_id = ? AND receiver_id = ? AND read_at IS NULL", withUserID, readerID)

	if upToMessageID != nil {
		var cutoff models.Message
		err := r.db.Select("created_at").
			Where("id = ? AND ((sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?))",
				*upToMessageID,
				readerID, withUserID,
				withUserID, readerID,
			).
			First(&cutoff).Error
		if err != nil {
			return 0, nil, err
		}

		query = query.Where("created_at <= ?", cutoff.CreatedAt)
	}

	now := time.Now().UTC()
	result := query.Update("read_at", now)
	if result.Error != nil {
		return 0, nil, result.Error
	}

	return result.RowsAffected, &now, nil
}
