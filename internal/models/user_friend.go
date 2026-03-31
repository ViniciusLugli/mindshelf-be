package models

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	Pending  Status = "pending"
	Accepted Status = "accepted"
	Rejected Status = "rejected"
)

func (s *Status) Scan(value interface{}) error {
	if value == nil {
		*s = ""
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		if str, ok := value.(string); ok {
			*s = Status(str)
			return nil
		}
		return fmt.Errorf("Can not convert %v to Status", value)
	}
	*s = Status(bytes)
	return nil
}

func (s Status) Value() (driver.Value, error) {
	return string(s), nil
}

type UserFriend struct {
	UserID    uuid.UUID `gorm:"primaryKey"`
	FriendID  uuid.UUID `gorm:"primaryKey"`
	Status    Status    `gorm:"type:status"`
	CreatedAt time.Time

	User   User `gorm:"foreignKey:UserID"`
	Friend User `gorm:"foreignKey:FriendID"`
}
