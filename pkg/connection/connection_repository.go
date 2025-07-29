package connection

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) Create(conn *Connection) error {
	return r.DB.Create(conn).Error
}

func (r *Repository) FindByUserID(userID uuid.UUID) ([]Connection, error) {
	var conns []Connection
	if err := r.DB.Where("user_id = ?", userID).Find(&conns).Error; err != nil {
		return nil, err
	}
	return conns, nil
}

func (r *Repository) Delete(userID, targetID uuid.UUID) error {
	return r.DB.Where("user_id = ? AND target_id = ?", userID, targetID).Delete(&Connection{}).Error
}
