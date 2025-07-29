package connection

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Connection struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID   uuid.UUID `gorm:"type:uuid" json:"user_id"`
	TargetID uuid.UUID `gorm:"type:uuid" json:"target_id"`
	gorm.Model
}
