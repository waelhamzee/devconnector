package post

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID  uuid.UUID `gorm:"type:uuid" json:"user_id"`
	Content string    `json:"content"`
	gorm.Model
}
