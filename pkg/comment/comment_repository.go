package comment

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

func (r *Repository) Create(comment *Comment) error {
	return r.DB.Create(comment).Error
}

func (r *Repository) FindByPostID(postID uuid.UUID) ([]Comment, error) {
	var comments []Comment
	if err := r.DB.Where("post_id = ?", postID).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *Repository) Delete(id uuid.UUID) error {
	return r.DB.Delete(&Comment{}, "id = ?", id).Error
}
