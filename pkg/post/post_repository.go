package post

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

func (r *Repository) Create(post *Post) error {
	return r.DB.Create(post).Error
}

func (r *Repository) FindByID(id uuid.UUID) (*Post, error) {
	var post Post
	if err := r.DB.First(&post, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *Repository) FindAll() ([]Post, error) {
	var posts []Post
	if err := r.DB.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *Repository) Delete(id uuid.UUID) error {
	return r.DB.Delete(&Post{}, "id = ?", id).Error
}
