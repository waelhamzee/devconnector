package auth

import (
	u "github.com/waelhamzee/devconnector/pkg/user"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) FindByEmail(email string) (*u.User, error) {
	var user u.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *AuthRepository) Create(user *u.User) error {
	return r.db.Create(user).Error
}
