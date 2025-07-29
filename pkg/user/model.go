package user

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name     string    `json:"name"`
	Email    string    `gorm:"uniqueIndex" json:"email"`
	Password string    `json:"-"`
}
