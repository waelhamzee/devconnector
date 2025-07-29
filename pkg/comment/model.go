package comment

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	PostID  uuid.UUID `gorm:"type:uuid" json:"post_id"`
	UserID  uuid.UUID `gorm:"type:uuid" json:"user_id"`
	Content string    `json:"content"`
	gorm.Model
}
